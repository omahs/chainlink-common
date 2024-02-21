package internal

import (
	libocr "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type pluginProviderClient struct {
	*configProviderClient
	contractTransmitter libocr.ContractTransmitter
	chainReader         types.ChainReader
	codec               types.Codec
}

func (p *pluginProviderClient) ClientConn() grpc.ClientConnInterface { return p.cc }

func newPluginProviderClient(b *BrokerExt, cc grpc.ClientConnInterface) *pluginProviderClient {
	p := &pluginProviderClient{configProviderClient: newConfigProviderClient(b.WithName("PluginProviderClient"), cc)}
	p.contractTransmitter = &contractTransmitterClient{b, pb.NewContractTransmitterClient(p.cc)}
	return p
}

func (p *pluginProviderClient) ContractTransmitter() libocr.ContractTransmitter {
	return p.contractTransmitter
}

type PluginProviderServer struct{}

func (p PluginProviderServer) ConnToProvider(conn grpc.ClientConnInterface, broker Broker, brokerCfg BrokerConfig) types.PluginProvider {
	be := &BrokerExt{Broker: broker, BrokerConfig: brokerCfg}
	return newPluginProviderClient(be, conn)
}

func (p *pluginProviderClient) ChainReader() types.ChainReader {
	return p.chainReader
}

func (p *pluginProviderClient) Codec() types.Codec {
	return p.codec
}
