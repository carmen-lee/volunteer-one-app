// Code generated by mockery v2.26.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// UsersController is an autogenerated mock type for the UsersController type
type UsersController struct {
	mock.Mock
}

// Create provides a mock function with given fields: c
func (_m *UsersController) Create(c *gin.Context) {
	_m.Called(c)
}

// Delete provides a mock function with given fields: c
func (_m *UsersController) Delete(c *gin.Context) {
	_m.Called(c)
}

// One provides a mock function with given fields: c
func (_m *UsersController) One(c *gin.Context) {
	_m.Called(c)
}

// Update provides a mock function with given fields: c
func (_m *UsersController) Update(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewUsersController interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsersController creates a new instance of UsersController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsersController(t mockConstructorTestingTNewUsersController) *UsersController {
	mock := &UsersController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
