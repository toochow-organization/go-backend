// Code generated by mockery v2.36.0. DO NOT EDIT.

package mock

import (
	context "context"

	apis "github.com/toochow-organization/bego/external/apis"

	mock "github.com/stretchr/testify/mock"
)

// ModelStorage is an autogenerated mock type for the ModelStorage type
type ModelStorage struct {
	mock.Mock
}

type ModelStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *ModelStorage) EXPECT() *ModelStorage_Expecter {
	return &ModelStorage_Expecter{mock: &_m.Mock}
}

// ListModels provides a mock function with given fields: ctx
func (_m *ModelStorage) ListModels(ctx context.Context) ([]apis.Model, error) {
	ret := _m.Called(ctx)

	var r0 []apis.Model
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]apis.Model, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []apis.Model); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]apis.Model)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModelStorage_ListModels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListModels'
type ModelStorage_ListModels_Call struct {
	*mock.Call
}

// ListModels is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ModelStorage_Expecter) ListModels(ctx interface{}) *ModelStorage_ListModels_Call {
	return &ModelStorage_ListModels_Call{Call: _e.mock.On("ListModels", ctx)}
}

func (_c *ModelStorage_ListModels_Call) Run(run func(ctx context.Context)) *ModelStorage_ListModels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ModelStorage_ListModels_Call) Return(_a0 []apis.Model, _a1 error) *ModelStorage_ListModels_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ModelStorage_ListModels_Call) RunAndReturn(run func(context.Context) ([]apis.Model, error)) *ModelStorage_ListModels_Call {
	_c.Call.Return(run)
	return _c
}

// NewModelStorage creates a new instance of ModelStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewModelStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *ModelStorage {
	mock := &ModelStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
