// Code generated by MockGen. DO NOT EDIT.
// Source: repository/repo.go

// Package mock_repository is a generated GoMock package.
package repository

import (
	gomock "github.com/golang/mock/gomock"
	models "lockdown/models"
	reflect "reflect"
)

// MockRepo is a mock of Repo interface
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddTraderRegistrationDetails mocks base method
func (m *MockRepo) AddTraderRegistrationDetails(traderDetails models.TraderDetailsDb, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTraderRegistrationDetails", traderDetails, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTraderRegistrationDetails indicates an expected call of AddTraderRegistrationDetails
func (mr *MockRepoMockRecorder) AddTraderRegistrationDetails(traderDetails, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTraderRegistrationDetails", reflect.TypeOf((*MockRepo)(nil).AddTraderRegistrationDetails), traderDetails, id)
}

// GenerateUniqueId mocks base method
func (m *MockRepo) GenerateUniqueId() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUniqueId")
	ret0, _ := ret[0].(int)
	return ret0
}

// GenerateUniqueId indicates an expected call of GenerateUniqueId
func (mr *MockRepoMockRecorder) GenerateUniqueId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUniqueId", reflect.TypeOf((*MockRepo)(nil).GenerateUniqueId))
}

// GetAllTraderRegistrationDetails mocks base method
func (m *MockRepo) GetAllTraderRegistrationDetails() ([]models.TraderDetailsDb, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTraderRegistrationDetails")
	ret0, _ := ret[0].([]models.TraderDetailsDb)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTraderRegistrationDetails indicates an expected call of GetAllTraderRegistrationDetails
func (mr *MockRepoMockRecorder) GetAllTraderRegistrationDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTraderRegistrationDetails", reflect.TypeOf((*MockRepo)(nil).GetAllTraderRegistrationDetails))
}
