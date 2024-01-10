package test

import (
	"bytes"
	"context"
	"fmt"
	"reflect"

	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/loop/internal"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type MockConn struct {
	grpc.ClientConnInterface
}

const ReportingPluginWithMedianProviderName = "reporting-plugin-with-median-provider"

type staticReportingPluginWithMedianProvider struct {
	staticService
}

func NewStaticReportingPluginWithMedianProvider(lggr logger.Logger) staticReportingPluginWithMedianProvider {
	return staticReportingPluginWithMedianProvider{staticService{lggr: logger.Named(lggr, "staticReportingPluginWithMedianProvider")}}
}

func (s staticReportingPluginWithMedianProvider) HealthReport() map[string]error {
	hp := s.staticService.HealthReport()
	services.CopyHealth(hp, NewStaticMedianProvider(s.lggr).HealthReport())
	services.CopyHealth(hp, newStaticPluginFactory(s.lggr).HealthReport())
	return hp
}

func (s staticReportingPluginWithMedianProvider) ConnToProvider(conn grpc.ClientConnInterface, broker internal.Broker, brokerConfig internal.BrokerConfig) types.MedianProvider {
	return NewStaticMedianProvider(s.lggr)
}

func (s staticReportingPluginWithMedianProvider) NewReportingPluginFactory(ctx context.Context, config types.ReportingPluginServiceConfig, provider types.MedianProvider, pipelineRunner types.PipelineRunnerService, telemetry types.TelemetryClient, errorLog types.ErrorLog) (types.ReportingPluginFactory, error) {
	//TODO validate config
	ocd := provider.OffchainConfigDigester()
	gotDigestPrefix, err := ocd.ConfigDigestPrefix()
	if err != nil {
		return nil, fmt.Errorf("failed to get ConfigDigestPrefix: %w", err)
	}
	if gotDigestPrefix != configDigestPrefix {
		return nil, fmt.Errorf("expected ConfigDigestPrefix %x but got %x", configDigestPrefix, gotDigestPrefix)
	}
	gotDigest, err := ocd.ConfigDigest(contractConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get ConfigDigest: %w", err)
	}
	if gotDigest != configDigest {
		return nil, fmt.Errorf("expected ConfigDigest %x but got %x", configDigest, gotDigest)
	}
	cct := provider.ContractConfigTracker()
	gotBlockHeight, err := cct.LatestBlockHeight(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get LatestBlockHeight: %w", err)
	}
	if gotBlockHeight != blockHeight {
		return nil, fmt.Errorf("expected LatestBlockHeight %d but got %d", blockHeight, gotBlockHeight)
	}
	gotChangedInBlock, gotConfigDigest, err := cct.LatestConfigDetails(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get LatestConfigDetails: %w", err)
	}
	if gotChangedInBlock != changedInBlock {
		return nil, fmt.Errorf("expected changedInBlock %d but got %d", changedInBlock, gotChangedInBlock)
	}
	if gotConfigDigest != configDigest {
		return nil, fmt.Errorf("expected ConfigDigest %s but got %s", configDigest, gotConfigDigest)
	}
	gotContractConfig, err := cct.LatestConfig(ctx, changedInBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to get LatestConfig: %w", err)
	}
	if !reflect.DeepEqual(gotContractConfig, contractConfig) {
		return nil, fmt.Errorf("expected ContractConfig %v but got %v", contractConfig, gotContractConfig)
	}
	ct := provider.ContractTransmitter()
	gotAccount, err := ct.FromAccount()
	if err != nil {
		return nil, fmt.Errorf("failed to get FromAccount: %w", err)
	}
	if gotAccount != account {
		return nil, fmt.Errorf("expectd FromAccount %s but got %s", account, gotAccount)
	}
	gotConfigDigest, gotEpoch, err := ct.LatestConfigDigestAndEpoch(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get LatestConfigDigestAndEpoch: %w", err)
	}
	if gotConfigDigest != configDigest {
		return nil, fmt.Errorf("expected ConfigDigest %s but got %s", configDigest, gotConfigDigest)
	}
	if gotEpoch != epoch {
		return nil, fmt.Errorf("expected Epoch %d but got %d", epoch, gotEpoch)
	}
	err = ct.Transmit(ctx, reportContext, report, sigs)
	if err != nil {
		return nil, fmt.Errorf("failed to Transmit")
	}
	rc := provider.ReportCodec()
	gotReport, err := rc.BuildReport(pobs)
	if err != nil {
		return nil, fmt.Errorf("failed to BuildReport: %w", err)
	}
	if !bytes.Equal(gotReport, report) {
		return nil, fmt.Errorf("expected Report %x but got %x", report, gotReport)
	}
	gotMedianValue, err := rc.MedianFromReport(report)
	if err != nil {
		return nil, fmt.Errorf("failed to get MedianFromReport: %w", err)
	}
	if medianValue.Cmp(gotMedianValue) != 0 {
		return nil, fmt.Errorf("expected MedianValue %s but got %s", medianValue, gotMedianValue)
	}
	gotMax, err := rc.MaxReportLength(n)
	if err != nil {
		return nil, fmt.Errorf("failed to get MaxReportLength: %w", err)
	}
	if gotMax != max {
		return nil, fmt.Errorf("expected MaxReportLength %d but got %d", max, gotMax)
	}
	mc := provider.MedianContract()
	gotConfigDigest, gotEpoch, gotRound, err := mc.LatestRoundRequested(ctx, lookbackDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to get LatestRoundRequested: %w", err)
	}
	if gotConfigDigest != configDigest {
		return nil, fmt.Errorf("expected ConfigDigest %s but got %s", configDigest, gotConfigDigest)
	}
	if gotEpoch != epoch {
		return nil, fmt.Errorf("expected Epoch %d but got %d", epoch, gotEpoch)
	}
	if gotRound != round {
		return nil, fmt.Errorf("expected Round %d but got %d", round, gotRound)
	}
	gotConfigDigest, gotEpoch, gotRound, gotLatestAnswer, gotLatestTimestamp, err := mc.LatestTransmissionDetails(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get LatestTransmissionDetails: %w", err)
	}
	if gotConfigDigest != configDigest {
		return nil, fmt.Errorf("expected ConfigDigest %s but got %s", configDigest, gotConfigDigest)
	}
	if gotEpoch != epoch {
		return nil, fmt.Errorf("expected Epoch %d but got %d", epoch, gotEpoch)
	}
	if gotRound != round {
		return nil, fmt.Errorf("expected Round %d but got %d", round, gotRound)
	}
	if latestAnswer.Cmp(gotLatestAnswer) != 0 {
		return nil, fmt.Errorf("expected LatestAnswer %s but got %s", latestAnswer, gotLatestAnswer)
	}
	if !gotLatestTimestamp.Equal(latestTimestamp) {
		return nil, fmt.Errorf("expected LatestTimestamp %s but got %s", latestTimestamp, gotLatestTimestamp)
	}
	occ := provider.OnchainConfigCodec()
	gotEncoded, err := occ.Encode(onchainConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to Encode: %w", err)
	}
	if !bytes.Equal(gotEncoded, encoded) {
		return nil, fmt.Errorf("expected Encoded %s but got %s", encoded, gotEncoded)
	}
	gotDecoded, err := occ.Decode(encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to Decode: %w", err)
	}
	if !reflect.DeepEqual(gotDecoded, onchainConfig) {
		return nil, fmt.Errorf("expected OnchainConfig %s but got %s", onchainConfig, gotDecoded)
	}
	//TODO validate pipelineRunner
	if err2 := errorLog.SaveError(ctx, errMsg); err2 != nil {
		return nil, fmt.Errorf("failed to save error: %w", err2)
	}
	endpoint, err := telemetry.NewEndpoint(ctx, network, chainID, contractID, telemType)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate endpoint: %w", err)
	}
	err = endpoint.SendLog(ctx, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to send log: %w", err)
	}
	return newStaticPluginFactory(s.lggr), nil
}

type staticReportingPluginWithPluginProvider struct {
	staticService
}

func NewStaticReportingPluginWithPluginProvider(lggr logger.Logger) staticReportingPluginWithPluginProvider {
	return staticReportingPluginWithPluginProvider{staticService{lggr: logger.Named(lggr, "staticReportingPluginWithPluginProvider")}}
}

func (s staticReportingPluginWithPluginProvider) HealthReport() map[string]error {
	hp := s.staticService.HealthReport()
	services.CopyHealth(hp, NewStaticPluginProvider(s.lggr).HealthReport())
	services.CopyHealth(hp, newStaticPluginFactory(s.lggr).HealthReport())
	return hp
}

func (s staticReportingPluginWithPluginProvider) ConnToProvider(conn grpc.ClientConnInterface, broker internal.Broker, brokerConfig internal.BrokerConfig) types.PluginProvider {
	return NewStaticPluginProvider(s.lggr)
}

func (s staticReportingPluginWithPluginProvider) NewReportingPluginFactory(ctx context.Context, config types.ReportingPluginServiceConfig, provider types.PluginProvider, pipelineRunner types.PipelineRunnerService, telemetry types.TelemetryClient, errorLog types.ErrorLog) (types.ReportingPluginFactory, error) {
	ocd := provider.OffchainConfigDigester()
	gotDigestPrefix, err := ocd.ConfigDigestPrefix()
	if err != nil {
		return nil, fmt.Errorf("failed to get ConfigDigestPrefix: %w", err)
	}
	if gotDigestPrefix != configDigestPrefix {
		return nil, fmt.Errorf("expected ConfigDigestPrefix %x but got %x", configDigestPrefix, gotDigestPrefix)
	}
	gotDigest, err := ocd.ConfigDigest(contractConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get ConfigDigest: %w", err)
	}
	if gotDigest != configDigest {
		return nil, fmt.Errorf("expected ConfigDigest %x but got %x", configDigest, gotDigest)
	}
	cct := provider.ContractConfigTracker()
	gotBlockHeight, err := cct.LatestBlockHeight(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get LatestBlockHeight: %w", err)
	}
	if gotBlockHeight != blockHeight {
		return nil, fmt.Errorf("expected LatestBlockHeight %d but got %d", blockHeight, gotBlockHeight)
	}
	gotChangedInBlock, gotConfigDigest, err := cct.LatestConfigDetails(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get LatestConfigDetails: %w", err)
	}
	if gotChangedInBlock != changedInBlock {
		return nil, fmt.Errorf("expected changedInBlock %d but got %d", changedInBlock, gotChangedInBlock)
	}
	if gotConfigDigest != configDigest {
		return nil, fmt.Errorf("expected ConfigDigest %s but got %s", configDigest, gotConfigDigest)
	}
	gotContractConfig, err := cct.LatestConfig(ctx, changedInBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to get LatestConfig: %w", err)
	}
	if !reflect.DeepEqual(gotContractConfig, contractConfig) {
		return nil, fmt.Errorf("expected ContractConfig %v but got %v", contractConfig, gotContractConfig)
	}
	ct := provider.ContractTransmitter()
	gotAccount, err := ct.FromAccount()
	if err != nil {
		return nil, fmt.Errorf("failed to get FromAccount: %w", err)
	}
	if gotAccount != account {
		return nil, fmt.Errorf("expectd FromAccount %s but got %s", account, gotAccount)
	}
	gotConfigDigest, gotEpoch, err := ct.LatestConfigDigestAndEpoch(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get LatestConfigDigestAndEpoch: %w", err)
	}
	if gotConfigDigest != configDigest {
		return nil, fmt.Errorf("expected ConfigDigest %s but got %s", configDigest, gotConfigDigest)
	}
	if gotEpoch != epoch {
		return nil, fmt.Errorf("expected Epoch %d but got %d", epoch, gotEpoch)
	}
	err = ct.Transmit(ctx, reportContext, report, sigs)
	if err != nil {
		return nil, fmt.Errorf("failed to Transmit")
	}
	if err2 := errorLog.SaveError(ctx, errMsg); err2 != nil {
		return nil, fmt.Errorf("failed to save error: %w", err2)
	}
	tr, err := pipelineRunner.ExecuteRun(ctx, spec, vars, options)
	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}
	if !reflect.DeepEqual(tr, taskResults) {
		return nil, fmt.Errorf("expected TaskResults %+v but got %+v", taskResults, tr)
	}
	endpoint, err := telemetry.NewEndpoint(ctx, network, chainID, contractID, telemType)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate endpoint: %w", err)
	}
	err = endpoint.SendLog(ctx, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to send log: %w", err)
	}
	return newStaticPluginFactory(s.lggr), nil
}
