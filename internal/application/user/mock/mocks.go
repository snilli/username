// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package user_test

import (
	"context"
	"userman/internal/application/common/service"
	user0 "userman/internal/application/user"
	"userman/internal/domain/user"

	mock "github.com/stretchr/testify/mock"
)

// NewMockUserService creates a new instance of MockUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserService {
	mock := &MockUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockUserService is an autogenerated mock type for the UserService type
type MockUserService struct {
	mock.Mock
}

type MockUserService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserService) EXPECT() *MockUserService_Expecter {
	return &MockUserService_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function for the type MockUserService
func (_mock *MockUserService) CreateUser(ctx context.Context, name string, email string, password string) (*user.User, error) {
	ret := _mock.Called(ctx, name, email, password)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *user.User
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string, string) (*user.User, error)); ok {
		return returnFunc(ctx, name, email, password)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string, string) *user.User); ok {
		r0 = returnFunc(ctx, name, email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = returnFunc(ctx, name, email, password)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockUserService_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockUserService_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx
//   - name
//   - email
//   - password
func (_e *MockUserService_Expecter) CreateUser(ctx interface{}, name interface{}, email interface{}, password interface{}) *MockUserService_CreateUser_Call {
	return &MockUserService_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, name, email, password)}
}

func (_c *MockUserService_CreateUser_Call) Run(run func(ctx context.Context, name string, email string, password string)) *MockUserService_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockUserService_CreateUser_Call) Return(user1 *user.User, err error) *MockUserService_CreateUser_Call {
	_c.Call.Return(user1, err)
	return _c
}

func (_c *MockUserService_CreateUser_Call) RunAndReturn(run func(ctx context.Context, name string, email string, password string) (*user.User, error)) *MockUserService_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUserByID provides a mock function for the type MockUserService
func (_mock *MockUserService) DeleteUserByID(ctx context.Context, id string) (*user.User, error) {
	ret := _mock.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserByID")
	}

	var r0 *user.User
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) (*user.User, error)); ok {
		return returnFunc(ctx, id)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) *user.User); ok {
		r0 = returnFunc(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = returnFunc(ctx, id)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockUserService_DeleteUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUserByID'
type MockUserService_DeleteUserByID_Call struct {
	*mock.Call
}

// DeleteUserByID is a helper method to define mock.On call
//   - ctx
//   - id
func (_e *MockUserService_Expecter) DeleteUserByID(ctx interface{}, id interface{}) *MockUserService_DeleteUserByID_Call {
	return &MockUserService_DeleteUserByID_Call{Call: _e.mock.On("DeleteUserByID", ctx, id)}
}

func (_c *MockUserService_DeleteUserByID_Call) Run(run func(ctx context.Context, id string)) *MockUserService_DeleteUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserService_DeleteUserByID_Call) Return(user1 *user.User, err error) *MockUserService_DeleteUserByID_Call {
	_c.Call.Return(user1, err)
	return _c
}

func (_c *MockUserService_DeleteUserByID_Call) RunAndReturn(run func(ctx context.Context, id string) (*user.User, error)) *MockUserService_DeleteUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllUser provides a mock function for the type MockUserService
func (_mock *MockUserService) GetAllUser(ctx context.Context, input *user0.GetAllUserInput) (*common.CursorPaginated[*user.User], error) {
	ret := _mock.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for GetAllUser")
	}

	var r0 *common.CursorPaginated[*user.User]
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *user0.GetAllUserInput) (*common.CursorPaginated[*user.User], error)); ok {
		return returnFunc(ctx, input)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *user0.GetAllUserInput) *common.CursorPaginated[*user.User]); ok {
		r0 = returnFunc(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.CursorPaginated[*user.User])
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *user0.GetAllUserInput) error); ok {
		r1 = returnFunc(ctx, input)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockUserService_GetAllUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllUser'
type MockUserService_GetAllUser_Call struct {
	*mock.Call
}

// GetAllUser is a helper method to define mock.On call
//   - ctx
//   - input
func (_e *MockUserService_Expecter) GetAllUser(ctx interface{}, input interface{}) *MockUserService_GetAllUser_Call {
	return &MockUserService_GetAllUser_Call{Call: _e.mock.On("GetAllUser", ctx, input)}
}

func (_c *MockUserService_GetAllUser_Call) Run(run func(ctx context.Context, input *user0.GetAllUserInput)) *MockUserService_GetAllUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*user0.GetAllUserInput))
	})
	return _c
}

func (_c *MockUserService_GetAllUser_Call) Return(cursorPaginated *common.CursorPaginated[*user.User], err error) *MockUserService_GetAllUser_Call {
	_c.Call.Return(cursorPaginated, err)
	return _c
}

func (_c *MockUserService_GetAllUser_Call) RunAndReturn(run func(ctx context.Context, input *user0.GetAllUserInput) (*common.CursorPaginated[*user.User], error)) *MockUserService_GetAllUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function for the type MockUserService
func (_mock *MockUserService) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	ret := _mock.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *user.User
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) (*user.User, error)); ok {
		return returnFunc(ctx, email)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) *user.User); ok {
		r0 = returnFunc(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = returnFunc(ctx, email)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockUserService_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type MockUserService_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - ctx
//   - email
func (_e *MockUserService_Expecter) GetUserByEmail(ctx interface{}, email interface{}) *MockUserService_GetUserByEmail_Call {
	return &MockUserService_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", ctx, email)}
}

func (_c *MockUserService_GetUserByEmail_Call) Run(run func(ctx context.Context, email string)) *MockUserService_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserService_GetUserByEmail_Call) Return(user1 *user.User, err error) *MockUserService_GetUserByEmail_Call {
	_c.Call.Return(user1, err)
	return _c
}

func (_c *MockUserService_GetUserByEmail_Call) RunAndReturn(run func(ctx context.Context, email string) (*user.User, error)) *MockUserService_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByID provides a mock function for the type MockUserService
func (_mock *MockUserService) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	ret := _mock.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *user.User
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) (*user.User, error)); ok {
		return returnFunc(ctx, id)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) *user.User); ok {
		r0 = returnFunc(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = returnFunc(ctx, id)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockUserService_GetUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByID'
type MockUserService_GetUserByID_Call struct {
	*mock.Call
}

// GetUserByID is a helper method to define mock.On call
//   - ctx
//   - id
func (_e *MockUserService_Expecter) GetUserByID(ctx interface{}, id interface{}) *MockUserService_GetUserByID_Call {
	return &MockUserService_GetUserByID_Call{Call: _e.mock.On("GetUserByID", ctx, id)}
}

func (_c *MockUserService_GetUserByID_Call) Run(run func(ctx context.Context, id string)) *MockUserService_GetUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserService_GetUserByID_Call) Return(user1 *user.User, err error) *MockUserService_GetUserByID_Call {
	_c.Call.Return(user1, err)
	return _c
}

func (_c *MockUserService_GetUserByID_Call) RunAndReturn(run func(ctx context.Context, id string) (*user.User, error)) *MockUserService_GetUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserByID provides a mock function for the type MockUserService
func (_mock *MockUserService) UpdateUserByID(ctx context.Context, id string, name string, email string) (*user.User, error) {
	ret := _mock.Called(ctx, id, name, email)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserByID")
	}

	var r0 *user.User
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string, string) (*user.User, error)); ok {
		return returnFunc(ctx, id, name, email)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, string, string) *user.User); ok {
		r0 = returnFunc(ctx, id, name, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = returnFunc(ctx, id, name, email)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockUserService_UpdateUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserByID'
type MockUserService_UpdateUserByID_Call struct {
	*mock.Call
}

// UpdateUserByID is a helper method to define mock.On call
//   - ctx
//   - id
//   - name
//   - email
func (_e *MockUserService_Expecter) UpdateUserByID(ctx interface{}, id interface{}, name interface{}, email interface{}) *MockUserService_UpdateUserByID_Call {
	return &MockUserService_UpdateUserByID_Call{Call: _e.mock.On("UpdateUserByID", ctx, id, name, email)}
}

func (_c *MockUserService_UpdateUserByID_Call) Run(run func(ctx context.Context, id string, name string, email string)) *MockUserService_UpdateUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockUserService_UpdateUserByID_Call) Return(user1 *user.User, err error) *MockUserService_UpdateUserByID_Call {
	_c.Call.Return(user1, err)
	return _c
}

func (_c *MockUserService_UpdateUserByID_Call) RunAndReturn(run func(ctx context.Context, id string, name string, email string) (*user.User, error)) *MockUserService_UpdateUserByID_Call {
	_c.Call.Return(run)
	return _c
}
