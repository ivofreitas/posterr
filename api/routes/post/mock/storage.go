// Code generated by MockGen. DO NOT EDIT.
// Source: api/routes/post/storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mysql "strider-backend-test.com/adapter/mysql"
	model "strider-backend-test.com/api/routes/post/model"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// BeginTransaction mocks base method.
func (m *MockRepository) BeginTransaction(ctx context.Context) (mysql.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTransaction", ctx)
	ret0, _ := ret[0].(mysql.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTransaction indicates an expected call of BeginTransaction.
func (mr *MockRepositoryMockRecorder) BeginTransaction(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockRepository)(nil).BeginTransaction), ctx)
}

// Commit mocks base method.
func (m *MockRepository) Commit(ctx context.Context, tx mysql.Transaction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Commit", ctx, tx)
}

// Commit indicates an expected call of Commit.
func (mr *MockRepositoryMockRecorder) Commit(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockRepository)(nil).Commit), ctx, tx)
}

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, tx mysql.Transaction, post *model.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, tx, post)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, tx, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, tx, post)
}

// GetPostCountToday mocks base method.
func (m *MockRepository) GetPostCountToday(ctx context.Context, userID string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostCountToday", ctx, userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostCountToday indicates an expected call of GetPostCountToday.
func (mr *MockRepositoryMockRecorder) GetPostCountToday(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostCountToday", reflect.TypeOf((*MockRepository)(nil).GetPostCountToday), ctx, userID)
}

// List mocks base method.
func (m *MockRepository) List(ctx context.Context, limit, offset int) ([]*model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset)
	ret0, _ := ret[0].([]*model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepositoryMockRecorder) List(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List), ctx, limit, offset)
}

// ListByFollower mocks base method.
func (m *MockRepository) ListByFollower(ctx context.Context, follower string, limit, offset int) ([]*model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByFollower", ctx, follower, limit, offset)
	ret0, _ := ret[0].([]*model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByFollower indicates an expected call of ListByFollower.
func (mr *MockRepositoryMockRecorder) ListByFollower(ctx, follower, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByFollower", reflect.TypeOf((*MockRepository)(nil).ListByFollower), ctx, follower, limit, offset)
}

// ListByUser mocks base method.
func (m *MockRepository) ListByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUser", ctx, userID, limit, offset)
	ret0, _ := ret[0].([]*model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUser indicates an expected call of ListByUser.
func (mr *MockRepositoryMockRecorder) ListByUser(ctx, userID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUser", reflect.TypeOf((*MockRepository)(nil).ListByUser), ctx, userID, limit, offset)
}

// QueryContext mocks base method.
func (m *MockRepository) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContext", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContext indicates an expected call of QueryContext.
func (mr *MockRepositoryMockRecorder) QueryContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockRepository)(nil).QueryContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockRepository) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockRepositoryMockRecorder) QueryRowContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockRepository)(nil).QueryRowContext), varargs...)
}

// Rollback mocks base method.
func (m *MockRepository) Rollback(ctx context.Context, tx mysql.Transaction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Rollback", ctx, tx)
}

// Rollback indicates an expected call of Rollback.
func (mr *MockRepositoryMockRecorder) Rollback(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockRepository)(nil).Rollback), ctx, tx)
}
