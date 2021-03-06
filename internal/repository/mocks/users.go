// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks_repository

import (
	context "context"

	domain "github.com/paw1a/ecommerce-api/internal/domain"
	dto "github.com/paw1a/ecommerce-api/internal/domain/dto"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	testing "testing"
)

// Users is an autogenerated mock type for the Users type
type Users struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *Users) Create(ctx context.Context, user domain.User) (domain.User, error) {
	ret := _m.Called(ctx, user)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) domain.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, userID
func (_m *Users) Delete(ctx context.Context, userID primitive.ObjectID) error {
	ret := _m.Called(ctx, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx
func (_m *Users) FindAll(ctx context.Context) ([]domain.User, error) {
	ret := _m.Called(ctx)

	var r0 []domain.User
	if rf, ok := ret.Get(0).(func(context.Context) []domain.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByCredentials provides a mock function with given fields: ctx, email, password
func (_m *Users) FindByCredentials(ctx context.Context, email string, password string) (domain.User, error) {
	ret := _m.Called(ctx, email, password)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) domain.User); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, userID
func (_m *Users) FindByID(ctx context.Context, userID primitive.ObjectID) (domain.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) domain.User); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserInfo provides a mock function with given fields: ctx, userID
func (_m *Users) FindUserInfo(ctx context.Context, userID primitive.ObjectID) (domain.UserInfo, error) {
	ret := _m.Called(ctx, userID)

	var r0 domain.UserInfo
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) domain.UserInfo); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(domain.UserInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, userInput, userID
func (_m *Users) Update(ctx context.Context, userInput dto.UpdateUserInput, userID primitive.ObjectID) (domain.User, error) {
	ret := _m.Called(ctx, userInput, userID)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, dto.UpdateUserInput, primitive.ObjectID) domain.User); ok {
		r0 = rf(ctx, userInput, userID)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.UpdateUserInput, primitive.ObjectID) error); ok {
		r1 = rf(ctx, userInput, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsers creates a new instance of Users. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsers(t testing.TB) *Users {
	mock := &Users{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
