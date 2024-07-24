package coralogix

import (
	"encoding/gob"

	"github.com/doublecloud/tross/library/go/core/log"
	"github.com/doublecloud/tross/library/go/core/metrics"
	"github.com/doublecloud/tross/library/go/core/xerrors"
	"github.com/doublecloud/tross/transfer_manager/go/pkg/abstract"
	"github.com/doublecloud/tross/transfer_manager/go/pkg/abstract/coordinator"
	server "github.com/doublecloud/tross/transfer_manager/go/pkg/abstract/model"
	"github.com/doublecloud/tross/transfer_manager/go/pkg/middlewares"
	"github.com/doublecloud/tross/transfer_manager/go/pkg/providers"
)

func init() {
	gob.Register(new(CoralogixDestination))
	providers.Register(ProviderType, New)
	abstract.RegisterProviderName(ProviderType, "Coralogix")
	server.RegisterDestination(ProviderType, destinationModelFactory)
}

func destinationModelFactory() server.Destination {
	return new(CoralogixDestination)
}

const ProviderType = abstract.ProviderType("coralogix")

// To verify providers contract implementation
var (
	_ providers.Sinker = (*Provider)(nil)
)

type Provider struct {
	logger   log.Logger
	registry metrics.Registry
	cp       coordinator.Coordinator
	transfer *server.Transfer
}

func (p Provider) Sink(config middlewares.Config) (abstract.Sinker, error) {
	dst, ok := p.transfer.Dst.(*CoralogixDestination)
	if !ok {
		return nil, xerrors.Errorf("unexpected target type: %T", p.transfer.Dst)
	}
	return NewSink(dst, p.logger, p.registry)
}

func (p Provider) Type() abstract.ProviderType {
	return ProviderType
}

func New(lgr log.Logger, registry metrics.Registry, cp coordinator.Coordinator, transfer *server.Transfer) providers.Provider {
	return &Provider{
		logger:   lgr,
		registry: registry,
		cp:       cp,
		transfer: transfer,
	}
}
