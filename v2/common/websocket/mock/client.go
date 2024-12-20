// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	websocket "github.com/adshao/go-binance/v2/common/websocket"
	gomock "github.com/golang/mock/gomock"
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

// GetReadChannel mocks base method.
func (m *MockClient) GetReadChannel() <-chan []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadChannel")
	ret0, _ := ret[0].(<-chan []byte)
	return ret0
}

// GetReadChannel indicates an expected call of GetReadChannel.
func (mr *MockClientMockRecorder) GetReadChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadChannel", reflect.TypeOf((*MockClient)(nil).GetReadChannel))
}

// GetReadErrorChannel mocks base method.
func (m *MockClient) GetReadErrorChannel() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadErrorChannel")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// GetReadErrorChannel indicates an expected call of GetReadErrorChannel.
func (mr *MockClientMockRecorder) GetReadErrorChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadErrorChannel", reflect.TypeOf((*MockClient)(nil).GetReadErrorChannel))
}

// GetReconnectCount mocks base method.
func (m *MockClient) GetReconnectCount() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReconnectCount")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetReconnectCount indicates an expected call of GetReconnectCount.
func (mr *MockClientMockRecorder) GetReconnectCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReconnectCount", reflect.TypeOf((*MockClient)(nil).GetReconnectCount))
}

// Wait mocks base method.
func (m *MockClient) Wait(timeout time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Wait", timeout)
}

// Wait indicates an expected call of Wait.
func (mr *MockClientMockRecorder) Wait(timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockClient)(nil).Wait), timeout)
}

// Write mocks base method.
func (m *MockClient) Write(id string, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", id, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockClientMockRecorder) Write(id, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockClient)(nil).Write), id, data)
}

// WriteSync mocks base method.
func (m *MockClient) WriteSync(id string, data []byte, timeout time.Duration) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteSync", id, data, timeout)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteSync indicates an expected call of WriteSync.
func (mr *MockClientMockRecorder) WriteSync(id, data, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteSync", reflect.TypeOf((*MockClient)(nil).WriteSync), id, data, timeout)
}

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// ReadMessage mocks base method.
func (m *MockConnection) ReadMessage() (int, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessage")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadMessage indicates an expected call of ReadMessage.
func (mr *MockConnectionMockRecorder) ReadMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockConnection)(nil).ReadMessage))
}

// RestoreConnection mocks base method.
func (m *MockConnection) RestoreConnection() (websocket.Connection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreConnection")
	ret0, _ := ret[0].(websocket.Connection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreConnection indicates an expected call of RestoreConnection.
func (mr *MockConnectionMockRecorder) RestoreConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreConnection", reflect.TypeOf((*MockConnection)(nil).RestoreConnection))
}

// WriteMessage mocks base method.
func (m *MockConnection) WriteMessage(messageType int, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessage", messageType, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessage indicates an expected call of WriteMessage.
func (mr *MockConnectionMockRecorder) WriteMessage(messageType, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessage", reflect.TypeOf((*MockConnection)(nil).WriteMessage), messageType, data)
}