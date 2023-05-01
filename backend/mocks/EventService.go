// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	models "github.com/VolunteerOne/volunteer-one-app/backend/models"
)

// EventService is an autogenerated mock type for the EventService type
type EventService struct {
	mock.Mock
}

// Bind provides a mock function with given fields: _a0, _a1
func (_m *EventService) Bind(_a0 *gin.Context, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gin.Context, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateEvent provides a mock function with given fields: _a0
func (_m *EventService) CreateEvent(_a0 models.Event) (models.Event, error) {
	ret := _m.Called(_a0)

	var r0 models.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Event) (models.Event, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(models.Event) models.Event); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	if rf, ok := ret.Get(1).(func(models.Event) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEvent provides a mock function with given fields: _a0
func (_m *EventService) DeleteEvent(_a0 models.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEventById provides a mock function with given fields: _a0
func (_m *EventService) GetEventById(_a0 string) (models.Event, error) {
	ret := _m.Called(_a0)

	var r0 models.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Event, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) models.Event); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvents provides a mock function with given fields:
func (_m *EventService) GetEvents() ([]models.Event, error) {
	ret := _m.Called()

	var r0 []models.Event
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Event, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEvent provides a mock function with given fields: _a0
func (_m *EventService) UpdateEvent(_a0 models.Event) (models.Event, error) {
	ret := _m.Called(_a0)

	var r0 models.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Event) (models.Event, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(models.Event) models.Event); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	if rf, ok := ret.Get(1).(func(models.Event) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEventService interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventService creates a new instance of EventService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventService(t mockConstructorTestingTNewEventService) *EventService {
	mock := &EventService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
