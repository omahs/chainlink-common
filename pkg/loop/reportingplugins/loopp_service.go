package reportingplugins

import (
	"context"
	"fmt"
	"os/exec"

	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

var _ ocrtypes.ReportingPluginFactory = (*LOOPPService)(nil)

// LOOPPService is a [types.Service] that maintains an internal [types.PluginClient].
type LOOPPService struct {
	internal.PluginService[*GRPCService[types.PluginProvider], types.ReportingPluginFactory]
}

// NewLOOPPService returns a new [*PluginService].
// cmd must return a new exec.Cmd each time it is called.
// We use a `conn` here rather than a provider so that we can enforce proxy providers being passed in.
func NewLOOPPService(
	lggr logger.Logger,
	grpcOpts loop.GRPCOpts,
	cmd func() *exec.Cmd,
	config types.ReportingPluginServiceConfig,
	providerConn grpc.ClientConnInterface,
	pipelineRunner types.PipelineRunnerService,
	telemetryService types.TelemetryService,
	errorLog types.ErrorLog,
) *LOOPPService {
	newService := func(ctx context.Context, instance any) (types.ReportingPluginFactory, services.HealthReporter, error) {
		plug, ok := instance.(types.ReportingPluginClient)
		if !ok {
			return nil, nil, fmt.Errorf("expected GenericPluginClient but got %T", instance)
		}
		//TODO plug.Start(ctx)? (how to close?)
		factory, err := plug.NewReportingPluginFactory(ctx, config, providerConn, pipelineRunner, telemetryService, errorLog)
		if err != nil {
			return nil, nil, err
		}
		return factory, plug, nil
	}
	stopCh := make(chan struct{})
	lggr = logger.Named(lggr, "GenericService")
	var ps LOOPPService
	broker := internal.BrokerConfig{StopCh: stopCh, Logger: lggr, GRPCOpts: grpcOpts}
	ps.Init(PluginServiceName, &GRPCService[types.PluginProvider]{BrokerConfig: broker}, newService, lggr, cmd, stopCh)
	return &ps
}

func (g *LOOPPService) NewReportingPlugin(config ocrtypes.ReportingPluginConfig) (ocrtypes.ReportingPlugin, ocrtypes.ReportingPluginInfo, error) {
	if err := g.Wait(); err != nil {
		return nil, ocrtypes.ReportingPluginInfo{}, err
	}
	return g.Service.NewReportingPlugin(config)
}
