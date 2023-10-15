// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/sesaquecruz/go-payment-processor/internal/application/entity"
	mock "github.com/stretchr/testify/mock"
)

// PaymentMock is an autogenerated mock type for the Payment type
type PaymentMock struct {
	mock.Mock
}

type PaymentMock_Expecter struct {
	mock *mock.Mock
}

func (_m *PaymentMock) EXPECT() *PaymentMock_Expecter {
	return &PaymentMock_Expecter{mock: &_m.Mock}
}

// Process provides a mock function with given fields: ctx, transaction
func (_m *PaymentMock) Process(ctx context.Context, transaction entity.Transaction) (*entity.Payment, error) {
	ret := _m.Called(ctx, transaction)

	var r0 *entity.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Transaction) (*entity.Payment, error)); ok {
		return rf(ctx, transaction)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Transaction) *entity.Payment); ok {
		r0 = rf(ctx, transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Payment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Transaction) error); ok {
		r1 = rf(ctx, transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PaymentMock_Process_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Process'
type PaymentMock_Process_Call struct {
	*mock.Call
}

// Process is a helper method to define mock.On call
//   - ctx context.Context
//   - transaction entity.Transaction
func (_e *PaymentMock_Expecter) Process(ctx interface{}, transaction interface{}) *PaymentMock_Process_Call {
	return &PaymentMock_Process_Call{Call: _e.mock.On("Process", ctx, transaction)}
}

func (_c *PaymentMock_Process_Call) Run(run func(ctx context.Context, transaction entity.Transaction)) *PaymentMock_Process_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.Transaction))
	})
	return _c
}

func (_c *PaymentMock_Process_Call) Return(_a0 *entity.Payment, _a1 error) *PaymentMock_Process_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PaymentMock_Process_Call) RunAndReturn(run func(context.Context, entity.Transaction) (*entity.Payment, error)) *PaymentMock_Process_Call {
	_c.Call.Return(run)
	return _c
}

// NewPaymentMock creates a new instance of PaymentMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentMock {
	mock := &PaymentMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}