package client

import (
	"context"
	"time"

	cachecash "github.com/cachecashproject/go-cachecash"
	"github.com/cachecashproject/go-cachecash/ccmsg"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/trace"
)

type fetchGroup struct {
	bundle *ccmsg.TicketBundle
	err    error
	notify []chan DownloadResult
}

func (cl *client) schedule(ctx context.Context, path string, queue chan *fetchGroup) {
	defer close(queue)

	var chunkSize uint64
	var rangeBegin uint64

	minimumBacklogDepth := uint64(0)
	bundleRequestInterval := 0
	schedulerNotify := make(chan bool, 64)

	for {
		cl.l.WithFields(logrus.Fields{
			"chunkRangeBegin": rangeBegin,
			"byteRangeBegin":  rangeBegin * chunkSize,
		}).Info("requesting bundle")
		bundles, err := cl.requestBundles(ctx, path, rangeBegin*chunkSize)
		if err != nil {
			err = errors.Wrapf(err, "failed to fetch chunk-group at offset %d", rangeBegin)
			queue <- &fetchGroup{
				err: err,
			}
			cl.l.Error("encountered an error, shutting down scheduler")
			return
		}

		for _, bundle := range bundles {
			chunks := len(bundle.TicketRequest)
			cl.l.WithFields(logrus.Fields{
				"len(chunks)": chunks,
			}).Info("pushing bundle to downloader")

			// For each chunk in TicketBundle, dispatch a request to the appropriate cache.
			chunkResults := make([]*chunkRequest, chunks)

			fetchGroup := &fetchGroup{
				bundle: bundle,
				notify: []chan DownloadResult{},
			}

			for i := 0; i < chunks; i++ {
				b := &chunkRequest{
					bundle: bundle,
					idx:    i,
				}
				chunkResults[i] = b

				// TODO: we identify caches by public key but connections are currently distinguished by addr
				cid := (cacheID)(bundle.CacheInfo[i].Addr.ConnectionString())
				cc, ok := cl.cacheConns[cid]
				if !ok {
					var err error
					ci := bundle.CacheInfo[i]

					// XXX: It's problematic to pass ctx here, because canceling the context will destroy the cache connections!
					// (It should only cancel this particular chunk-group request.)
					cc, err = cl.publisherConn.newCacheConnection(ctx, cl.l, ci.Addr.ConnectionString(), ci.Pubkey.GetPublicKey())
					if err != nil {
						cl.l.WithError(err).Error("failed to connect to cache")
						b.err = err
					}

					cl.cacheConns[cid] = cc
					go cc.Run(ctx)
				}

				clientNotify := make(chan DownloadResult, 128)
				fetchGroup.notify = append(fetchGroup.notify, clientNotify)
				cc.QueueRequest(DownloadTask{
					req:             b,
					clientNotify:    clientNotify,
					schedulerNotify: schedulerNotify,
				})
			}

			queue <- fetchGroup
			rangeBegin += uint64(chunks)

			if rangeBegin >= bundle.Metadata.ChunkCount() {
				cl.l.Info("got all bundles, terminating scheduler")
				return
			}
			chunkSize = bundle.Metadata.ChunkSize

			minimumBacklogDepth = uint64(bundle.Metadata.MinimumBacklogDepth)
			bundleRequestInterval = int(bundle.Metadata.BundleRequestInterval)
		}

		cl.waitUntilNextRequest(schedulerNotify, minimumBacklogDepth, bundleRequestInterval)
	}
}

func (cl *client) waitUntilNextRequest(schedulerNotify chan bool, minimumBacklogDepth uint64, bundleRequestInterval int) {
	for {
		interval := time.Duration(bundleRequestInterval) * time.Second
		intervalRemaining := interval - time.Since(cl.lastBundleRequest)

		select {
		case <-schedulerNotify:
			cl.l.WithFields(logrus.Fields{
				"minimumBacklogDepth": minimumBacklogDepth,
			}).Debug("checking cache backlog depth")
			if cl.checkBacklogDepth(minimumBacklogDepth) {
				cl.l.Info("cache backlog is running low, requesting new bundle")
				return
			}
		case <-time.After(intervalRemaining):
			cl.l.WithFields(logrus.Fields{
				"interval": bundleRequestInterval,
			}).Info("interval reached, requesting new bundles")
			return
		}
	}
}

func (cl *client) checkBacklogDepth(n uint64) bool {
	for _, c := range cl.cacheConns {
		if c.BacklogLength() <= n {
			return true
		}
	}
	return false
}

type chunkRequest struct {
	bundle *ccmsg.TicketBundle
	idx    int

	encData []byte // Singly-encrypted data.
	err     error
}

func (cl *client) requestBundles(ctx context.Context, path string, rangeBegin uint64) ([]*ccmsg.TicketBundle, error) {
	ctx, span := trace.StartSpan(ctx, "cachecash.com/Client/requestBundle")
	defer span.End()
	cl.l.Info("enumerating backlog length")

	backlogs := make(map[string]uint64)
	for _, cc := range cl.cacheConns {
		cl.l.WithFields(logrus.Fields{
			"cache": cc.PublicKey(),
		}).Info("backlog length: ", cc.BacklogLength())
		backlogs[string(cc.PublicKeyBytes())] = cc.BacklogLength()
	}

	req := &ccmsg.ContentRequest{
		ClientPublicKey: cachecash.PublicKeyMessage(cl.publicKey),
		Path:            path,
		RangeBegin:      rangeBegin,
		RangeEnd:        0, // "continue to the end of the object"
		BacklogDepth:    backlogs,
	}
	cl.l.Infof("sending content request to publisher: %v", req)

	// Send request to publisher; get TicketBundle in response.
	resp, err := cl.publisherConn.GetContent(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request bundle from publisher")
	}
	bundles := resp.Bundles
	for _, bundle := range bundles {
		cl.l.Info("got ticket bundle from publisher for escrow: ", bundle.GetRemainder().GetEscrowId())
		// cl.l.Debugf("got ticket bundle from publisher: %v", proto.MarshalTextString(bundle))
	}

	cl.lastBundleRequest = time.Now()

	return bundles, nil
}
