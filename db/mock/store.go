// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Rich-dev-team/stock-prediction-backend/db/sqlc (interfaces: Store)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
	context "context"
	reflect "reflect"

	db "github.com/Rich-dev-team/stock-prediction-backend/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateBatchStockPrice mocks base method.
func (m *MockStore) CreateBatchStockPrice(arg0 context.Context, arg1 db.CreateBatchStockPriceParams) ([]db.StockPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBatchStockPrice", arg0, arg1)
	ret0, _ := ret[0].([]db.StockPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBatchStockPrice indicates an expected call of CreateBatchStockPrice.
func (mr *MockStoreMockRecorder) CreateBatchStockPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBatchStockPrice", reflect.TypeOf((*MockStore)(nil).CreateBatchStockPrice), arg0, arg1)
}

// CreateCompany mocks base method.
func (m *MockStore) CreateCompany(arg0 context.Context, arg1 db.CreateCompanyParams) (db.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", arg0, arg1)
	ret0, _ := ret[0].(db.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockStoreMockRecorder) CreateCompany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockStore)(nil).CreateCompany), arg0, arg1)
}

// CreateStockPrice mocks base method.
func (m *MockStore) CreateStockPrice(arg0 context.Context, arg1 db.CreateStockPriceParams) (db.StockPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStockPrice", arg0, arg1)
	ret0, _ := ret[0].(db.StockPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStockPrice indicates an expected call of CreateStockPrice.
func (mr *MockStoreMockRecorder) CreateStockPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStockPrice", reflect.TypeOf((*MockStore)(nil).CreateStockPrice), arg0, arg1)
}

// CreateStocksTx mocks base method.
func (m *MockStore) CreateStocksTx(arg0 context.Context, arg1 db.CreateBatchStockPriceParams) ([]db.StockPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStocksTx", arg0, arg1)
	ret0, _ := ret[0].([]db.StockPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStocksTx indicates an expected call of CreateStocksTx.
func (mr *MockStoreMockRecorder) CreateStocksTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStocksTx", reflect.TypeOf((*MockStore)(nil).CreateStocksTx), arg0, arg1)
}

// DeleteCompany mocks base method.
func (m *MockStore) DeleteCompany(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockStoreMockRecorder) DeleteCompany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockStore)(nil).DeleteCompany), arg0, arg1)
}

// DeleteStockPrice mocks base method.
func (m *MockStore) DeleteStockPrice(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStockPrice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStockPrice indicates an expected call of DeleteStockPrice.
func (mr *MockStoreMockRecorder) DeleteStockPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStockPrice", reflect.TypeOf((*MockStore)(nil).DeleteStockPrice), arg0, arg1)
}

// GetCompanyById mocks base method.
func (m *MockStore) GetCompanyById(arg0 context.Context, arg1 int64) (db.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyById", arg0, arg1)
	ret0, _ := ret[0].(db.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyById indicates an expected call of GetCompanyById.
func (mr *MockStoreMockRecorder) GetCompanyById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyById", reflect.TypeOf((*MockStore)(nil).GetCompanyById), arg0, arg1)
}

// GetCompanyByName mocks base method.
func (m *MockStore) GetCompanyByName(arg0 context.Context, arg1 string) (db.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyByName", arg0, arg1)
	ret0, _ := ret[0].(db.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyByName indicates an expected call of GetCompanyByName.
func (mr *MockStoreMockRecorder) GetCompanyByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyByName", reflect.TypeOf((*MockStore)(nil).GetCompanyByName), arg0, arg1)
}

// ListAllStockPrice mocks base method.
func (m *MockStore) ListAllStockPrice(arg0 context.Context, arg1 db.ListAllStockPriceParams) ([]db.StockPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllStockPrice", arg0, arg1)
	ret0, _ := ret[0].([]db.StockPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllStockPrice indicates an expected call of ListAllStockPrice.
func (mr *MockStoreMockRecorder) ListAllStockPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllStockPrice", reflect.TypeOf((*MockStore)(nil).ListAllStockPrice), arg0, arg1)
}

// ListCompany mocks base method.
func (m *MockStore) ListCompany(arg0 context.Context, arg1 db.ListCompanyParams) ([]db.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCompany", arg0, arg1)
	ret0, _ := ret[0].([]db.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCompany indicates an expected call of ListCompany.
func (mr *MockStoreMockRecorder) ListCompany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCompany", reflect.TypeOf((*MockStore)(nil).ListCompany), arg0, arg1)
}

// ListStockPriceByRange mocks base method.
func (m *MockStore) ListStockPriceByRange(arg0 context.Context, arg1 db.ListStockPriceByRangeParams) ([]db.StockPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStockPriceByRange", arg0, arg1)
	ret0, _ := ret[0].([]db.StockPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStockPriceByRange indicates an expected call of ListStockPriceByRange.
func (mr *MockStoreMockRecorder) ListStockPriceByRange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStockPriceByRange", reflect.TypeOf((*MockStore)(nil).ListStockPriceByRange), arg0, arg1)
}

// ListStockPriceByRangeForUpdate mocks base method.
func (m *MockStore) ListStockPriceByRangeForUpdate(arg0 context.Context, arg1 db.ListStockPriceByRangeForUpdateParams) ([]db.StockPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStockPriceByRangeForUpdate", arg0, arg1)
	ret0, _ := ret[0].([]db.StockPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStockPriceByRangeForUpdate indicates an expected call of ListStockPriceByRangeForUpdate.
func (mr *MockStoreMockRecorder) ListStockPriceByRangeForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStockPriceByRangeForUpdate", reflect.TypeOf((*MockStore)(nil).ListStockPriceByRangeForUpdate), arg0, arg1)
}

// UpdateCompanyName mocks base method.
func (m *MockStore) UpdateCompanyName(arg0 context.Context, arg1 db.UpdateCompanyNameParams) (db.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompanyName", arg0, arg1)
	ret0, _ := ret[0].(db.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCompanyName indicates an expected call of UpdateCompanyName.
func (mr *MockStoreMockRecorder) UpdateCompanyName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompanyName", reflect.TypeOf((*MockStore)(nil).UpdateCompanyName), arg0, arg1)
}

// UpdateCompanyStockSymbol mocks base method.
func (m *MockStore) UpdateCompanyStockSymbol(arg0 context.Context, arg1 db.UpdateCompanyStockSymbolParams) (db.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompanyStockSymbol", arg0, arg1)
	ret0, _ := ret[0].(db.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCompanyStockSymbol indicates an expected call of UpdateCompanyStockSymbol.
func (mr *MockStoreMockRecorder) UpdateCompanyStockSymbol(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompanyStockSymbol", reflect.TypeOf((*MockStore)(nil).UpdateCompanyStockSymbol), arg0, arg1)
}

// UpdateStockPrice mocks base method.
func (m *MockStore) UpdateStockPrice(arg0 context.Context, arg1 db.UpdateStockPriceParams) (db.StockPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStockPrice", arg0, arg1)
	ret0, _ := ret[0].(db.StockPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStockPrice indicates an expected call of UpdateStockPrice.
func (mr *MockStoreMockRecorder) UpdateStockPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStockPrice", reflect.TypeOf((*MockStore)(nil).UpdateStockPrice), arg0, arg1)
}