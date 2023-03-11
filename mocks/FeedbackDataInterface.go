// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	feedbacks "immersiveApp/features/feedbacks"

	mock "github.com/stretchr/testify/mock"
)

// FeedbackDataInterface is an autogenerated mock type for the FeedbackDataInterface type
type FeedbackDataInterface struct {
	mock.Mock
}

// Edit provides a mock function with given fields: feedbackEntity, id
func (_m *FeedbackDataInterface) Edit(feedbackEntity feedbacks.FeedbackEntity, id uint) error {
	ret := _m.Called(feedbackEntity, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(feedbacks.FeedbackEntity, uint) error); ok {
		r0 = rf(feedbackEntity, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: feedbackEntity
func (_m *FeedbackDataInterface) Insert(feedbackEntity feedbacks.FeedbackEntity) (uint, error) {
	ret := _m.Called(feedbackEntity)

	var r0 uint
	if rf, ok := ret.Get(0).(func(feedbacks.FeedbackEntity) uint); ok {
		r0 = rf(feedbackEntity)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(feedbacks.FeedbackEntity) error); ok {
		r1 = rf(feedbackEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: id
func (_m *FeedbackDataInterface) Remove(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields:
func (_m *FeedbackDataInterface) SelectAll() ([]feedbacks.FeedbackEntity, error) {
	ret := _m.Called()

	var r0 []feedbacks.FeedbackEntity
	if rf, ok := ret.Get(0).(func() []feedbacks.FeedbackEntity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]feedbacks.FeedbackEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: id
func (_m *FeedbackDataInterface) SelectById(id uint) (feedbacks.FeedbackEntity, error) {
	ret := _m.Called(id)

	var r0 feedbacks.FeedbackEntity
	if rf, ok := ret.Get(0).(func(uint) feedbacks.FeedbackEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(feedbacks.FeedbackEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFeedbackDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewFeedbackDataInterface creates a new instance of FeedbackDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFeedbackDataInterface(t mockConstructorTestingTNewFeedbackDataInterface) *FeedbackDataInterface {
	mock := &FeedbackDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
