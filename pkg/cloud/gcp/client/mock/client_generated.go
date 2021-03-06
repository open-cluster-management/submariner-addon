// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	compute "google.golang.org/api/compute/v1"
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

// DeleteFirewallRule mocks base method.
func (m *MockInterface) DeleteFirewallRule(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFirewallRule", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFirewallRule indicates an expected call of DeleteFirewallRule.
func (mr *MockInterfaceMockRecorder) DeleteFirewallRule(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFirewallRule", reflect.TypeOf((*MockInterface)(nil).DeleteFirewallRule), name)
}

// DisablePublicIP mocks base method.
func (m *MockInterface) DisablePublicIP(instance *compute.Instance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisablePublicIP", instance)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisablePublicIP indicates an expected call of DisablePublicIP.
func (mr *MockInterfaceMockRecorder) DisablePublicIP(instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisablePublicIP", reflect.TypeOf((*MockInterface)(nil).DisablePublicIP), instance)
}

// EnablePublicIP mocks base method.
func (m *MockInterface) EnablePublicIP(instance *compute.Instance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePublicIP", instance)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnablePublicIP indicates an expected call of EnablePublicIP.
func (mr *MockInterfaceMockRecorder) EnablePublicIP(instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePublicIP", reflect.TypeOf((*MockInterface)(nil).EnablePublicIP), instance)
}

// GetFirewallRule mocks base method.
func (m *MockInterface) GetFirewallRule(name string) (*compute.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirewallRule", name)
	ret0, _ := ret[0].(*compute.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirewallRule indicates an expected call of GetFirewallRule.
func (mr *MockInterfaceMockRecorder) GetFirewallRule(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirewallRule", reflect.TypeOf((*MockInterface)(nil).GetFirewallRule), name)
}

// GetInstance mocks base method.
func (m *MockInterface) GetInstance(zone, instance string) (*compute.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstance", zone, instance)
	ret0, _ := ret[0].(*compute.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstance indicates an expected call of GetInstance.
func (mr *MockInterfaceMockRecorder) GetInstance(zone, instance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*MockInterface)(nil).GetInstance), zone, instance)
}

// GetProjectID mocks base method.
func (m *MockInterface) GetProjectID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetProjectID indicates an expected call of GetProjectID.
func (mr *MockInterfaceMockRecorder) GetProjectID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectID", reflect.TypeOf((*MockInterface)(nil).GetProjectID))
}

// InsertFirewallRule mocks base method.
func (m *MockInterface) InsertFirewallRule(rule *compute.Firewall) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertFirewallRule", rule)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertFirewallRule indicates an expected call of InsertFirewallRule.
func (mr *MockInterfaceMockRecorder) InsertFirewallRule(rule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertFirewallRule", reflect.TypeOf((*MockInterface)(nil).InsertFirewallRule), rule)
}

// UpdateFirewallRule mocks base method.
func (m *MockInterface) UpdateFirewallRule(name string, rule *compute.Firewall) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFirewallRule", name, rule)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFirewallRule indicates an expected call of UpdateFirewallRule.
func (mr *MockInterfaceMockRecorder) UpdateFirewallRule(name, rule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFirewallRule", reflect.TypeOf((*MockInterface)(nil).UpdateFirewallRule), name, rule)
}
