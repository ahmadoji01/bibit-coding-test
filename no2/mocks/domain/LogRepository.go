// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "go_bibit_test/domain"

	mock "github.com/stretchr/testify/mock"
)

// LogRepository is an autogenerated mock type for the LogRepository type
type LogRepository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *LogRepository) Fetch(ctx context.Context, cursor string, num int64) ([]domain.Log, string, error) {
	ret := _m.Called(ctx, cursor, num)

	var r0 []domain.Log
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []domain.Log); ok {
		r0 = rf(ctx, cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Log)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) string); ok {
		r1 = rf(ctx, cursor, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int64) error); ok {
		r2 = rf(ctx, cursor, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *LogRepository) GetByID(ctx context.Context, id int64) (domain.Log, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Log
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Log); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Log)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, l
func (_m *LogRepository) Store(ctx context.Context, l *domain.Log) error {
	ret := _m.Called(ctx, l)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Log) error); ok {
		r0 = rf(ctx, l)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
