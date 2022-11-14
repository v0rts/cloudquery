// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/batch/v1 (interfaces: BatchV1Interface)

// Package v1 is a generated GoMock package.
package v1

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	rest "k8s.io/client-go/rest"
)

// MockBatchV1Interface is a mock of BatchV1Interface interface.
type MockBatchV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockBatchV1InterfaceMockRecorder
}

// MockBatchV1InterfaceMockRecorder is the mock recorder for MockBatchV1Interface.
type MockBatchV1InterfaceMockRecorder struct {
	mock *MockBatchV1Interface
}

// NewMockBatchV1Interface creates a new mock instance.
func NewMockBatchV1Interface(ctrl *gomock.Controller) *MockBatchV1Interface {
	mock := &MockBatchV1Interface{ctrl: ctrl}
	mock.recorder = &MockBatchV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBatchV1Interface) EXPECT() *MockBatchV1InterfaceMockRecorder {
	return m.recorder
}

// CronJobs mocks base method.
func (m *MockBatchV1Interface) CronJobs(arg0 string) v1.CronJobInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CronJobs", arg0)
	ret0, _ := ret[0].(v1.CronJobInterface)
	return ret0
}

// CronJobs indicates an expected call of CronJobs.
func (mr *MockBatchV1InterfaceMockRecorder) CronJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CronJobs", reflect.TypeOf((*MockBatchV1Interface)(nil).CronJobs), arg0)
}

// Jobs mocks base method.
func (m *MockBatchV1Interface) Jobs(arg0 string) v1.JobInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Jobs", arg0)
	ret0, _ := ret[0].(v1.JobInterface)
	return ret0
}

// Jobs indicates an expected call of Jobs.
func (mr *MockBatchV1InterfaceMockRecorder) Jobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jobs", reflect.TypeOf((*MockBatchV1Interface)(nil).Jobs), arg0)
}

// RESTClient mocks base method.
func (m *MockBatchV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockBatchV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockBatchV1Interface)(nil).RESTClient))
}
