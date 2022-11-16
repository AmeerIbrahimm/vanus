// Code generated by MockGen. DO NOT EDIT.
// Source: entry.go

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	block "github.com/linkall-labs/vanus/internal/store/block"
)

// MockExtensionAttributeCallback is a mock of ExtensionAttributeCallback interface.
type MockExtensionAttributeCallback struct {
	ctrl     *gomock.Controller
	recorder *MockExtensionAttributeCallbackMockRecorder
}

// MockExtensionAttributeCallbackMockRecorder is the mock recorder for MockExtensionAttributeCallback.
type MockExtensionAttributeCallbackMockRecorder struct {
	mock *MockExtensionAttributeCallback
}

// NewMockExtensionAttributeCallback creates a new mock instance.
func NewMockExtensionAttributeCallback(ctrl *gomock.Controller) *MockExtensionAttributeCallback {
	mock := &MockExtensionAttributeCallback{ctrl: ctrl}
	mock.recorder = &MockExtensionAttributeCallbackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtensionAttributeCallback) EXPECT() *MockExtensionAttributeCallbackMockRecorder {
	return m.recorder
}

// OnAttribute mocks base method.
func (m *MockExtensionAttributeCallback) OnAttribute(attr, val []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnAttribute", attr, val)
}

// OnAttribute indicates an expected call of OnAttribute.
func (mr *MockExtensionAttributeCallbackMockRecorder) OnAttribute(attr, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnAttribute", reflect.TypeOf((*MockExtensionAttributeCallback)(nil).OnAttribute), attr, val)
}

// MockEntry is a mock of Entry interface.
type MockEntry struct {
	ctrl     *gomock.Controller
	recorder *MockEntryMockRecorder
}

// MockEntryMockRecorder is the mock recorder for MockEntry.
type MockEntryMockRecorder struct {
	mock *MockEntry
}

// NewMockEntry creates a new mock instance.
func NewMockEntry(ctrl *gomock.Controller) *MockEntry {
	mock := &MockEntry{ctrl: ctrl}
	mock.recorder = &MockEntryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntry) EXPECT() *MockEntryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockEntry) Get(ordinal int) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ordinal)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockEntryMockRecorder) Get(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEntry)(nil).Get), ordinal)
}

// GetBytes mocks base method.
func (m *MockEntry) GetBytes(ordinal int) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBytes", ordinal)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetBytes indicates an expected call of GetBytes.
func (mr *MockEntryMockRecorder) GetBytes(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBytes", reflect.TypeOf((*MockEntry)(nil).GetBytes), ordinal)
}

// GetExtensionAttribute mocks base method.
func (m *MockEntry) GetExtensionAttribute(arg0 []byte) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtensionAttribute", arg0)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetExtensionAttribute indicates an expected call of GetExtensionAttribute.
func (mr *MockEntryMockRecorder) GetExtensionAttribute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtensionAttribute", reflect.TypeOf((*MockEntry)(nil).GetExtensionAttribute), arg0)
}

// GetInt64 mocks base method.
func (m *MockEntry) GetInt64(ordinal int) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt64", ordinal)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetInt64 indicates an expected call of GetInt64.
func (mr *MockEntryMockRecorder) GetInt64(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt64", reflect.TypeOf((*MockEntry)(nil).GetInt64), ordinal)
}

// GetString mocks base method.
func (m *MockEntry) GetString(ordinal int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetString", ordinal)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetString indicates an expected call of GetString.
func (mr *MockEntryMockRecorder) GetString(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockEntry)(nil).GetString), ordinal)
}

// GetTime mocks base method.
func (m *MockEntry) GetTime(ordinal int) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTime", ordinal)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTime indicates an expected call of GetTime.
func (mr *MockEntryMockRecorder) GetTime(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*MockEntry)(nil).GetTime), ordinal)
}

// GetUint16 mocks base method.
func (m *MockEntry) GetUint16(ordinal int) uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUint16", ordinal)
	ret0, _ := ret[0].(uint16)
	return ret0
}

// GetUint16 indicates an expected call of GetUint16.
func (mr *MockEntryMockRecorder) GetUint16(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUint16", reflect.TypeOf((*MockEntry)(nil).GetUint16), ordinal)
}

// GetUint64 mocks base method.
func (m *MockEntry) GetUint64(ordinal int) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUint64", ordinal)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetUint64 indicates an expected call of GetUint64.
func (mr *MockEntryMockRecorder) GetUint64(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUint64", reflect.TypeOf((*MockEntry)(nil).GetUint64), ordinal)
}

// RangeExtensionAttributes mocks base method.
func (m *MockEntry) RangeExtensionAttributes(cb block.ExtensionAttributeCallback) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RangeExtensionAttributes", cb)
}

// RangeExtensionAttributes indicates an expected call of RangeExtensionAttributes.
func (mr *MockEntryMockRecorder) RangeExtensionAttributes(cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeExtensionAttributes", reflect.TypeOf((*MockEntry)(nil).RangeExtensionAttributes), cb)
}

// MockOptionalAttributeCallback is a mock of OptionalAttributeCallback interface.
type MockOptionalAttributeCallback struct {
	ctrl     *gomock.Controller
	recorder *MockOptionalAttributeCallbackMockRecorder
}

// MockOptionalAttributeCallbackMockRecorder is the mock recorder for MockOptionalAttributeCallback.
type MockOptionalAttributeCallbackMockRecorder struct {
	mock *MockOptionalAttributeCallback
}

// NewMockOptionalAttributeCallback creates a new mock instance.
func NewMockOptionalAttributeCallback(ctrl *gomock.Controller) *MockOptionalAttributeCallback {
	mock := &MockOptionalAttributeCallback{ctrl: ctrl}
	mock.recorder = &MockOptionalAttributeCallbackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOptionalAttributeCallback) EXPECT() *MockOptionalAttributeCallbackMockRecorder {
	return m.recorder
}

// OnAttribute mocks base method.
func (m *MockOptionalAttributeCallback) OnAttribute(ordinal int, val interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnAttribute", ordinal, val)
}

// OnAttribute indicates an expected call of OnAttribute.
func (mr *MockOptionalAttributeCallbackMockRecorder) OnAttribute(ordinal, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnAttribute", reflect.TypeOf((*MockOptionalAttributeCallback)(nil).OnAttribute), ordinal, val)
}

// OnBytes mocks base method.
func (m *MockOptionalAttributeCallback) OnBytes(ordinal int, val []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnBytes", ordinal, val)
}

// OnBytes indicates an expected call of OnBytes.
func (mr *MockOptionalAttributeCallbackMockRecorder) OnBytes(ordinal, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnBytes", reflect.TypeOf((*MockOptionalAttributeCallback)(nil).OnBytes), ordinal, val)
}

// OnInt64 mocks base method.
func (m *MockOptionalAttributeCallback) OnInt64(ordinal int, val int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnInt64", ordinal, val)
}

// OnInt64 indicates an expected call of OnInt64.
func (mr *MockOptionalAttributeCallbackMockRecorder) OnInt64(ordinal, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnInt64", reflect.TypeOf((*MockOptionalAttributeCallback)(nil).OnInt64), ordinal, val)
}

// OnString mocks base method.
func (m *MockOptionalAttributeCallback) OnString(ordinal int, val string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnString", ordinal, val)
}

// OnString indicates an expected call of OnString.
func (mr *MockOptionalAttributeCallbackMockRecorder) OnString(ordinal, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnString", reflect.TypeOf((*MockOptionalAttributeCallback)(nil).OnString), ordinal, val)
}

// OnTime mocks base method.
func (m *MockOptionalAttributeCallback) OnTime(ordinal int, val time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnTime", ordinal, val)
}

// OnTime indicates an expected call of OnTime.
func (mr *MockOptionalAttributeCallbackMockRecorder) OnTime(ordinal, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnTime", reflect.TypeOf((*MockOptionalAttributeCallback)(nil).OnTime), ordinal, val)
}

// OnUint16 mocks base method.
func (m *MockOptionalAttributeCallback) OnUint16(ordinal int, val uint16) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnUint16", ordinal, val)
}

// OnUint16 indicates an expected call of OnUint16.
func (mr *MockOptionalAttributeCallbackMockRecorder) OnUint16(ordinal, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUint16", reflect.TypeOf((*MockOptionalAttributeCallback)(nil).OnUint16), ordinal, val)
}

// OnUint64 mocks base method.
func (m *MockOptionalAttributeCallback) OnUint64(ordinal int, val uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnUint64", ordinal, val)
}

// OnUint64 indicates an expected call of OnUint64.
func (mr *MockOptionalAttributeCallbackMockRecorder) OnUint64(ordinal, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUint64", reflect.TypeOf((*MockOptionalAttributeCallback)(nil).OnUint64), ordinal, val)
}

// MockEntryExt is a mock of EntryExt interface.
type MockEntryExt struct {
	ctrl     *gomock.Controller
	recorder *MockEntryExtMockRecorder
}

// MockEntryExtMockRecorder is the mock recorder for MockEntryExt.
type MockEntryExtMockRecorder struct {
	mock *MockEntryExt
}

// NewMockEntryExt creates a new mock instance.
func NewMockEntryExt(ctrl *gomock.Controller) *MockEntryExt {
	mock := &MockEntryExt{ctrl: ctrl}
	mock.recorder = &MockEntryExtMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntryExt) EXPECT() *MockEntryExtMockRecorder {
	return m.recorder
}

// ExtensionAttributeCount mocks base method.
func (m *MockEntryExt) ExtensionAttributeCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtensionAttributeCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// ExtensionAttributeCount indicates an expected call of ExtensionAttributeCount.
func (mr *MockEntryExtMockRecorder) ExtensionAttributeCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtensionAttributeCount", reflect.TypeOf((*MockEntryExt)(nil).ExtensionAttributeCount))
}

// Get mocks base method.
func (m *MockEntryExt) Get(ordinal int) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ordinal)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockEntryExtMockRecorder) Get(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEntryExt)(nil).Get), ordinal)
}

// GetBytes mocks base method.
func (m *MockEntryExt) GetBytes(ordinal int) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBytes", ordinal)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetBytes indicates an expected call of GetBytes.
func (mr *MockEntryExtMockRecorder) GetBytes(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBytes", reflect.TypeOf((*MockEntryExt)(nil).GetBytes), ordinal)
}

// GetExtensionAttribute mocks base method.
func (m *MockEntryExt) GetExtensionAttribute(arg0 []byte) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtensionAttribute", arg0)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetExtensionAttribute indicates an expected call of GetExtensionAttribute.
func (mr *MockEntryExtMockRecorder) GetExtensionAttribute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtensionAttribute", reflect.TypeOf((*MockEntryExt)(nil).GetExtensionAttribute), arg0)
}

// GetInt64 mocks base method.
func (m *MockEntryExt) GetInt64(ordinal int) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInt64", ordinal)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetInt64 indicates an expected call of GetInt64.
func (mr *MockEntryExtMockRecorder) GetInt64(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt64", reflect.TypeOf((*MockEntryExt)(nil).GetInt64), ordinal)
}

// GetString mocks base method.
func (m *MockEntryExt) GetString(ordinal int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetString", ordinal)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetString indicates an expected call of GetString.
func (mr *MockEntryExtMockRecorder) GetString(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockEntryExt)(nil).GetString), ordinal)
}

// GetTime mocks base method.
func (m *MockEntryExt) GetTime(ordinal int) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTime", ordinal)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTime indicates an expected call of GetTime.
func (mr *MockEntryExtMockRecorder) GetTime(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*MockEntryExt)(nil).GetTime), ordinal)
}

// GetUint16 mocks base method.
func (m *MockEntryExt) GetUint16(ordinal int) uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUint16", ordinal)
	ret0, _ := ret[0].(uint16)
	return ret0
}

// GetUint16 indicates an expected call of GetUint16.
func (mr *MockEntryExtMockRecorder) GetUint16(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUint16", reflect.TypeOf((*MockEntryExt)(nil).GetUint16), ordinal)
}

// GetUint64 mocks base method.
func (m *MockEntryExt) GetUint64(ordinal int) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUint64", ordinal)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetUint64 indicates an expected call of GetUint64.
func (mr *MockEntryExtMockRecorder) GetUint64(ordinal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUint64", reflect.TypeOf((*MockEntryExt)(nil).GetUint64), ordinal)
}

// OptionalAttributeCount mocks base method.
func (m *MockEntryExt) OptionalAttributeCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OptionalAttributeCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// OptionalAttributeCount indicates an expected call of OptionalAttributeCount.
func (mr *MockEntryExtMockRecorder) OptionalAttributeCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OptionalAttributeCount", reflect.TypeOf((*MockEntryExt)(nil).OptionalAttributeCount))
}

// RangeExtensionAttributes mocks base method.
func (m *MockEntryExt) RangeExtensionAttributes(cb block.ExtensionAttributeCallback) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RangeExtensionAttributes", cb)
}

// RangeExtensionAttributes indicates an expected call of RangeExtensionAttributes.
func (mr *MockEntryExtMockRecorder) RangeExtensionAttributes(cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeExtensionAttributes", reflect.TypeOf((*MockEntryExt)(nil).RangeExtensionAttributes), cb)
}

// RangeOptionalAttributes mocks base method.
func (m *MockEntryExt) RangeOptionalAttributes(cb block.OptionalAttributeCallback) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RangeOptionalAttributes", cb)
}

// RangeOptionalAttributes indicates an expected call of RangeOptionalAttributes.
func (mr *MockEntryExtMockRecorder) RangeOptionalAttributes(cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeOptionalAttributes", reflect.TypeOf((*MockEntryExt)(nil).RangeOptionalAttributes), cb)
}
