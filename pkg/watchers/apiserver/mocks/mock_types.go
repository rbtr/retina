// Code generated by MockGen. DO NOT EDIT.
// Source: types.go
//
// Generated by this command:
//
//	mockgen -source=types.go -destination=mocks/mock_types.go -package=mocks .
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIHostResolver is a mock of IHostResolver interface.
type MockIHostResolver struct {
	ctrl     *gomock.Controller
	recorder *MockIHostResolverMockRecorder
}

// MockIHostResolverMockRecorder is the mock recorder for MockIHostResolver.
type MockIHostResolverMockRecorder struct {
	mock *MockIHostResolver
}

// NewMockIHostResolver creates a new mock instance.
func NewMockIHostResolver(ctrl *gomock.Controller) *MockIHostResolver {
	mock := &MockIHostResolver{ctrl: ctrl}
	mock.recorder = &MockIHostResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHostResolver) EXPECT() *MockIHostResolverMockRecorder {
	return m.recorder
}

// LookupHost mocks base method.
func (m *MockIHostResolver) LookupHost(context context.Context, host string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupHost", context, host)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupHost indicates an expected call of LookupHost.
func (mr *MockIHostResolverMockRecorder) LookupHost(context, host any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupHost", reflect.TypeOf((*MockIHostResolver)(nil).LookupHost), context, host)
}
