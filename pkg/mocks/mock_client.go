// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/cluster-api-provider-cloudstack/pkg/cloud (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	cloudstack "github.com/apache/cloudstack-go/v2/cloudstack"
	v1beta1 "github.com/aws/cluster-api-provider-cloudstack/api/v1beta1"
	cloud "github.com/aws/cluster-api-provider-cloudstack/pkg/cloud"
	gomock "github.com/golang/mock/gomock"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AssignVMToLoadBalancerRule mocks base method.
func (m *MockClient) AssignVMToLoadBalancerRule(arg0 *v1beta1.CloudStackCluster, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignVMToLoadBalancerRule", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignVMToLoadBalancerRule indicates an expected call of AssignVMToLoadBalancerRule.
func (mr *MockClientMockRecorder) AssignVMToLoadBalancerRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignVMToLoadBalancerRule", reflect.TypeOf((*MockClient)(nil).AssignVMToLoadBalancerRule), arg0, arg1)
}

// AssociateAffinityGroup mocks base method.
func (m *MockClient) AssociateAffinityGroup(arg0 *v1beta1.CloudStackMachine, arg1 cloud.AffinityGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssociateAffinityGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssociateAffinityGroup indicates an expected call of AssociateAffinityGroup.
func (mr *MockClientMockRecorder) AssociateAffinityGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateAffinityGroup", reflect.TypeOf((*MockClient)(nil).AssociateAffinityGroup), arg0, arg1)
}

// DeleteAffinityGroup mocks base method.
func (m *MockClient) DeleteAffinityGroup(arg0 *cloud.AffinityGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAffinityGroup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAffinityGroup indicates an expected call of DeleteAffinityGroup.
func (mr *MockClientMockRecorder) DeleteAffinityGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAffinityGroup", reflect.TypeOf((*MockClient)(nil).DeleteAffinityGroup), arg0)
}

// DestroyVMInstance mocks base method.
func (m *MockClient) DestroyVMInstance(arg0 *v1beta1.CloudStackMachine) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyVMInstance", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyVMInstance indicates an expected call of DestroyVMInstance.
func (mr *MockClientMockRecorder) DestroyVMInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyVMInstance", reflect.TypeOf((*MockClient)(nil).DestroyVMInstance), arg0)
}

// DissassociateAffinityGroup mocks base method.
func (m *MockClient) DissassociateAffinityGroup(arg0 *v1beta1.CloudStackMachine, arg1 cloud.AffinityGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DissassociateAffinityGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DissassociateAffinityGroup indicates an expected call of DissassociateAffinityGroup.
func (mr *MockClientMockRecorder) DissassociateAffinityGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DissassociateAffinityGroup", reflect.TypeOf((*MockClient)(nil).DissassociateAffinityGroup), arg0, arg1)
}

// FetchAffinityGroup mocks base method.
func (m *MockClient) FetchAffinityGroup(arg0 *cloud.AffinityGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAffinityGroup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchAffinityGroup indicates an expected call of FetchAffinityGroup.
func (mr *MockClientMockRecorder) FetchAffinityGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAffinityGroup", reflect.TypeOf((*MockClient)(nil).FetchAffinityGroup), arg0)
}

// GetOrCreateAffinityGroup mocks base method.
func (m *MockClient) GetOrCreateAffinityGroup(arg0 *v1beta1.CloudStackCluster, arg1 *cloud.AffinityGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateAffinityGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetOrCreateAffinityGroup indicates an expected call of GetOrCreateAffinityGroup.
func (mr *MockClientMockRecorder) GetOrCreateAffinityGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateAffinityGroup", reflect.TypeOf((*MockClient)(nil).GetOrCreateAffinityGroup), arg0, arg1)
}

// GetOrCreateCluster mocks base method.
func (m *MockClient) GetOrCreateCluster(arg0 *v1beta1.CloudStackCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateCluster", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetOrCreateCluster indicates an expected call of GetOrCreateCluster.
func (mr *MockClientMockRecorder) GetOrCreateCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateCluster", reflect.TypeOf((*MockClient)(nil).GetOrCreateCluster), arg0)
}

// GetOrCreateLoadBalancerRule mocks base method.
func (m *MockClient) GetOrCreateLoadBalancerRule(arg0 *v1beta1.CloudStackCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateLoadBalancerRule", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetOrCreateLoadBalancerRule indicates an expected call of GetOrCreateLoadBalancerRule.
func (mr *MockClientMockRecorder) GetOrCreateLoadBalancerRule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateLoadBalancerRule", reflect.TypeOf((*MockClient)(nil).GetOrCreateLoadBalancerRule), arg0)
}

// GetOrCreateNetwork mocks base method.
func (m *MockClient) GetOrCreateNetwork(arg0 *v1beta1.CloudStackCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateNetwork", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetOrCreateNetwork indicates an expected call of GetOrCreateNetwork.
func (mr *MockClientMockRecorder) GetOrCreateNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateNetwork", reflect.TypeOf((*MockClient)(nil).GetOrCreateNetwork), arg0)
}

// GetOrCreateVMInstance mocks base method.
func (m *MockClient) GetOrCreateVMInstance(arg0 *v1beta1.CloudStackMachine, arg1 *v1beta10.Machine, arg2 *v1beta1.CloudStackCluster, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateVMInstance", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetOrCreateVMInstance indicates an expected call of GetOrCreateVMInstance.
func (mr *MockClientMockRecorder) GetOrCreateVMInstance(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateVMInstance", reflect.TypeOf((*MockClient)(nil).GetOrCreateVMInstance), arg0, arg1, arg2, arg3)
}

// OpenFirewallRules mocks base method.
func (m *MockClient) OpenFirewallRules(arg0 *v1beta1.CloudStackCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFirewallRules", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenFirewallRules indicates an expected call of OpenFirewallRules.
func (mr *MockClientMockRecorder) OpenFirewallRules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFirewallRules", reflect.TypeOf((*MockClient)(nil).OpenFirewallRules), arg0)
}

// ResolveLoadBalancerRuleDetails mocks base method.
func (m *MockClient) ResolveLoadBalancerRuleDetails(arg0 *v1beta1.CloudStackCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveLoadBalancerRuleDetails", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResolveLoadBalancerRuleDetails indicates an expected call of ResolveLoadBalancerRuleDetails.
func (mr *MockClientMockRecorder) ResolveLoadBalancerRuleDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveLoadBalancerRuleDetails", reflect.TypeOf((*MockClient)(nil).ResolveLoadBalancerRuleDetails), arg0)
}

// ResolveNetwork mocks base method.
func (m *MockClient) ResolveNetwork(arg0 *v1beta1.CloudStackCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveNetwork", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResolveNetwork indicates an expected call of ResolveNetwork.
func (mr *MockClientMockRecorder) ResolveNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveNetwork", reflect.TypeOf((*MockClient)(nil).ResolveNetwork), arg0)
}

// ResolvePublicIPDetails mocks base method.
func (m *MockClient) ResolvePublicIPDetails(arg0 *v1beta1.CloudStackCluster) (*cloudstack.PublicIpAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolvePublicIPDetails", arg0)
	ret0, _ := ret[0].(*cloudstack.PublicIpAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolvePublicIPDetails indicates an expected call of ResolvePublicIPDetails.
func (mr *MockClientMockRecorder) ResolvePublicIPDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolvePublicIPDetails", reflect.TypeOf((*MockClient)(nil).ResolvePublicIPDetails), arg0)
}

// ResolveVMInstanceDetails mocks base method.
func (m *MockClient) ResolveVMInstanceDetails(arg0 *v1beta1.CloudStackMachine) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveVMInstanceDetails", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResolveVMInstanceDetails indicates an expected call of ResolveVMInstanceDetails.
func (mr *MockClientMockRecorder) ResolveVMInstanceDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveVMInstanceDetails", reflect.TypeOf((*MockClient)(nil).ResolveVMInstanceDetails), arg0)
}
