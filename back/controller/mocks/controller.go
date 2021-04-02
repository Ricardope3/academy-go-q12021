// Code generated by MockGen. DO NOT EDIT.
// Source: controller/controller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ricardope3/academy-go-q12021/back/models"
)

// MockEntity is a mock of Entity interface.
type MockEntity struct {
	ctrl     *gomock.Controller
	recorder *MockEntityMockRecorder
}

// MockEntityMockRecorder is the mock recorder for MockEntity.
type MockEntityMockRecorder struct {
	mock *MockEntity
}

// NewMockEntity creates a new mock instance.
func NewMockEntity(ctrl *gomock.Controller) *MockEntity {
	mock := &MockEntity{ctrl: ctrl}
	mock.recorder = &MockEntityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntity) EXPECT() *MockEntityMockRecorder {
	return m.recorder
}

// GetPokemonFromCSV mocks base method.
func (m *MockEntity) GetPokemonFromCSV(requestedId int) ([]models.Pokemon, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPokemonFromCSV", requestedId)
	ret0, _ := ret[0].([]models.Pokemon)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// GetPokemonFromCSV indicates an expected call of GetPokemonFromCSV.
func (mr *MockEntityMockRecorder) GetPokemonFromCSV(requestedId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPokemonFromCSV", reflect.TypeOf((*MockEntity)(nil).GetPokemonFromCSV), requestedId)
}

// SaveCSV mocks base method.
func (m *MockEntity) SaveCSV(todoArray []models.Todo) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCSV", todoArray)
	ret0, _ := ret[0].(int)
	return ret0
}

// SaveCSV indicates an expected call of SaveCSV.
func (mr *MockEntityMockRecorder) SaveCSV(todoArray interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCSV", reflect.TypeOf((*MockEntity)(nil).SaveCSV), todoArray)
}
