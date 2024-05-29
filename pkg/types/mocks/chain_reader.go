// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	query "github.com/smartcontractkit/chainlink-common/pkg/types/query"
	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink-common/pkg/types"
)

// ChainReader is an autogenerated mock type for the ChainReader type
type ChainReader struct {
	mock.Mock
}

// Bind provides a mock function with given fields: ctx, bindings
func (_m *ChainReader) Bind(ctx context.Context, bindings []types.BoundContract) error {
	ret := _m.Called(ctx, bindings)

	if len(ret) == 0 {
		panic("no return value specified for Bind")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.BoundContract) error); ok {
		r0 = rf(ctx, bindings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *ChainReader) Close() error {
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

// GetLatestValue provides a mock function with given fields: ctx, contractName, method, params, returnVal
func (_m *ChainReader) GetLatestValue(ctx context.Context, contractName string, method string, params interface{}, returnVal interface{}) error {
	ret := _m.Called(ctx, contractName, method, params, returnVal)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestValue")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}, interface{}) error); ok {
		r0 = rf(ctx, contractName, method, params, returnVal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HealthReport provides a mock function with given fields:
func (_m *ChainReader) HealthReport() map[string]error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthReport")
	}

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *ChainReader) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// QueryKey provides a mock function with given fields: ctx, contractName, filter, limitAndSort, sequenceDataType
func (_m *ChainReader) QueryKey(ctx context.Context, contractName string, filter query.KeyFilter, limitAndSort query.LimitAndSort, sequenceDataType interface{}) ([]types.Sequence, error) {
	ret := _m.Called(ctx, contractName, filter, limitAndSort, sequenceDataType)

	if len(ret) == 0 {
		panic("no return value specified for QueryKey")
	}

	var r0 []types.Sequence
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, query.KeyFilter, query.LimitAndSort, interface{}) ([]types.Sequence, error)); ok {
		return rf(ctx, contractName, filter, limitAndSort, sequenceDataType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, query.KeyFilter, query.LimitAndSort, interface{}) []types.Sequence); ok {
		r0 = rf(ctx, contractName, filter, limitAndSort, sequenceDataType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Sequence)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, query.KeyFilter, query.LimitAndSort, interface{}) error); ok {
		r1 = rf(ctx, contractName, filter, limitAndSort, sequenceDataType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ready provides a mock function with given fields:
func (_m *ChainReader) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *ChainReader) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewChainReader creates a new instance of ChainReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChainReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChainReader {
	mock := &ChainReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
