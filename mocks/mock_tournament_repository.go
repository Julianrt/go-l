// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jairogloz/go-l/pkg/ports (interfaces: TournamentRepository)
//
// Generated by this command:
//
//	mockgen -destination=mocks/mock_tournament_repository.go -package=mocks github.com/jairogloz/go-l/pkg/ports TournamentRepository
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/jairogloz/go-l/pkg/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockTournamentRepository is a mock of TournamentRepository interface.
type MockTournamentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTournamentRepositoryMockRecorder
}

// MockTournamentRepositoryMockRecorder is the mock recorder for MockTournamentRepository.
type MockTournamentRepositoryMockRecorder struct {
	mock *MockTournamentRepository
}

// NewMockTournamentRepository creates a new mock instance.
func NewMockTournamentRepository(ctrl *gomock.Controller) *MockTournamentRepository {
	mock := &MockTournamentRepository{ctrl: ctrl}
	mock.recorder = &MockTournamentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTournamentRepository) EXPECT() *MockTournamentRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockTournamentRepository) Insert(arg0 context.Context, arg1 *domain.Tournament) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockTournamentRepositoryMockRecorder) Insert(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTournamentRepository)(nil).Insert), arg0, arg1)
}
