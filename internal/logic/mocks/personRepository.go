// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/polnaya-katuxa/ds-lab-01/internal/models"
)

// PersonRepository is an autogenerated mock type for the personRepository type
type PersonRepository struct {
	mock.Mock
}

type PersonRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *PersonRepository) EXPECT() *PersonRepository_Expecter {
	return &PersonRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *PersonRepository) Create(_a0 context.Context, _a1 models.Person) (*models.Person, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Person) (*models.Person, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Person) *models.Person); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Person) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type PersonRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 models.Person
func (_e *PersonRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *PersonRepository_Create_Call {
	return &PersonRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *PersonRepository_Create_Call) Run(run func(_a0 context.Context, _a1 models.Person)) *PersonRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Person))
	})
	return _c
}

func (_c *PersonRepository_Create_Call) Return(_a0 *models.Person, _a1 error) *PersonRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersonRepository_Create_Call) RunAndReturn(run func(context.Context, models.Person) (*models.Person, error)) *PersonRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *PersonRepository) Delete(_a0 context.Context, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersonRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type PersonRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int
func (_e *PersonRepository_Expecter) Delete(_a0 interface{}, _a1 interface{}) *PersonRepository_Delete_Call {
	return &PersonRepository_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *PersonRepository_Delete_Call) Run(run func(_a0 context.Context, _a1 int)) *PersonRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *PersonRepository_Delete_Call) Return(_a0 error) *PersonRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PersonRepository_Delete_Call) RunAndReturn(run func(context.Context, int) error) *PersonRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Edit provides a mock function with given fields: _a0, _a1
func (_m *PersonRepository) Edit(_a0 context.Context, _a1 models.Person) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Edit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Person) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersonRepository_Edit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Edit'
type PersonRepository_Edit_Call struct {
	*mock.Call
}

// Edit is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 models.Person
func (_e *PersonRepository_Expecter) Edit(_a0 interface{}, _a1 interface{}) *PersonRepository_Edit_Call {
	return &PersonRepository_Edit_Call{Call: _e.mock.On("Edit", _a0, _a1)}
}

func (_c *PersonRepository_Edit_Call) Run(run func(_a0 context.Context, _a1 models.Person)) *PersonRepository_Edit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Person))
	})
	return _c
}

func (_c *PersonRepository_Edit_Call) Return(_a0 error) *PersonRepository_Edit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PersonRepository_Edit_Call) RunAndReturn(run func(context.Context, models.Person) error) *PersonRepository_Edit_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *PersonRepository) Get(_a0 context.Context, _a1 int) (*models.Person, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*models.Person, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Person); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type PersonRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int
func (_e *PersonRepository_Expecter) Get(_a0 interface{}, _a1 interface{}) *PersonRepository_Get_Call {
	return &PersonRepository_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *PersonRepository_Get_Call) Run(run func(_a0 context.Context, _a1 int)) *PersonRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *PersonRepository_Get_Call) Return(_a0 *models.Person, _a1 error) *PersonRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersonRepository_Get_Call) RunAndReturn(run func(context.Context, int) (*models.Person, error)) *PersonRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: _a0
func (_m *PersonRepository) GetAll(_a0 context.Context) ([]models.Person, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []models.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Person, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Person); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonRepository_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type PersonRepository_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *PersonRepository_Expecter) GetAll(_a0 interface{}) *PersonRepository_GetAll_Call {
	return &PersonRepository_GetAll_Call{Call: _e.mock.On("GetAll", _a0)}
}

func (_c *PersonRepository_GetAll_Call) Run(run func(_a0 context.Context)) *PersonRepository_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *PersonRepository_GetAll_Call) Return(_a0 []models.Person, _a1 error) *PersonRepository_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersonRepository_GetAll_Call) RunAndReturn(run func(context.Context) ([]models.Person, error)) *PersonRepository_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// NewPersonRepository creates a new instance of PersonRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPersonRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PersonRepository {
	mock := &PersonRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
