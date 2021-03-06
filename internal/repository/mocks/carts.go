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

// Carts is an autogenerated mock type for the Carts type
type Carts struct {
	mock.Mock
}

// AddCartItem provides a mock function with given fields: ctx, cartItem, cartID
func (_m *Carts) AddCartItem(ctx context.Context, cartItem domain.CartItem, cartID primitive.ObjectID) (domain.CartItem, error) {
	ret := _m.Called(ctx, cartItem, cartID)

	var r0 domain.CartItem
	if rf, ok := ret.Get(0).(func(context.Context, domain.CartItem, primitive.ObjectID) domain.CartItem); ok {
		r0 = rf(ctx, cartItem, cartID)
	} else {
		r0 = ret.Get(0).(domain.CartItem)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.CartItem, primitive.ObjectID) error); ok {
		r1 = rf(ctx, cartItem, cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClearCart provides a mock function with given fields: ctx, cartID
func (_m *Carts) ClearCart(ctx context.Context, cartID primitive.ObjectID) error {
	ret := _m.Called(ctx, cartID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, cartID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, cart
func (_m *Carts) Create(ctx context.Context, cart domain.Cart) (domain.Cart, error) {
	ret := _m.Called(ctx, cart)

	var r0 domain.Cart
	if rf, ok := ret.Get(0).(func(context.Context, domain.Cart) domain.Cart); ok {
		r0 = rf(ctx, cart)
	} else {
		r0 = ret.Get(0).(domain.Cart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Cart) error); ok {
		r1 = rf(ctx, cart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, cartID
func (_m *Carts) Delete(ctx context.Context, cartID primitive.ObjectID) error {
	ret := _m.Called(ctx, cartID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, cartID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCartItem provides a mock function with given fields: ctx, productID, cartID
func (_m *Carts) DeleteCartItem(ctx context.Context, productID primitive.ObjectID, cartID primitive.ObjectID) error {
	ret := _m.Called(ctx, productID, cartID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = rf(ctx, productID, cartID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx
func (_m *Carts) FindAll(ctx context.Context) ([]domain.Cart, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Cart
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Cart); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Cart)
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

// FindByID provides a mock function with given fields: ctx, cartID
func (_m *Carts) FindByID(ctx context.Context, cartID primitive.ObjectID) (domain.Cart, error) {
	ret := _m.Called(ctx, cartID)

	var r0 domain.Cart
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) domain.Cart); ok {
		r0 = rf(ctx, cartID)
	} else {
		r0 = ret.Get(0).(domain.Cart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCartItems provides a mock function with given fields: ctx, cartID
func (_m *Carts) FindCartItems(ctx context.Context, cartID primitive.ObjectID) ([]domain.CartItem, error) {
	ret := _m.Called(ctx, cartID)

	var r0 []domain.CartItem
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) []domain.CartItem); ok {
		r0 = rf(ctx, cartID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.CartItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, cartInput, cartID
func (_m *Carts) Update(ctx context.Context, cartInput dto.UpdateCartInput, cartID primitive.ObjectID) (domain.Cart, error) {
	ret := _m.Called(ctx, cartInput, cartID)

	var r0 domain.Cart
	if rf, ok := ret.Get(0).(func(context.Context, dto.UpdateCartInput, primitive.ObjectID) domain.Cart); ok {
		r0 = rf(ctx, cartInput, cartID)
	} else {
		r0 = ret.Get(0).(domain.Cart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.UpdateCartInput, primitive.ObjectID) error); ok {
		r1 = rf(ctx, cartInput, cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCartItem provides a mock function with given fields: ctx, cartItem, cartID
func (_m *Carts) UpdateCartItem(ctx context.Context, cartItem domain.CartItem, cartID primitive.ObjectID) (domain.CartItem, error) {
	ret := _m.Called(ctx, cartItem, cartID)

	var r0 domain.CartItem
	if rf, ok := ret.Get(0).(func(context.Context, domain.CartItem, primitive.ObjectID) domain.CartItem); ok {
		r0 = rf(ctx, cartItem, cartID)
	} else {
		r0 = ret.Get(0).(domain.CartItem)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.CartItem, primitive.ObjectID) error); ok {
		r1 = rf(ctx, cartItem, cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCarts creates a new instance of Carts. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCarts(t testing.TB) *Carts {
	mock := &Carts{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
