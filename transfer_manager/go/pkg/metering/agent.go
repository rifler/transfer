package metering

import (
	"context"
	"sync"
	"time"

	"github.com/doublecloud/tross/transfer_manager/go/pkg/abstract"
)

var (
	commonAgent   MeteringAgent
	commonAgentMu sync.Mutex = sync.Mutex{}
)

type MeteringAgent interface {
	// RunPusher starts background metrics pushing process. RunPusher must not be called after Stop
	// Pusher is stopped either when Stop method is called or incoming Context is Done.
	RunPusher(ctx context.Context, interval time.Duration) error
	// Stop is used to stop metrics pusher (if it was run). Stop must not be called concurrently with RunPusher
	Stop() error
	SetOpts(config *MeteringOpts) error
	CountInputRows(items []abstract.ChangeItem)
	CountOutputRows(items []abstract.ChangeItem)
}

type Writer interface {
	Write(data string) error
	Close() error
}
