// Code generated by mockery. DO NOT EDIT.

package usecase

import (
	context "context"

	usecase "github.com/sesaquecruz/go-payment-processor/internal/application/usecase"
	mock "github.com/stretchr/testify/mock"
)

// ProcessPaymentMock is an autogenerated mock type for the ProcessPayment type
type ProcessPaymentMock struct {
	mock.Mock
}

type ProcessPaymentMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ProcessPaymentMock) EXPECT() *ProcessPaymentMock_Expecter {
	return &ProcessPaymentMock_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, input
func (_m *ProcessPaymentMock) Execute(ctx context.Context, input usecase.ProcessPaymentInput) (*usecase.ProcessPaymentOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *usecase.ProcessPaymentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, usecase.ProcessPaymentInput) (*usecase.ProcessPaymentOutput, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, usecase.ProcessPaymentInput) *usecase.ProcessPaymentOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usecase.ProcessPaymentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, usecase.ProcessPaymentInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessPaymentMock_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type ProcessPaymentMock_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - input usecase.ProcessPaymentInput
func (_e *ProcessPaymentMock_Expecter) Execute(ctx interface{}, input interface{}) *ProcessPaymentMock_Execute_Call {
	return &ProcessPaymentMock_Execute_Call{Call: _e.mock.On("Execute", ctx, input)}
}

func (_c *ProcessPaymentMock_Execute_Call) Run(run func(ctx context.Context, input usecase.ProcessPaymentInput)) *ProcessPaymentMock_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(usecase.ProcessPaymentInput))
	})
	return _c
}

func (_c *ProcessPaymentMock_Execute_Call) Return(_a0 *usecase.ProcessPaymentOutput, _a1 error) *ProcessPaymentMock_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProcessPaymentMock_Execute_Call) RunAndReturn(run func(context.Context, usecase.ProcessPaymentInput) (*usecase.ProcessPaymentOutput, error)) *ProcessPaymentMock_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewProcessPaymentMock creates a new instance of ProcessPaymentMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProcessPaymentMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProcessPaymentMock {
	mock := &ProcessPaymentMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}