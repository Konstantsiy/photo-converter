// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	gomock0 "io"
	"reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// DownloadFile mocks base method.
func (m *MockStorage) DownloadFile(fileID string) (gomock0.ReadSeeker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", fileID)
	ret0, _ := ret[0].(gomock0.ReadSeeker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadFile indicates an expected call of DownloadFile.
func (mr *MockStorageMockRecorder) DownloadFile(fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockStorage)(nil).DownloadFile), fileID)
}

// GetDownloadURL mocks base method.
func (m *MockStorage) GetDownloadURL(fileID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDownloadURL", fileID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDownloadURL indicates an expected call of GetDownloadURL.
func (mr *MockStorageMockRecorder) GetDownloadURL(fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadURL", reflect.TypeOf((*MockStorage)(nil).GetDownloadURL), fileID)
}

// UploadFile mocks base method.
func (m *MockStorage) UploadFile(file gomock0.ReadSeeker, fileID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", file, fileID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockStorageMockRecorder) UploadFile(file, fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockStorage)(nil).UploadFile), file, fileID)
}
