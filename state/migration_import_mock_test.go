// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: TransactionRunner,StateDocumentFactory,DocModelNamespace,RemoteEntitiesDescription,RelationNetworksDescription,RemoteApplicationsDescription)

// Package state is a generated GoMock package.
package state

import (
	gomock "github.com/golang/mock/gomock"
	description "github.com/juju/description"
	txn "gopkg.in/mgo.v2/txn"
	reflect "reflect"
)

// MockTransactionRunner is a mock of TransactionRunner interface
type MockTransactionRunner struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRunnerMockRecorder
}

// MockTransactionRunnerMockRecorder is the mock recorder for MockTransactionRunner
type MockTransactionRunnerMockRecorder struct {
	mock *MockTransactionRunner
}

// NewMockTransactionRunner creates a new mock instance
func NewMockTransactionRunner(ctrl *gomock.Controller) *MockTransactionRunner {
	mock := &MockTransactionRunner{ctrl: ctrl}
	mock.recorder = &MockTransactionRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionRunner) EXPECT() *MockTransactionRunnerMockRecorder {
	return m.recorder
}

// RunTransaction mocks base method
func (m *MockTransactionRunner) RunTransaction(arg0 []txn.Op) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTransaction indicates an expected call of RunTransaction
func (mr *MockTransactionRunnerMockRecorder) RunTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransaction", reflect.TypeOf((*MockTransactionRunner)(nil).RunTransaction), arg0)
}

// MockStateDocumentFactory is a mock of StateDocumentFactory interface
type MockStateDocumentFactory struct {
	ctrl     *gomock.Controller
	recorder *MockStateDocumentFactoryMockRecorder
}

// MockStateDocumentFactoryMockRecorder is the mock recorder for MockStateDocumentFactory
type MockStateDocumentFactoryMockRecorder struct {
	mock *MockStateDocumentFactory
}

// NewMockStateDocumentFactory creates a new mock instance
func NewMockStateDocumentFactory(ctrl *gomock.Controller) *MockStateDocumentFactory {
	mock := &MockStateDocumentFactory{ctrl: ctrl}
	mock.recorder = &MockStateDocumentFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateDocumentFactory) EXPECT() *MockStateDocumentFactoryMockRecorder {
	return m.recorder
}

// MakeFirewallRuleDoc mocks base method
func (m *MockStateDocumentFactory) MakeFirewallRuleDoc(arg0 description.FirewallRule) *firewallRulesDoc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeFirewallRuleDoc", arg0)
	ret0, _ := ret[0].(*firewallRulesDoc)
	return ret0
}

// MakeFirewallRuleDoc indicates an expected call of MakeFirewallRuleDoc
func (mr *MockStateDocumentFactoryMockRecorder) MakeFirewallRuleDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeFirewallRuleDoc", reflect.TypeOf((*MockStateDocumentFactory)(nil).MakeFirewallRuleDoc), arg0)
}

// MakeRemoteApplicationDoc mocks base method
func (m *MockStateDocumentFactory) MakeRemoteApplicationDoc(arg0 description.RemoteApplication) *remoteApplicationDoc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRemoteApplicationDoc", arg0)
	ret0, _ := ret[0].(*remoteApplicationDoc)
	return ret0
}

// MakeRemoteApplicationDoc indicates an expected call of MakeRemoteApplicationDoc
func (mr *MockStateDocumentFactoryMockRecorder) MakeRemoteApplicationDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRemoteApplicationDoc", reflect.TypeOf((*MockStateDocumentFactory)(nil).MakeRemoteApplicationDoc), arg0)
}

// MakeStatusDoc mocks base method
func (m *MockStateDocumentFactory) MakeStatusDoc(arg0 description.Status) statusDoc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeStatusDoc", arg0)
	ret0, _ := ret[0].(statusDoc)
	return ret0
}

// MakeStatusDoc indicates an expected call of MakeStatusDoc
func (mr *MockStateDocumentFactoryMockRecorder) MakeStatusDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeStatusDoc", reflect.TypeOf((*MockStateDocumentFactory)(nil).MakeStatusDoc), arg0)
}

// MakeStatusOp mocks base method
func (m *MockStateDocumentFactory) MakeStatusOp(arg0 string, arg1 statusDoc) txn.Op {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeStatusOp", arg0, arg1)
	ret0, _ := ret[0].(txn.Op)
	return ret0
}

// MakeStatusOp indicates an expected call of MakeStatusOp
func (mr *MockStateDocumentFactoryMockRecorder) MakeStatusOp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeStatusOp", reflect.TypeOf((*MockStateDocumentFactory)(nil).MakeStatusOp), arg0, arg1)
}

// NewRemoteApplication mocks base method
func (m *MockStateDocumentFactory) NewRemoteApplication(arg0 *remoteApplicationDoc) *RemoteApplication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRemoteApplication", arg0)
	ret0, _ := ret[0].(*RemoteApplication)
	return ret0
}

// NewRemoteApplication indicates an expected call of NewRemoteApplication
func (mr *MockStateDocumentFactoryMockRecorder) NewRemoteApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRemoteApplication", reflect.TypeOf((*MockStateDocumentFactory)(nil).NewRemoteApplication), arg0)
}

// MockDocModelNamespace is a mock of DocModelNamespace interface
type MockDocModelNamespace struct {
	ctrl     *gomock.Controller
	recorder *MockDocModelNamespaceMockRecorder
}

// MockDocModelNamespaceMockRecorder is the mock recorder for MockDocModelNamespace
type MockDocModelNamespaceMockRecorder struct {
	mock *MockDocModelNamespace
}

// NewMockDocModelNamespace creates a new mock instance
func NewMockDocModelNamespace(ctrl *gomock.Controller) *MockDocModelNamespace {
	mock := &MockDocModelNamespace{ctrl: ctrl}
	mock.recorder = &MockDocModelNamespaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDocModelNamespace) EXPECT() *MockDocModelNamespaceMockRecorder {
	return m.recorder
}

// DocID mocks base method
func (m *MockDocModelNamespace) DocID(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockDocModelNamespaceMockRecorder) DocID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockDocModelNamespace)(nil).DocID), arg0)
}

// MockRemoteEntitiesDescription is a mock of RemoteEntitiesDescription interface
type MockRemoteEntitiesDescription struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteEntitiesDescriptionMockRecorder
}

// MockRemoteEntitiesDescriptionMockRecorder is the mock recorder for MockRemoteEntitiesDescription
type MockRemoteEntitiesDescriptionMockRecorder struct {
	mock *MockRemoteEntitiesDescription
}

// NewMockRemoteEntitiesDescription creates a new mock instance
func NewMockRemoteEntitiesDescription(ctrl *gomock.Controller) *MockRemoteEntitiesDescription {
	mock := &MockRemoteEntitiesDescription{ctrl: ctrl}
	mock.recorder = &MockRemoteEntitiesDescriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemoteEntitiesDescription) EXPECT() *MockRemoteEntitiesDescriptionMockRecorder {
	return m.recorder
}

// DocID mocks base method
func (m *MockRemoteEntitiesDescription) DocID(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockRemoteEntitiesDescriptionMockRecorder) DocID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockRemoteEntitiesDescription)(nil).DocID), arg0)
}

// RemoteEntities mocks base method
func (m *MockRemoteEntitiesDescription) RemoteEntities() []description.RemoteEntity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteEntities")
	ret0, _ := ret[0].([]description.RemoteEntity)
	return ret0
}

// RemoteEntities indicates an expected call of RemoteEntities
func (mr *MockRemoteEntitiesDescriptionMockRecorder) RemoteEntities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteEntities", reflect.TypeOf((*MockRemoteEntitiesDescription)(nil).RemoteEntities))
}

// MockRelationNetworksDescription is a mock of RelationNetworksDescription interface
type MockRelationNetworksDescription struct {
	ctrl     *gomock.Controller
	recorder *MockRelationNetworksDescriptionMockRecorder
}

// MockRelationNetworksDescriptionMockRecorder is the mock recorder for MockRelationNetworksDescription
type MockRelationNetworksDescriptionMockRecorder struct {
	mock *MockRelationNetworksDescription
}

// NewMockRelationNetworksDescription creates a new mock instance
func NewMockRelationNetworksDescription(ctrl *gomock.Controller) *MockRelationNetworksDescription {
	mock := &MockRelationNetworksDescription{ctrl: ctrl}
	mock.recorder = &MockRelationNetworksDescriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRelationNetworksDescription) EXPECT() *MockRelationNetworksDescriptionMockRecorder {
	return m.recorder
}

// DocID mocks base method
func (m *MockRelationNetworksDescription) DocID(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockRelationNetworksDescriptionMockRecorder) DocID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockRelationNetworksDescription)(nil).DocID), arg0)
}

// RelationNetworks mocks base method
func (m *MockRelationNetworksDescription) RelationNetworks() []description.RelationNetwork {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationNetworks")
	ret0, _ := ret[0].([]description.RelationNetwork)
	return ret0
}

// RelationNetworks indicates an expected call of RelationNetworks
func (mr *MockRelationNetworksDescriptionMockRecorder) RelationNetworks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationNetworks", reflect.TypeOf((*MockRelationNetworksDescription)(nil).RelationNetworks))
}

// MockRemoteApplicationsDescription is a mock of RemoteApplicationsDescription interface
type MockRemoteApplicationsDescription struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteApplicationsDescriptionMockRecorder
}

// MockRemoteApplicationsDescriptionMockRecorder is the mock recorder for MockRemoteApplicationsDescription
type MockRemoteApplicationsDescriptionMockRecorder struct {
	mock *MockRemoteApplicationsDescription
}

// NewMockRemoteApplicationsDescription creates a new mock instance
func NewMockRemoteApplicationsDescription(ctrl *gomock.Controller) *MockRemoteApplicationsDescription {
	mock := &MockRemoteApplicationsDescription{ctrl: ctrl}
	mock.recorder = &MockRemoteApplicationsDescriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemoteApplicationsDescription) EXPECT() *MockRemoteApplicationsDescriptionMockRecorder {
	return m.recorder
}

// DocID mocks base method
func (m *MockRemoteApplicationsDescription) DocID(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockRemoteApplicationsDescriptionMockRecorder) DocID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockRemoteApplicationsDescription)(nil).DocID), arg0)
}

// MakeFirewallRuleDoc mocks base method
func (m *MockRemoteApplicationsDescription) MakeFirewallRuleDoc(arg0 description.FirewallRule) *firewallRulesDoc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeFirewallRuleDoc", arg0)
	ret0, _ := ret[0].(*firewallRulesDoc)
	return ret0
}

// MakeFirewallRuleDoc indicates an expected call of MakeFirewallRuleDoc
func (mr *MockRemoteApplicationsDescriptionMockRecorder) MakeFirewallRuleDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeFirewallRuleDoc", reflect.TypeOf((*MockRemoteApplicationsDescription)(nil).MakeFirewallRuleDoc), arg0)
}

// MakeRemoteApplicationDoc mocks base method
func (m *MockRemoteApplicationsDescription) MakeRemoteApplicationDoc(arg0 description.RemoteApplication) *remoteApplicationDoc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRemoteApplicationDoc", arg0)
	ret0, _ := ret[0].(*remoteApplicationDoc)
	return ret0
}

// MakeRemoteApplicationDoc indicates an expected call of MakeRemoteApplicationDoc
func (mr *MockRemoteApplicationsDescriptionMockRecorder) MakeRemoteApplicationDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRemoteApplicationDoc", reflect.TypeOf((*MockRemoteApplicationsDescription)(nil).MakeRemoteApplicationDoc), arg0)
}

// MakeStatusDoc mocks base method
func (m *MockRemoteApplicationsDescription) MakeStatusDoc(arg0 description.Status) statusDoc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeStatusDoc", arg0)
	ret0, _ := ret[0].(statusDoc)
	return ret0
}

// MakeStatusDoc indicates an expected call of MakeStatusDoc
func (mr *MockRemoteApplicationsDescriptionMockRecorder) MakeStatusDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeStatusDoc", reflect.TypeOf((*MockRemoteApplicationsDescription)(nil).MakeStatusDoc), arg0)
}

// MakeStatusOp mocks base method
func (m *MockRemoteApplicationsDescription) MakeStatusOp(arg0 string, arg1 statusDoc) txn.Op {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeStatusOp", arg0, arg1)
	ret0, _ := ret[0].(txn.Op)
	return ret0
}

// MakeStatusOp indicates an expected call of MakeStatusOp
func (mr *MockRemoteApplicationsDescriptionMockRecorder) MakeStatusOp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeStatusOp", reflect.TypeOf((*MockRemoteApplicationsDescription)(nil).MakeStatusOp), arg0, arg1)
}

// NewRemoteApplication mocks base method
func (m *MockRemoteApplicationsDescription) NewRemoteApplication(arg0 *remoteApplicationDoc) *RemoteApplication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRemoteApplication", arg0)
	ret0, _ := ret[0].(*RemoteApplication)
	return ret0
}

// NewRemoteApplication indicates an expected call of NewRemoteApplication
func (mr *MockRemoteApplicationsDescriptionMockRecorder) NewRemoteApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRemoteApplication", reflect.TypeOf((*MockRemoteApplicationsDescription)(nil).NewRemoteApplication), arg0)
}

// RemoteApplications mocks base method
func (m *MockRemoteApplicationsDescription) RemoteApplications() []description.RemoteApplication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteApplications")
	ret0, _ := ret[0].([]description.RemoteApplication)
	return ret0
}

// RemoteApplications indicates an expected call of RemoteApplications
func (mr *MockRemoteApplicationsDescriptionMockRecorder) RemoteApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteApplications", reflect.TypeOf((*MockRemoteApplicationsDescription)(nil).RemoteApplications))
}
