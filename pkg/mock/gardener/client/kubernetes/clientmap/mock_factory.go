// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/client/kubernetes/clientmap (interfaces: ClientSetFactory)

// Package clientmap is a generated GoMock package.
package clientmap

import (
	context "context"
	kubernetes "github.com/gardener/gardener/pkg/client/kubernetes"
	clientmap "github.com/gardener/gardener/pkg/client/kubernetes/clientmap"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClientSetFactory is a mock of ClientSetFactory interface
type MockClientSetFactory struct {
	ctrl     *gomock.Controller
	recorder *MockClientSetFactoryMockRecorder
}

// MockClientSetFactoryMockRecorder is the mock recorder for MockClientSetFactory
type MockClientSetFactoryMockRecorder struct {
	mock *MockClientSetFactory
}

// NewMockClientSetFactory creates a new mock instance
func NewMockClientSetFactory(ctrl *gomock.Controller) *MockClientSetFactory {
	mock := &MockClientSetFactory{ctrl: ctrl}
	mock.recorder = &MockClientSetFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientSetFactory) EXPECT() *MockClientSetFactoryMockRecorder {
	return m.recorder
}

// CalculateClientSetHash mocks base method
func (m *MockClientSetFactory) CalculateClientSetHash(arg0 context.Context, arg1 clientmap.ClientSetKey) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateClientSetHash", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateClientSetHash indicates an expected call of CalculateClientSetHash
func (mr *MockClientSetFactoryMockRecorder) CalculateClientSetHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateClientSetHash", reflect.TypeOf((*MockClientSetFactory)(nil).CalculateClientSetHash), arg0, arg1)
}

// NewClientSet mocks base method
func (m *MockClientSetFactory) NewClientSet(arg0 context.Context, arg1 clientmap.ClientSetKey) (kubernetes.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewClientSet", arg0, arg1)
	ret0, _ := ret[0].(kubernetes.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewClientSet indicates an expected call of NewClientSet
func (mr *MockClientSetFactoryMockRecorder) NewClientSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClientSet", reflect.TypeOf((*MockClientSetFactory)(nil).NewClientSet), arg0, arg1)
}