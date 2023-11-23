// Code generated by MockGen. DO NOT EDIT.
// Source: types/handler.go
//
// Generated by this command:
//
//	mockgen -source=types/handler.go -package=mocks -destination=types/mocks/interfaces.go PostDecorator
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "go.uber.org/mock/gomock"
)

// MockAnteDecorator is a mock of AnteDecorator interface.
type MockAnteDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockAnteDecoratorMockRecorder
}

// MockAnteDecoratorMockRecorder is the mock recorder for MockAnteDecorator.
type MockAnteDecoratorMockRecorder struct {
	mock *MockAnteDecorator
}

// NewMockAnteDecorator creates a new mock instance.
func NewMockAnteDecorator(ctrl *gomock.Controller) *MockAnteDecorator {
	mock := &MockAnteDecorator{ctrl: ctrl}
	mock.recorder = &MockAnteDecoratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnteDecorator) EXPECT() *MockAnteDecoratorMockRecorder {
	return m.recorder
}

// AnteHandle mocks base method.
func (m *MockAnteDecorator) AnteHandle(ctx types.Context, tx types.Tx, simulate bool, next types.AnteHandler) (types.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnteHandle", ctx, tx, simulate, next)
	ret0, _ := ret[0].(types.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnteHandle indicates an expected call of AnteHandle.
func (mr *MockAnteDecoratorMockRecorder) AnteHandle(ctx, tx, simulate, next any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnteHandle", reflect.TypeOf((*MockAnteDecorator)(nil).AnteHandle), ctx, tx, simulate, next)
}

// MockPostDecorator is a mock of PostDecorator interface.
type MockPostDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockPostDecoratorMockRecorder
}

// MockPostDecoratorMockRecorder is the mock recorder for MockPostDecorator.
type MockPostDecoratorMockRecorder struct {
	mock *MockPostDecorator
}

// NewMockPostDecorator creates a new mock instance.
func NewMockPostDecorator(ctrl *gomock.Controller) *MockPostDecorator {
	mock := &MockPostDecorator{ctrl: ctrl}
	mock.recorder = &MockPostDecoratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostDecorator) EXPECT() *MockPostDecoratorMockRecorder {
	return m.recorder
}

// PostHandle mocks base method.
func (m *MockPostDecorator) PostHandle(ctx types.Context, tx types.Tx, simulate, success bool, next types.PostHandler) (types.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostHandle", ctx, tx, simulate, success, next)
	ret0, _ := ret[0].(types.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostHandle indicates an expected call of PostHandle.
func (mr *MockPostDecoratorMockRecorder) PostHandle(ctx, tx, simulate, success, next any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostHandle", reflect.TypeOf((*MockPostDecorator)(nil).PostHandle), ctx, tx, simulate, success, next)
}