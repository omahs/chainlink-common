package loop_test

import (
	"testing"

	"github.com/hashicorp/go-plugin"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal/test"
)

func TestPluginRelayer(t *testing.T) {
	t.Parallel()

	lggr := logger.Test(t)
	stopCh := newStopCh(t)
	test.PluginTest(t, loop.PluginRelayerName, &loop.GRPCPluginRelayer{PluginServer: test.NewStaticPluginRelayer(lggr), BrokerConfig: loop.BrokerConfig{Logger: logger.Test(t), StopCh: stopCh}}, test.RunPluginRelayer)
}

func TestPluginRelayerExec(t *testing.T) {
	t.Parallel()
	stopCh := newStopCh(t)

	pr := newPluginRelayerExec(t, stopCh)

	test.RunPluginRelayer(t, pr)
}

func newPluginRelayerExec(t *testing.T, stopCh <-chan struct{}) loop.PluginRelayer {
	relayer := loop.GRPCPluginRelayer{BrokerConfig: loop.BrokerConfig{Logger: logger.Test(t), StopCh: stopCh}}
	cc := relayer.ClientConfig()
	cc.Cmd = NewHelperProcessCommand(loop.PluginRelayerName)
	c := plugin.NewClient(cc)
	t.Cleanup(c.Kill)
	client, err := c.Client()
	require.NoError(t, err)
	t.Cleanup(func() { _ = client.Close() })
	require.NoError(t, client.Ping())
	i, err := client.Dispense(loop.PluginRelayerName)
	require.NoError(t, err)
	return i.(loop.PluginRelayer)
}
