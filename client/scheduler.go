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

	var chunkRangeBegin uint64
	var byteRangeBegin uint64

	minimumBacklogDepth := uint64(0)
	bundleRequestInterval := 0
	schedulerNotify := make(chan bool, 64)

	for {
		cl.l.WithFields(logrus.Fields{
			"chunkRangeBegin": chunkRangeBegin,
			"byteRangeBegin":  byteRangeBegin,
		}).Info("requesting bundle")
		bundles, err := cl.requestBundles(ctx, path, byteRangeBegin)
		if err != nil {
			err = errors.Wrapf(err, "failed to fetch chunk-group at chunk offset %d", chunkRangeBegin)
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

				ci := bundle.CacheInfo[i]
				pubKey := ci.Pubkey.GetPublicKey()
				cid := (cacheID)(string(pubKey))
				cc, ok := cl.cacheConns[cid]
				if !ok {
					var err error
					cc, err = cl.publisherConn.newCacheConnection(cl.l, ci.Addr.ConnectionString(), pubKey)
					if err != nil {
						cl.l.WithError(err).Error("failed to connect to cache")
						// In future we should resubmit the bundle - but this is better than panicing.
						fetchGroup.err = err
						fetchGroup.bundle = nil
						queue <- fetchGroup
						return
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
			chunkRangeBegin += uint64(chunks)
			byteRangeBegin += uint64(chunks) * bundle.Metadata.ChunkSize

			if chunkRangeBegin >= bundle.Metadata.ChunkCount() {
				cl.l.Info("got all bundles, terminating scheduler")
				return
			}

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