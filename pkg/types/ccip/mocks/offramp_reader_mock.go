// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	ccip "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	mock "github.com/stretchr/testify/mock"
)

// OffRampReader is an autogenerated mock type for the OffRampReader type
type OffRampReader struct {
	mock.Mock
}

// Address provides a mock function with given fields: ctx
func (_m *OffRampReader) Address(ctx context.Context) (ccip.Address, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 ccip.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.Address, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.Address); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccip.Address)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChangeConfig provides a mock function with given fields: ctx, onchainConfig, offchainConfig
func (_m *OffRampReader) ChangeConfig(ctx context.Context, onchainConfig []byte, offchainConfig []byte) (ccip.Address, ccip.Address, error) {
	ret := _m.Called(ctx, onchainConfig, offchainConfig)

	if len(ret) == 0 {
		panic("no return value specified for ChangeConfig")
	}

	var r0 ccip.Address
	var r1 ccip.Address
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, []byte) (ccip.Address, ccip.Address, error)); ok {
		return rf(ctx, onchainConfig, offchainConfig)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, []byte) ccip.Address); ok {
		r0 = rf(ctx, onchainConfig, offchainConfig)
	} else {
		r0 = ret.Get(0).(ccip.Address)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, []byte) ccip.Address); ok {
		r1 = rf(ctx, onchainConfig, offchainConfig)
	} else {
		r1 = ret.Get(1).(ccip.Address)
	}

	if rf, ok := ret.Get(2).(func(context.Context, []byte, []byte) error); ok {
		r2 = rf(ctx, onchainConfig, offchainConfig)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Close provides a mock function with given fields:
func (_m *OffRampReader) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CurrentRateLimiterState provides a mock function with given fields: ctx
func (_m *OffRampReader) CurrentRateLimiterState(ctx context.Context) (ccip.TokenBucketRateLimit, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CurrentRateLimiterState")
	}

	var r0 ccip.TokenBucketRateLimit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.TokenBucketRateLimit, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.TokenBucketRateLimit); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccip.TokenBucketRateLimit)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeExecutionReport provides a mock function with given fields: ctx, report
func (_m *OffRampReader) DecodeExecutionReport(ctx context.Context, report []byte) (ccip.ExecReport, error) {
	ret := _m.Called(ctx, report)

	if len(ret) == 0 {
		panic("no return value specified for DecodeExecutionReport")
	}

	var r0 ccip.ExecReport
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) (ccip.ExecReport, error)); ok {
		return rf(ctx, report)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte) ccip.ExecReport); ok {
		r0 = rf(ctx, report)
	} else {
		r0 = ret.Get(0).(ccip.ExecReport)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, report)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncodeExecutionReport provides a mock function with given fields: ctx, report
func (_m *OffRampReader) EncodeExecutionReport(ctx context.Context, report ccip.ExecReport) ([]byte, error) {
	ret := _m.Called(ctx, report)

	if len(ret) == 0 {
		panic("no return value specified for EncodeExecutionReport")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ccip.ExecReport) ([]byte, error)); ok {
		return rf(ctx, report)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ccip.ExecReport) []byte); ok {
		r0 = rf(ctx, report)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ccip.ExecReport) error); ok {
		r1 = rf(ctx, report)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GasPriceEstimator provides a mock function with given fields: ctx
func (_m *OffRampReader) GasPriceEstimator(ctx context.Context) (ccip.GasPriceEstimatorExec, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GasPriceEstimator")
	}

	var r0 ccip.GasPriceEstimatorExec
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.GasPriceEstimatorExec, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.GasPriceEstimatorExec); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ccip.GasPriceEstimatorExec)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExecutionState provides a mock function with given fields: ctx, sequenceNumber
func (_m *OffRampReader) GetExecutionState(ctx context.Context, sequenceNumber uint64) (uint8, error) {
	ret := _m.Called(ctx, sequenceNumber)

	if len(ret) == 0 {
		panic("no return value specified for GetExecutionState")
	}

	var r0 uint8
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (uint8, error)); ok {
		return rf(ctx, sequenceNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) uint8); ok {
		r0 = rf(ctx, sequenceNumber)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, sequenceNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExecutionStateChangesBetweenSeqNums provides a mock function with given fields: ctx, seqNumMin, seqNumMax, confirmations
func (_m *OffRampReader) GetExecutionStateChangesBetweenSeqNums(ctx context.Context, seqNumMin uint64, seqNumMax uint64, confirmations int) ([]ccip.ExecutionStateChangedWithTxMeta, error) {
	ret := _m.Called(ctx, seqNumMin, seqNumMax, confirmations)

	if len(ret) == 0 {
		panic("no return value specified for GetExecutionStateChangesBetweenSeqNums")
	}

	var r0 []ccip.ExecutionStateChangedWithTxMeta
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, int) ([]ccip.ExecutionStateChangedWithTxMeta, error)); ok {
		return rf(ctx, seqNumMin, seqNumMax, confirmations)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, int) []ccip.ExecutionStateChangedWithTxMeta); ok {
		r0 = rf(ctx, seqNumMin, seqNumMax, confirmations)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccip.ExecutionStateChangedWithTxMeta)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, int) error); ok {
		r1 = rf(ctx, seqNumMin, seqNumMax, confirmations)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExecutionStateChangesForSeqNums provides a mock function with given fields: ctx, seqNums, confirmations
func (_m *OffRampReader) GetExecutionStateChangesForSeqNums(ctx context.Context, seqNums []ccip.SequenceNumberRange, confirmations int) ([]ccip.ExecutionStateChangedWithTxMeta, error) {
	ret := _m.Called(ctx, seqNums, confirmations)

	if len(ret) == 0 {
		panic("no return value specified for GetExecutionStateChangesForSeqNums")
	}

	var r0 []ccip.ExecutionStateChangedWithTxMeta
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []ccip.SequenceNumberRange, int) ([]ccip.ExecutionStateChangedWithTxMeta, error)); ok {
		return rf(ctx, seqNums, confirmations)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []ccip.SequenceNumberRange, int) []ccip.ExecutionStateChangedWithTxMeta); ok {
		r0 = rf(ctx, seqNums, confirmations)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccip.ExecutionStateChangedWithTxMeta)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []ccip.SequenceNumberRange, int) error); ok {
		r1 = rf(ctx, seqNums, confirmations)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRouter provides a mock function with given fields: ctx
func (_m *OffRampReader) GetRouter(ctx context.Context) (ccip.Address, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetRouter")
	}

	var r0 ccip.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.Address, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.Address); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccip.Address)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSourceToDestTokensMapping provides a mock function with given fields: ctx
func (_m *OffRampReader) GetSourceToDestTokensMapping(ctx context.Context) (map[ccip.Address]ccip.Address, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetSourceToDestTokensMapping")
	}

	var r0 map[ccip.Address]ccip.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[ccip.Address]ccip.Address, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[ccip.Address]ccip.Address); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[ccip.Address]ccip.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStaticConfig provides a mock function with given fields: ctx
func (_m *OffRampReader) GetStaticConfig(ctx context.Context) (ccip.OffRampStaticConfig, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetStaticConfig")
	}

	var r0 ccip.OffRampStaticConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.OffRampStaticConfig, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.OffRampStaticConfig); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccip.OffRampStaticConfig)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokens provides a mock function with given fields: ctx
func (_m *OffRampReader) GetTokens(ctx context.Context) (ccip.OffRampTokens, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetTokens")
	}

	var r0 ccip.OffRampTokens
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.OffRampTokens, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.OffRampTokens); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccip.OffRampTokens)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSenderNonces provides a mock function with given fields: ctx, senders
func (_m *OffRampReader) ListSenderNonces(ctx context.Context, senders []ccip.Address) (map[ccip.Address]uint64, error) {
	ret := _m.Called(ctx, senders)

	if len(ret) == 0 {
		panic("no return value specified for ListSenderNonces")
	}

	var r0 map[ccip.Address]uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []ccip.Address) (map[ccip.Address]uint64, error)); ok {
		return rf(ctx, senders)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []ccip.Address) map[ccip.Address]uint64); ok {
		r0 = rf(ctx, senders)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[ccip.Address]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []ccip.Address) error); ok {
		r1 = rf(ctx, senders)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OffchainConfig provides a mock function with given fields: ctx
func (_m *OffRampReader) OffchainConfig(ctx context.Context) (ccip.ExecOffchainConfig, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for OffchainConfig")
	}

	var r0 ccip.ExecOffchainConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.ExecOffchainConfig, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.ExecOffchainConfig); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccip.ExecOffchainConfig)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnchainConfig provides a mock function with given fields: ctx
func (_m *OffRampReader) OnchainConfig(ctx context.Context) (ccip.ExecOnchainConfig, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for OnchainConfig")
	}

	var r0 ccip.ExecOnchainConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.ExecOnchainConfig, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.ExecOnchainConfig); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccip.ExecOnchainConfig)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOffRampReader creates a new instance of OffRampReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOffRampReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *OffRampReader {
	mock := &OffRampReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
