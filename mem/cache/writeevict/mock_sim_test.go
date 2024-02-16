// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sarchlab/akita/v4/sim (interfaces: Port,Buffer)

package writeevict

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sim "github.com/sarchlab/akita/v4/sim"
)

// MockPort is a mock of Port interface.
type MockPort struct {
	ctrl     *gomock.Controller
	recorder *MockPortMockRecorder
}

// MockPortMockRecorder is the mock recorder for MockPort.
type MockPortMockRecorder struct {
	mock *MockPort
}

// NewMockPort creates a new mock instance.
func NewMockPort(ctrl *gomock.Controller) *MockPort {
	mock := &MockPort{ctrl: ctrl}
	mock.recorder = &MockPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPort) EXPECT() *MockPortMockRecorder {
	return m.recorder
}

// AcceptHook mocks base method.
func (m *MockPort) AcceptHook(arg0 sim.Hook) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AcceptHook", arg0)
}

// AcceptHook indicates an expected call of AcceptHook.
func (mr *MockPortMockRecorder) AcceptHook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptHook", reflect.TypeOf((*MockPort)(nil).AcceptHook), arg0)
}

// CanSend mocks base method.
func (m *MockPort) CanSend() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSend")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanSend indicates an expected call of CanSend.
func (mr *MockPortMockRecorder) CanSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSend", reflect.TypeOf((*MockPort)(nil).CanSend))
}

// Component mocks base method.
func (m *MockPort) Component() sim.Component {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Component")
	ret0, _ := ret[0].(sim.Component)
	return ret0
}

// Component indicates an expected call of Component.
func (mr *MockPortMockRecorder) Component() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Component", reflect.TypeOf((*MockPort)(nil).Component))
}

// Deliver mocks base method.
func (m *MockPort) Deliver(arg0 sim.Msg) *sim.SendError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deliver", arg0)
	ret0, _ := ret[0].(*sim.SendError)
	return ret0
}

// Deliver indicates an expected call of Deliver.
func (mr *MockPortMockRecorder) Deliver(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deliver", reflect.TypeOf((*MockPort)(nil).Deliver), arg0)
}

// Hooks mocks base method.
func (m *MockPort) Hooks() []sim.Hook {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hooks")
	ret0, _ := ret[0].([]sim.Hook)
	return ret0
}

// Hooks indicates an expected call of Hooks.
func (mr *MockPortMockRecorder) Hooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hooks", reflect.TypeOf((*MockPort)(nil).Hooks))
}

// Name mocks base method.
func (m *MockPort) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockPortMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPort)(nil).Name))
}

// NotifyAvailable mocks base method.
func (m *MockPort) NotifyAvailable(arg0 sim.VTimeInSec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyAvailable", arg0)
}

// NotifyAvailable indicates an expected call of NotifyAvailable.
func (mr *MockPortMockRecorder) NotifyAvailable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyAvailable", reflect.TypeOf((*MockPort)(nil).NotifyAvailable), arg0)
}

// NumHooks mocks base method.
func (m *MockPort) NumHooks() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumHooks")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumHooks indicates an expected call of NumHooks.
func (mr *MockPortMockRecorder) NumHooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumHooks", reflect.TypeOf((*MockPort)(nil).NumHooks))
}

// PeekIncoming mocks base method.
func (m *MockPort) PeekIncoming() sim.Msg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeekIncoming")
	ret0, _ := ret[0].(sim.Msg)
	return ret0
}

// PeekIncoming indicates an expected call of PeekIncoming.
func (mr *MockPortMockRecorder) PeekIncoming() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeekIncoming", reflect.TypeOf((*MockPort)(nil).PeekIncoming))
}

// PeekOutgoing mocks base method.
func (m *MockPort) PeekOutgoing() sim.Msg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeekOutgoing")
	ret0, _ := ret[0].(sim.Msg)
	return ret0
}

// PeekOutgoing indicates an expected call of PeekOutgoing.
func (mr *MockPortMockRecorder) PeekOutgoing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeekOutgoing", reflect.TypeOf((*MockPort)(nil).PeekOutgoing))
}

// RetrieveIncoming mocks base method.
func (m *MockPort) RetrieveIncoming(arg0 sim.VTimeInSec) sim.Msg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveIncoming", arg0)
	ret0, _ := ret[0].(sim.Msg)
	return ret0
}

// RetrieveIncoming indicates an expected call of RetrieveIncoming.
func (mr *MockPortMockRecorder) RetrieveIncoming(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveIncoming", reflect.TypeOf((*MockPort)(nil).RetrieveIncoming), arg0)
}

// RetrieveOutgoing mocks base method.
func (m *MockPort) RetrieveOutgoing() sim.Msg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveOutgoing")
	ret0, _ := ret[0].(sim.Msg)
	return ret0
}

// RetrieveOutgoing indicates an expected call of RetrieveOutgoing.
func (mr *MockPortMockRecorder) RetrieveOutgoing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveOutgoing", reflect.TypeOf((*MockPort)(nil).RetrieveOutgoing))
}

// Send mocks base method.
func (m *MockPort) Send(arg0 sim.Msg) *sim.SendError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(*sim.SendError)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockPortMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockPort)(nil).Send), arg0)
}

// SetConnection mocks base method.
func (m *MockPort) SetConnection(arg0 sim.Connection) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConnection", arg0)
}

// SetConnection indicates an expected call of SetConnection.
func (mr *MockPortMockRecorder) SetConnection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConnection", reflect.TypeOf((*MockPort)(nil).SetConnection), arg0)
}

// MockBuffer is a mock of Buffer interface.
type MockBuffer struct {
	ctrl     *gomock.Controller
	recorder *MockBufferMockRecorder
}

// MockBufferMockRecorder is the mock recorder for MockBuffer.
type MockBufferMockRecorder struct {
	mock *MockBuffer
}

// NewMockBuffer creates a new mock instance.
func NewMockBuffer(ctrl *gomock.Controller) *MockBuffer {
	mock := &MockBuffer{ctrl: ctrl}
	mock.recorder = &MockBufferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuffer) EXPECT() *MockBufferMockRecorder {
	return m.recorder
}

// AcceptHook mocks base method.
func (m *MockBuffer) AcceptHook(arg0 sim.Hook) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AcceptHook", arg0)
}

// AcceptHook indicates an expected call of AcceptHook.
func (mr *MockBufferMockRecorder) AcceptHook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptHook", reflect.TypeOf((*MockBuffer)(nil).AcceptHook), arg0)
}

// CanPush mocks base method.
func (m *MockBuffer) CanPush() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanPush")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanPush indicates an expected call of CanPush.
func (mr *MockBufferMockRecorder) CanPush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanPush", reflect.TypeOf((*MockBuffer)(nil).CanPush))
}

// Capacity mocks base method.
func (m *MockBuffer) Capacity() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capacity")
	ret0, _ := ret[0].(int)
	return ret0
}

// Capacity indicates an expected call of Capacity.
func (mr *MockBufferMockRecorder) Capacity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capacity", reflect.TypeOf((*MockBuffer)(nil).Capacity))
}

// Clear mocks base method.
func (m *MockBuffer) Clear() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear.
func (mr *MockBufferMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockBuffer)(nil).Clear))
}

// Hooks mocks base method.
func (m *MockBuffer) Hooks() []sim.Hook {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hooks")
	ret0, _ := ret[0].([]sim.Hook)
	return ret0
}

// Hooks indicates an expected call of Hooks.
func (mr *MockBufferMockRecorder) Hooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hooks", reflect.TypeOf((*MockBuffer)(nil).Hooks))
}

// Name mocks base method.
func (m *MockBuffer) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockBufferMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockBuffer)(nil).Name))
}

// NumHooks mocks base method.
func (m *MockBuffer) NumHooks() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumHooks")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumHooks indicates an expected call of NumHooks.
func (mr *MockBufferMockRecorder) NumHooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumHooks", reflect.TypeOf((*MockBuffer)(nil).NumHooks))
}

// Peek mocks base method.
func (m *MockBuffer) Peek() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peek")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Peek indicates an expected call of Peek.
func (mr *MockBufferMockRecorder) Peek() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peek", reflect.TypeOf((*MockBuffer)(nil).Peek))
}

// Pop mocks base method.
func (m *MockBuffer) Pop() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pop")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Pop indicates an expected call of Pop.
func (mr *MockBufferMockRecorder) Pop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pop", reflect.TypeOf((*MockBuffer)(nil).Pop))
}

// Push mocks base method.
func (m *MockBuffer) Push(arg0 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Push", arg0)
}

// Push indicates an expected call of Push.
func (mr *MockBufferMockRecorder) Push(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockBuffer)(nil).Push), arg0)
}

// Size mocks base method.
func (m *MockBuffer) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockBufferMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockBuffer)(nil).Size))
}
