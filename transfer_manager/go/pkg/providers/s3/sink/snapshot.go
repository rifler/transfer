package sink

import "github.com/doublecloud/tross/transfer_manager/go/pkg/serializer"

type Snapshot interface {
	Read(buf []byte) (n int, err error)
	FeedChannel() chan<- []byte
	Close()
}

type snapshotHolder struct {
	uploadDone chan error
	snapshot   Snapshot
	serializer serializer.BatchSerializer
}
