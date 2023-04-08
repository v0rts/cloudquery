// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset (interfaces: Interface)

// Package apiextensions is a generated GoMock package.
package apiextensions

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	discovery "k8s.io/client-go/discovery"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// ApiextensionsV1 mocks base method.
func (m *MockInterface) ApiextensionsV1() v1.ApiextensionsV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApiextensionsV1")
	ret0, _ := ret[0].(v1.ApiextensionsV1Interface)
	return ret0
}

// ApiextensionsV1 indicates an expected call of ApiextensionsV1.
func (mr *MockInterfaceMockRecorder) ApiextensionsV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApiextensionsV1", reflect.TypeOf((*MockInterface)(nil).ApiextensionsV1))
}

// ApiextensionsV1beta1 mocks base method.
func (m *MockInterface) ApiextensionsV1beta1() v1beta1.ApiextensionsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApiextensionsV1beta1")
	ret0, _ := ret[0].(v1beta1.ApiextensionsV1beta1Interface)
	return ret0
}

// ApiextensionsV1beta1 indicates an expected call of ApiextensionsV1beta1.
func (mr *MockInterfaceMockRecorder) ApiextensionsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApiextensionsV1beta1", reflect.TypeOf((*MockInterface)(nil).ApiextensionsV1beta1))
}

// Discovery mocks base method.
func (m *MockInterface) Discovery() discovery.DiscoveryInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discovery")
	ret0, _ := ret[0].(discovery.DiscoveryInterface)
	return ret0
}

// Discovery indicates an expected call of Discovery.
func (mr *MockInterfaceMockRecorder) Discovery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discovery", reflect.TypeOf((*MockInterface)(nil).Discovery))
}
