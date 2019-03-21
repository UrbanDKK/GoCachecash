package publisher

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/cachecashproject/go-cachecash/ccmsg"
	"github.com/cachecashproject/go-cachecash/common"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ed25519"
	"google.golang.org/grpc"
)

// An Application is the top-level content publisher.  It takes a configuration struct.  Its children are the several
// protocol servers (that deal with clients, caches, and so forth).
type Application interface {
	common.StarterShutdowner
}

type ConfigFile struct {
	Config      *Config            `json:"config"`
	Escrows     []*Escrow          `json:"escrows"`
	UpstreamURL string             `json:"upstreamURL"`
	PrivateKey  ed25519.PrivateKey `json:"privateKey"`
	Database    string             `json:"database"`
}

// XXX: Right now, this is shared between the client- and cache-facing servers.
type Config struct {
	ClientProtocolAddr   string
	CacheProtocolAddr    string
	StatusAddr           string
	DefaultCacheDuration time.Duration
}

func (c *Config) FillDefaults() {
	if c.ClientProtocolAddr == "" {
		c.ClientProtocolAddr = ":8080"
	}
	if c.CacheProtocolAddr == "" {
		c.CacheProtocolAddr = ":8082"
	}
	if c.StatusAddr == "" {
		c.StatusAddr = ":8100"
	}
	if c.DefaultCacheDuration == 0 {
		c.DefaultCacheDuration = 5 * time.Minute
	}
}

type application struct {
	l *logrus.Logger

	clientProtocolServer *clientProtocolServer
	cacheProtocolServer  *cacheProtocolServer
	statusServer         *statusServer
	// TODO: ...
}

var _ Application = (*application)(nil)

// XXX: Should this take p as an argument, or be responsible for setting it up?
func NewApplication(l *logrus.Logger, p *ContentPublisher, conf *Config) (Application, error) {
	conf.FillDefaults()

	clientProtocolServer, err := newClientProtocolServer(l, p, conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create client protocol server")
	}

	cacheProtocolServer, err := newCacheProtocolServer(l, p, conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cache protocol server")
	}

	statusServer, err := newStatusServer(l, p, conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create status server")
	}

	return &application{
		l:                    l,
		clientProtocolServer: clientProtocolServer,
		cacheProtocolServer:  cacheProtocolServer,
		statusServer:         statusServer,
	}, nil
}

func (a *application) Start() error {
	if err := a.clientProtocolServer.Start(); err != nil {
		return errors.Wrap(err, "failed to start client protocol server")
	}
	if err := a.cacheProtocolServer.Start(); err != nil {
		return errors.Wrap(err, "failed to start cache protocol server")
	}
	if err := a.statusServer.Start(); err != nil {
		return errors.Wrap(err, "failed to start status server")
	}
	return nil
}

func (a *application) Shutdown(ctx context.Context) error {
	if err := a.clientProtocolServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "failed to shut down client protocol server")
	}
	if err := a.cacheProtocolServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "failed to shut down cache protocol server")
	}
	if err := a.statusServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "failed to shut down status server")
	}
	return nil
}

type clientProtocolServer struct {
	l          *logrus.Logger
	conf       *Config
	publisher  *ContentPublisher
	grpcServer *grpc.Server
	httpServer *http.Server
}

var _ common.StarterShutdowner = (*clientProtocolServer)(nil)

func newClientProtocolServer(l *logrus.Logger, p *ContentPublisher, conf *Config) (*clientProtocolServer, error) {
	grpcServer := grpc.NewServer()
	ccmsg.RegisterClientPublisherServer(grpcServer, &grpcClientPublisherServer{publisher: p})

	httpServer := wrapGrpc(grpcServer)

	return &clientProtocolServer{
		l:          l,
		conf:       conf,
		publisher:  p,
		grpcServer: grpcServer,
		httpServer: httpServer,
	}, nil
}

func wrapGrpc(grpcServer *grpc.Server) *http.Server {
	wrappedServer := grpcweb.WrapServer(grpcServer)

	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	return &http.Server{
		// Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}
}

func (s *clientProtocolServer) Start() error {
	s.l.Info("clientProtocolServer - Start - enter")

	lis, err := net.Listen("tcp", s.conf.ClientProtocolAddr)
	if err != nil {
		return errors.Wrap(err, "failed to bind listener")
	}

	// httpLis, err := net.Listen("tcp", s.conf.ClientProtocolHttpAddr)
	httpLis, err := net.Listen("tcp", ":8043")
	if err != nil {
		return errors.Wrap(err, "failed to bind listener")
	}

	go func() {
		// This will block until we call `Stop`.
		if err := s.grpcServer.Serve(lis); err != nil {
			s.l.WithError(err).Error("failed to serve clientProtocolServer(grpc)")
		}
	}()

	go func() {
		// This will block until we call `Stop`.
		if err := s.httpServer.Serve(httpLis); err != nil {
			s.l.WithError(err).Error("failed to serve clientProtocolServer(http)")
		}
	}()

	s.l.Info("clientProtocolServer - Start - exit")
	return nil
}

func (s *clientProtocolServer) Shutdown(ctx context.Context) error {
	// TODO: Should use `GracefulStop` until context expires, and then fall back on `Stop`.
	s.grpcServer.Stop()

	return nil
}

type cacheProtocolServer struct {
	l          *logrus.Logger
	conf       *Config
	publisher  *ContentPublisher
	grpcServer *grpc.Server
}

var _ common.StarterShutdowner = (*cacheProtocolServer)(nil)

func newCacheProtocolServer(l *logrus.Logger, p *ContentPublisher, conf *Config) (*cacheProtocolServer, error) {
	grpcServer := grpc.NewServer()
	ccmsg.RegisterCachePublisherServer(grpcServer, &grpcCachePublisherServer{publisher: p})

	return &cacheProtocolServer{
		l:          l,
		conf:       conf,
		publisher:  p,
		grpcServer: grpcServer,
	}, nil
}

func (s *cacheProtocolServer) Start() error {
	s.l.Info("cacheProtocolServer - Start - enter")

	lis, err := net.Listen("tcp", s.conf.CacheProtocolAddr)
	if err != nil {
		return errors.Wrap(err, "failed to bind listener")
	}

	go func() {
		// This will block until we call `Stop`.
		if err := s.grpcServer.Serve(lis); err != nil {
			s.l.WithError(err).Error("failed to serve cacheProtocolServer")
		}
	}()

	s.l.Info("cacheProtocolServer - Start - exit")
	return nil
}

func (s *cacheProtocolServer) Shutdown(ctx context.Context) error {
	// TODO: Should use `GracefulStop` until context expires, and then fall back on `Stop`.
	s.grpcServer.Stop()

	return nil
}
