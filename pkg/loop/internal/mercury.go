package internal

import (
	"context"

	"google.golang.org/grpc"

	"github.com/mwitkow/grpc-proxy/proxy"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	mercury_common_internal "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/mercury/common"
	mercury_v3_internal "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/mercury/v3"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb"
	mercury_pb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/mercury"
	mercury_v3_pb "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/mercury/v3"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/mercury"
	mercury_v1 "github.com/smartcontractkit/chainlink-common/pkg/types/mercury/v1"
	mercury_v2 "github.com/smartcontractkit/chainlink-common/pkg/types/mercury/v2"
	mercury_v3 "github.com/smartcontractkit/chainlink-common/pkg/types/mercury/v3"
)

type MercuryAdapterClient struct {
	*pluginClient
	*serviceClient

	mercury mercury_pb.MercuryAdapterClient
}

func NewMercuryAdapterClient(broker Broker, brokerCfg BrokerConfig, conn *grpc.ClientConn) *MercuryAdapterClient {
	brokerCfg.Logger = logger.Named(brokerCfg.Logger, "MercuryAdapterClient")
	pc := newPluginClient(broker, brokerCfg, conn)
	// TODO: this is difficult to parse. the plugin client seems to be used to mean different things,
	// or at least is more 'primary' than the service client and the mercury client. i don't understand the asymmetry.
	return &MercuryAdapterClient{
		pluginClient:  pc,
		serviceClient: newServiceClient(pc.brokerExt, pc),
		mercury:       mercury_pb.NewMercuryAdapterClient(pc),
	}
}

func (c *MercuryAdapterClient) NewMercuryV1Factory(ctx context.Context,
	provider types.MercuryProvider, dataSource mercury_v1.DataSource) {
	panic("TODO")
}

func (c *MercuryAdapterClient) NewMercuryV2Factory(ctx context.Context,
	provider types.MercuryProvider, dataSource mercury_v2.DataSource) {
	panic("TODO")
}

// TODO: unlike median, the existing code to create a factory for mercury v3 does not use a provider. not sure what to do about that.
// https://github.com/smartcontractkit/chainlink-data-streams/blob/a6e3fe8ff2a12886b111f341639bae3cbf478501/mercury/v3/mercury.go#L65C92-L65C105
func (c *MercuryAdapterClient) NewMercuryV3Factory(ctx context.Context,
	provider types.MercuryProvider, dataSource mercury_v3.DataSource,
	occ mercury.OnchainConfigCodec, reportCodec mercury_v3.ReportCodec) (types.ReportingPluginFactory, error) {

	// every time a new client is created, we have to ensure that all the external dependencies are satisfied.
	// at this layer of the stack, all of those dependencies are other gRPC services.
	// some of those services are hosted in the same process as the client itself and others may be remote.
	newMercuryClientFn := func(ctx context.Context) (id uint32, deps resources, err error) {
		// the local resources for mercury are the DataSource
		dataSourceID, dsRes, err := c.serveNew("DataSource", func(s *grpc.Server) {
			mercury_v3_pb.RegisterDataSourceServer(s, mercury_v3_internal.NewDataSourceServer(dataSource))
		})
		if err != nil {
			return 0, nil, err
		}
		deps.Add(dsRes)

		// the proxyable resources for mercury are the Provider,  which may or may not be local to the client process. (legacy vs loopp)
		var (
			providerID  uint32
			providerRes resource
		)
		if grpcProvider, ok := provider.(GRPCClientConn); ok {
			providerID, providerRes, err = c.serve("MercuryProvider", proxy.NewProxy(grpcProvider.ClientConn()))
		} else {
			providerID, providerRes, err = c.serveNew("MercuryProvider", func(s *grpc.Server) {
				// this doesn't compile b/c there is a mixture of common and specific types. need to
				// figure out what can be reused from common and what needs to be mercury-specific.
				pb.RegisterServiceServer(s, &serviceServer{srv: provider})
				pb.RegisterOffchainConfigDigesterServer(s, &offchainConfigDigesterServer{impl: provider.OffchainConfigDigester()})
				pb.RegisterContractConfigTrackerServer(s, &contractConfigTrackerServer{impl: provider.ContractConfigTracker()})
				pb.RegisterContractTransmitterServer(s, &contractTransmitterServer{impl: provider.ContractTransmitter()})
				//pb.RegisterChainReaderServer(s, &chainReaderServer{impl: provider.ChainReader()})
				mercury_v3_pb.RegisterReportCodecServer(s, mercury_v3_internal.NewReportCodecServer(reportCodec))
				mercury_pb.RegisterOnchainConfigCodecServer(s, mercury_common_internal.NewOnchainConfigCodecServer(occ))
			})
		}
		if err != nil {
			return 0, nil, err
		}
		deps.Add(providerRes)

		reply, err := c.mercury.NewMercuryV3Factory(ctx, &mercury_pb.NewMercuryV3FactoryRequest{
			MercuryProviderID: providerID,
			DataSourceV3ID:    dataSourceID,
		})
		if err != nil {
			return 0, nil, err
		}
		return reply.MercuryV3FactoryID, deps, nil
	}

	cc := c.newClientConn("MercuryV3Factory", newMercuryClientFn)
	return newReportingPluginFactoryClient(c.pluginClient.brokerExt, cc), nil
}
