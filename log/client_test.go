package log

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"testing"
	"time"

	"github.com/cachecashproject/go-cachecash/keypair"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	// set this in the environment to see log messages output as they are written
	// to the store.
	seeOutput = os.Getenv("DEBUG")
)

const (
	iters = 50
	count = 100
)

func muteStderr(t *testing.T) *os.File {
	// this hack allows us to swallow a lot of "hook failed" messages that are spit to stderr.
	var (
		r   *os.File
		err error
	)

	stderr := os.Stderr
	r, os.Stderr, err = os.Pipe()
	assert.Nil(t, err)
	go func() {
		// make errcheck happy
		if _, err := io.Copy(ioutil.Discard, r); err != nil {
			fmt.Println(err)
		}
		r.Close()
	}()
	return stderr
}

func assertLogEqual(t *testing.T, filename string) {
	r, err := NewReader(filename)
	assert.Nil(t, err)

	uniq := map[string]int{}

	for {
		e, err := r.NextProto()
		if err != nil {
			if err != io.EOF {
				assert.Nil(t, err)
			}
			break
		}
		uniq[e.Message]++
	}

	assert.Equal(t, count, len(uniq))
	for _, value := range uniq {
		assert.Equal(t, iters, value)
	}
}

func writeLogs(l *logrus.Logger, done chan struct{}) {
	if seeOutput == "" {
		l.SetOutput(ioutil.Discard)
	}

	ready := make(chan struct{})

	wg := &sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int, wg *sync.WaitGroup) {
			<-ready

			for x := 0; x < iters; x++ {
				l.Infof("%v", i)
			}

			wg.Done()
		}(i, wg)
	}

	close(ready)
	wg.Wait()
	close(done)
}

func getLogFile(t *testing.T, c *Client) string {
	d, err := os.Open(c.logDir)
	assert.Nil(t, err)
	fis, err := d.Readdir(-1)
	assert.Nil(t, err)
	var logFile string

	for _, fi := range fis {
		if fi.Mode()&os.ModeType == 0 && fi.Size() > 0 {
			logFile = fi.Name()
			break
		}
	}

	return path.Join(c.logDir, logFile)
}

func setupClient(t *testing.T, listenAddress, dir string, deliver bool) (*Client, *logrus.Logger, string) {
	if dir == "" {
		var err error
		dir, err = ioutil.TempDir("", "")
		assert.Nil(t, err)
	}

	config := DefaultConfig()
	config.DeliverLogs = deliver

	kp, err := keypair.Generate()
	assert.Nil(t, err)
	c, err := NewClient(listenAddress, "test", dir, true, true, config, kp)
	assert.Nil(t, err)

	l := logrus.New()
	l.Hooks.Add(NewHook(c))

	return c, l, dir
}

func TestClientBasic(t *testing.T) {
	c, l, dir := setupClient(t, "", "", false)
	defer os.RemoveAll(dir)

	l.WithFields(logrus.Fields{"hi": "there", "number": 8675309}).Info("test")
	assert.Nil(t, c.Close())

	r, err := NewReader(getLogFile(t, c))
	assert.Nil(t, err)
	e, err := r.NextProto()
	assert.Nil(t, err)
	assert.Equal(t, e.Message, "test")
	assert.Equal(t, logrus.Level(e.Level), logrus.InfoLevel)

	for key, value := range e.Fields.Fields {
		switch key {
		case "hi":
			assert.Equal(t, value.GetStringValue(), "there")
		case "number":
			assert.Equal(t, value.GetStringValue(), "8675309")
		}
	}

	_, err = r.NextProto()
	assert.Equal(t, err, io.EOF)
	assert.Nil(t, r.Close())
}

func TestClientParallelWriters(t *testing.T) {
	c, l, dir := setupClient(t, "", "", false)
	defer os.RemoveAll(dir)

	writeLogs(l, make(chan struct{}))
	assert.Nil(t, c.Close())
	assertLogEqual(t, getLogFile(t, c))
}

func TestClientShipLogs(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	f.Close()
	defer os.Remove(f.Name())

	tp := NewTestPipeServer(f.Name())
	go func() {
		assert.Nil(t, tp.Serve(":0"))
	}()
	defer tp.Close()

	c, l, dir := setupClient(t, tp.ListenAddress(), "", true)
	defer os.RemoveAll(dir)

	writeLogs(l, make(chan struct{}))
	time.Sleep(defaultTickInterval * 2)
	c.Close()

	fi, err := os.Stat(f.Name())
	assert.Nil(t, err)
	assert.NotEmpty(t, fi.Size())

	assertLogEqual(t, f.Name())
}

func TestClientShipLogsIncompleteRedeliver(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	f.Close()
	defer os.Remove(f.Name())

	tp := NewTestPipeServer(f.Name())
	go func() {
		assert.Nil(t, tp.Serve(":0"))
	}()
	defer tp.Close()

	c, l, dir := setupClient(t, tp.ListenAddress(), "", true)
	defer os.RemoveAll(dir)

	// the main difference in this test is here. what we do instead of
	// synchronizing the writes is to let them choatically spam the log and
	// detect issues that way hopefully.
	//
	// functionally, we start logging, then, before the first heartbeat can tick
	// (about 500ms as of this writing), we cut off the heartbeat.
	//
	done := make(chan struct{})
	go writeLogs(l, done)
	time.Sleep(defaultTickInterval / 2)
	c.heartbeatCancel()
	<-done // allow logging to finish
	assert.Nil(t, c.Close())

	// at this point the receiving end, which is attached to the pipe server
	// above, should still be 0 since nothing has been appended to it yet, since
	// nothing finished heartbeating.
	fi, err := os.Stat(f.Name())
	assert.Nil(t, err)
	assert.Empty(t, fi.Size())

	// we then setup a new client, which should immediately pick up the file and
	// deliver it -- no additional work is required, the client does this on
	// boot.
	c, _, _ = setupClient(t, tp.ListenAddress(), dir, true)
	time.Sleep(defaultTickInterval * 2)
	defer c.Close()

	// then we check the file again. ha ha! data!
	fi, err = os.Stat(f.Name())
	assert.Nil(t, err)
	assert.NotEmpty(t, fi.Size())

	// now we check for log equivalence.
	assertLogEqual(t, f.Name())
}

func TestClientShipLogsOnError(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	f.Close()
	defer os.Remove(f.Name())

	tp := NewTestPipeServer(f.Name())
	go func() {
		assert.Nil(t, tp.Serve(":0"))
	}()
	defer tp.Close()

	c, l, dir := setupClient(t, tp.ListenAddress(), "", true)
	defer os.RemoveAll(dir)

	// the main difference in this test from the above one, is that this one throws an error during initial delivery.
	// we are going to make sure the logs were delivered anyway.
	done := make(chan struct{})
	go writeLogs(l, done)
	tp.Mutex.Lock()
	tp.RaiseError = errors.New("welp")
	tp.Mutex.Unlock()
	time.Sleep(defaultTickInterval * 2)
	<-done
	assert.Nil(t, c.Close())

	tp.Mutex.Lock()
	// then we clear the error condition
	tp.RaiseError = nil
	tp.Mutex.Unlock()

	// we then setup a new client, which should immediately pick up the file and
	// deliver it -- no additional work is required, the client does this on
	// boot.
	c, _, _ = setupClient(t, tp.ListenAddress(), dir, true)
	time.Sleep(defaultTickInterval * 2)
	defer c.Close()

	// then we check the file. ha ha! data!
	fi, err := os.Stat(f.Name())
	assert.Nil(t, err)
	assert.NotEmpty(t, fi.Size())

	// now we check for log equivalence.
	assertLogEqual(t, f.Name())
}

func TestClientShipLogsOnErrorFlapper(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	f.Close()
	defer os.Remove(f.Name())

	tp := NewTestPipeServer(f.Name())
	go func() {
		assert.Nil(t, tp.Serve(":0"))
	}()
	defer tp.Close()

	c, l, dir := setupClient(t, tp.ListenAddress(), "", true)
	defer os.RemoveAll(dir)

	done := make(chan struct{})

	// this test "flaps" an error by turning it on and off a few times to see the
	// behavior when the system thinks it's ok then it's suddenly, well, not.
	go func() {
		for {
			select {
			case <-done:
				return
			default:
			}

			tp.Mutex.Lock()
			if tp.RaiseError == nil {
				tp.RaiseError = errors.New("welp")
			} else {
				tp.RaiseError = nil
			}
			tp.Mutex.Unlock()

			time.Sleep(100 * time.Millisecond)
		}
	}()
	go writeLogs(l, make(chan struct{}))

	// we wait a little longer here because the message delivery slows down with the above mutex.
	time.Sleep(defaultTickInterval * 5)
	close(done)
	assert.Nil(t, c.Close())

	// then we clear the error condition
	tp.Mutex.Lock()
	tp.RaiseError = nil
	tp.Mutex.Unlock()

	// we then setup a new client, which should immediately pick up the file and
	// deliver it -- no additional work is required, the client does this on
	// boot.
	c, _, _ = setupClient(t, tp.ListenAddress(), dir, true)
	time.Sleep(defaultTickInterval * 2)
	defer c.Close()

	// then we check the file. ha ha! data!
	fi, err := os.Stat(f.Name())
	assert.Nil(t, err)
	assert.NotEmpty(t, fi.Size())

	// now we check for log equivalence.
	assertLogEqual(t, f.Name())
}

func TestClientBasicError(t *testing.T) {
	c, l, dir := setupClient(t, "", "", false)
	defer os.RemoveAll(dir)

	stderr := muteStderr(t)
	defer func() {
		os.Stderr.Close()
		os.Stderr = stderr
	}()
	done := make(chan struct{})
	go writeLogs(l, done)

	c.errorMutex.Lock()
	c.Error = errors.New("welp")
	c.errorMutex.Unlock()

	assert.NotNil(t, l.Hooks.Fire(logrus.InfoLevel, logrus.NewEntry(l)))
	assert.NotNil(t, c.Write(logrus.NewEntry(l)))
	assert.NotNil(t, c.Close())
	<-done // let logrus finish before moving forward
}

func TestClientCannotConnect(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	f.Close()
	defer os.Remove(f.Name())

	tp := NewTestPipeServer(f.Name())
	sync := make(chan struct{})
	go func() {
		close(sync)
		tp.Serve(":0")
	}()
	<-sync
	tp.Close() // immediately shut it down

	c, l, dir := setupClient(t, tp.ListenAddress(), "", true) // the service is down here
	defer os.RemoveAll(dir)

	done := make(chan struct{})
	go writeLogs(l, done)
	time.Sleep(5 * time.Second)
	<-done

	fi, err := os.Stat(f.Name())
	assert.Nil(t, err)
	assert.Empty(t, fi.Size())

	c.Close()

	sync = make(chan struct{})
	go func() {
		close(sync)
		tp.Serve(":0")
	}()
	<-sync
	c, _, _ = setupClient(t, tp.ListenAddress(), dir, true)
	time.Sleep(5 * time.Second) // heartbeat interval

	fi, err = os.Stat(f.Name())
	assert.Nil(t, err)
	assert.NotEmpty(t, fi.Size())

	assert.Nil(t, c.Close())
}

func TestClientBackoff(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	f.Close()
	defer os.Remove(f.Name())

	tp := NewTestPipeServer(f.Name())
	go func() {
		assert.Nil(t, tp.Serve(":0"))
	}()
	defer tp.Close()

	c, l, dir := setupClient(t, tp.ListenAddress(), "", true)
	defer os.RemoveAll(dir)

	tp.Mutex.Lock()
	tp.RaiseError = errors.New("welp")
	tp.Mutex.Unlock()

	done := make(chan struct{})
	go writeLogs(l, done)
	time.Sleep(5 * time.Second)
	<-done

	fi, err := os.Stat(f.Name())
	assert.Nil(t, err)
	assert.Empty(t, fi.Size())

	tp.Mutex.Lock()
	tp.RaiseError = nil
	tp.Mutex.Unlock()
	time.Sleep(time.Second)

	fi, err = os.Stat(f.Name())
	assert.Nil(t, err)
	assert.Empty(t, fi.Size())

	time.Sleep(2 * time.Second)

	fi, err = os.Stat(f.Name())
	assert.Nil(t, err)
	assert.NotEmpty(t, fi.Size())

	assert.Nil(t, c.Close())
}
