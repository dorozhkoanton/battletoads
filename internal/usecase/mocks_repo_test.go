// Code generated by MockGen. DO NOT EDIT.
// Source: contracts.go
//
// Generated by this command:
//
//	mockgen -source=contracts.go -destination=../usecase/mocks_repo_test.go -package=usecase_test
//

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"

	entity "github.com/dorozhkoanton/battletoads/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockTranslationRepo is a mock of TranslationRepo interface.
type MockTranslationRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTranslationRepoMockRecorder
	isgomock struct{}
}

// MockTranslationRepoMockRecorder is the mock recorder for MockTranslationRepo.
type MockTranslationRepoMockRecorder struct {
	mock *MockTranslationRepo
}

// NewMockTranslationRepo creates a new mock instance.
func NewMockTranslationRepo(ctrl *gomock.Controller) *MockTranslationRepo {
	mock := &MockTranslationRepo{ctrl: ctrl}
	mock.recorder = &MockTranslationRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslationRepo) EXPECT() *MockTranslationRepoMockRecorder {
	return m.recorder
}

// GetHistory mocks base method.
func (m *MockTranslationRepo) GetHistory(arg0 context.Context) ([]entity.Translation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistory", arg0)
	ret0, _ := ret[0].([]entity.Translation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistory indicates an expected call of GetHistory.
func (mr *MockTranslationRepoMockRecorder) GetHistory(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistory", reflect.TypeOf((*MockTranslationRepo)(nil).GetHistory), arg0)
}

// Store mocks base method.
func (m *MockTranslationRepo) Store(arg0 context.Context, arg1 entity.Translation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockTranslationRepoMockRecorder) Store(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockTranslationRepo)(nil).Store), arg0, arg1)
}

// MockTranslationWebAPI is a mock of TranslationWebAPI interface.
type MockTranslationWebAPI struct {
	ctrl     *gomock.Controller
	recorder *MockTranslationWebAPIMockRecorder
	isgomock struct{}
}

// MockTranslationWebAPIMockRecorder is the mock recorder for MockTranslationWebAPI.
type MockTranslationWebAPIMockRecorder struct {
	mock *MockTranslationWebAPI
}

// NewMockTranslationWebAPI creates a new mock instance.
func NewMockTranslationWebAPI(ctrl *gomock.Controller) *MockTranslationWebAPI {
	mock := &MockTranslationWebAPI{ctrl: ctrl}
	mock.recorder = &MockTranslationWebAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslationWebAPI) EXPECT() *MockTranslationWebAPIMockRecorder {
	return m.recorder
}

// Translate mocks base method.
func (m *MockTranslationWebAPI) Translate(arg0 entity.Translation) (entity.Translation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", arg0)
	ret0, _ := ret[0].(entity.Translation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Translate indicates an expected call of Translate.
func (mr *MockTranslationWebAPIMockRecorder) Translate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockTranslationWebAPI)(nil).Translate), arg0)
}

// MockIntervalCommandRepo is a mock of IntervalCommandRepo interface.
type MockIntervalCommandRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIntervalCommandRepoMockRecorder
	isgomock struct{}
}

// MockIntervalCommandRepoMockRecorder is the mock recorder for MockIntervalCommandRepo.
type MockIntervalCommandRepoMockRecorder struct {
	mock *MockIntervalCommandRepo
}

// NewMockIntervalCommandRepo creates a new mock instance.
func NewMockIntervalCommandRepo(ctrl *gomock.Controller) *MockIntervalCommandRepo {
	mock := &MockIntervalCommandRepo{ctrl: ctrl}
	mock.recorder = &MockIntervalCommandRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntervalCommandRepo) EXPECT() *MockIntervalCommandRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIntervalCommandRepo) Create(arg0 context.Context, arg1 entity.IntervalCommand) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIntervalCommandRepoMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIntervalCommandRepo)(nil).Create), arg0, arg1)
}

// MockScheduledCommandRepo is a mock of ScheduledCommandRepo interface.
type MockScheduledCommandRepo struct {
	ctrl     *gomock.Controller
	recorder *MockScheduledCommandRepoMockRecorder
	isgomock struct{}
}

// MockScheduledCommandRepoMockRecorder is the mock recorder for MockScheduledCommandRepo.
type MockScheduledCommandRepoMockRecorder struct {
	mock *MockScheduledCommandRepo
}

// NewMockScheduledCommandRepo creates a new mock instance.
func NewMockScheduledCommandRepo(ctrl *gomock.Controller) *MockScheduledCommandRepo {
	mock := &MockScheduledCommandRepo{ctrl: ctrl}
	mock.recorder = &MockScheduledCommandRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduledCommandRepo) EXPECT() *MockScheduledCommandRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockScheduledCommandRepo) Create(arg0 context.Context, arg1 entity.ScheduledCommand) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockScheduledCommandRepoMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockScheduledCommandRepo)(nil).Create), arg0, arg1)
}
