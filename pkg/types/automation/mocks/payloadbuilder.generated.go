// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	automation "github.com/smartcontractkit/chainlink-common/pkg/types/automation"

	mock "github.com/stretchr/testify/mock"
)

// MockPayloadBuilder is an autogenerated mock type for the PayloadBuilder type
type MockPayloadBuilder struct {
	mock.Mock
}

// BuildPayloads provides a mock function with given fields: _a0, _a1
func (_m *MockPayloadBuilder) BuildPayloads(_a0 context.Context, _a1 ...automation.CoordinatedBlockProposal) ([]automation.UpkeepPayload, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BuildPayloads")
	}

	var r0 []automation.UpkeepPayload
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...automation.CoordinatedBlockProposal) ([]automation.UpkeepPayload, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...automation.CoordinatedBlockProposal) []automation.UpkeepPayload); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]automation.UpkeepPayload)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...automation.CoordinatedBlockProposal) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockPayloadBuilder creates a new instance of MockPayloadBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPayloadBuilder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPayloadBuilder {
	mock := &MockPayloadBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
