// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sarchlab/akita/v3/mem/dram/internal/addressmapping (interfaces: Mapper)

package dram

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	addressmapping "github.com/sarchlab/akita/v3/mem/dram/internal/addressmapping"
)

// MockMapper is a mock of Mapper interface.
type MockMapper struct {
	ctrl     *gomock.Controller
	recorder *MockMapperMockRecorder
}

// MockMapperMockRecorder is the mock recorder for MockMapper.
type MockMapperMockRecorder struct {
	mock *MockMapper
}

// NewMockMapper creates a new mock instance.
func NewMockMapper(ctrl *gomock.Controller) *MockMapper {
	mock := &MockMapper{ctrl: ctrl}
	mock.recorder = &MockMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMapper) EXPECT() *MockMapperMockRecorder {
	return m.recorder
}

// Map mocks base method.
func (m *MockMapper) Map(arg0 uint64) addressmapping.Location {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map", arg0)
	ret0, _ := ret[0].(addressmapping.Location)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockMapperMockRecorder) Map(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockMapper)(nil).Map), arg0)
}
