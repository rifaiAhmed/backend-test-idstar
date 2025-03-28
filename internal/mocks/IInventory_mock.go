// Code generated by MockGen. DO NOT EDIT.
// Source: IInventory.go
//
// Generated by this command:
//
//	mockgen -source=IInventory.go -destination=../mocks/IInventory_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	models "backend-test/internal/models"
	context "context"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockIInventoryRepository is a mock of IInventoryRepository interface.
type MockIInventoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIInventoryRepositoryMockRecorder
	isgomock struct{}
}

// MockIInventoryRepositoryMockRecorder is the mock recorder for MockIInventoryRepository.
type MockIInventoryRepositoryMockRecorder struct {
	mock *MockIInventoryRepository
}

// NewMockIInventoryRepository creates a new mock instance.
func NewMockIInventoryRepository(ctrl *gomock.Controller) *MockIInventoryRepository {
	mock := &MockIInventoryRepository{ctrl: ctrl}
	mock.recorder = &MockIInventoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIInventoryRepository) EXPECT() *MockIInventoryRepositoryMockRecorder {
	return m.recorder
}

// BatchInsert mocks base method.
func (m *MockIInventoryRepository) BatchInsert(ctx context.Context, objs []models.Inventory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchInsert", ctx, objs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchInsert indicates an expected call of BatchInsert.
func (mr *MockIInventoryRepositoryMockRecorder) BatchInsert(ctx, objs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchInsert", reflect.TypeOf((*MockIInventoryRepository)(nil).BatchInsert), ctx, objs)
}

// CountData mocks base method.
func (m *MockIInventoryRepository) CountData(ctx context.Context, objComponent models.ComponentServerSide) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountData", ctx, objComponent)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountData indicates an expected call of CountData.
func (mr *MockIInventoryRepositoryMockRecorder) CountData(ctx, objComponent any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountData", reflect.TypeOf((*MockIInventoryRepository)(nil).CountData), ctx, objComponent)
}

// DeleteInv mocks base method.
func (m *MockIInventoryRepository) DeleteInv(ctx context.Context, ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInv", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInv indicates an expected call of DeleteInv.
func (mr *MockIInventoryRepositoryMockRecorder) DeleteInv(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInv", reflect.TypeOf((*MockIInventoryRepository)(nil).DeleteInv), ctx, ID)
}

// FindByID mocks base method.
func (m *MockIInventoryRepository) FindByID(ctx context.Context, ID int) (*models.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, ID)
	ret0, _ := ret[0].(*models.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockIInventoryRepositoryMockRecorder) FindByID(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockIInventoryRepository)(nil).FindByID), ctx, ID)
}

// GetAllInv mocks base method.
func (m *MockIInventoryRepository) GetAllInv(ctx context.Context, objComponent models.ComponentServerSide) (*[]models.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllInv", ctx, objComponent)
	ret0, _ := ret[0].(*[]models.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllInv indicates an expected call of GetAllInv.
func (mr *MockIInventoryRepositoryMockRecorder) GetAllInv(ctx, objComponent any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllInv", reflect.TypeOf((*MockIInventoryRepository)(nil).GetAllInv), ctx, objComponent)
}

// InsertInv mocks base method.
func (m *MockIInventoryRepository) InsertInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertInv", ctx, obj)
	ret0, _ := ret[0].(*models.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertInv indicates an expected call of InsertInv.
func (mr *MockIInventoryRepositoryMockRecorder) InsertInv(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertInv", reflect.TypeOf((*MockIInventoryRepository)(nil).InsertInv), ctx, obj)
}

// UpdateInv mocks base method.
func (m *MockIInventoryRepository) UpdateInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInv", ctx, obj)
	ret0, _ := ret[0].(*models.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInv indicates an expected call of UpdateInv.
func (mr *MockIInventoryRepositoryMockRecorder) UpdateInv(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInv", reflect.TypeOf((*MockIInventoryRepository)(nil).UpdateInv), ctx, obj)
}

// MockIInventoryService is a mock of IInventoryService interface.
type MockIInventoryService struct {
	ctrl     *gomock.Controller
	recorder *MockIInventoryServiceMockRecorder
	isgomock struct{}
}

// MockIInventoryServiceMockRecorder is the mock recorder for MockIInventoryService.
type MockIInventoryServiceMockRecorder struct {
	mock *MockIInventoryService
}

// NewMockIInventoryService creates a new mock instance.
func NewMockIInventoryService(ctrl *gomock.Controller) *MockIInventoryService {
	mock := &MockIInventoryService{ctrl: ctrl}
	mock.recorder = &MockIInventoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIInventoryService) EXPECT() *MockIInventoryServiceMockRecorder {
	return m.recorder
}

// CountData mocks base method.
func (m *MockIInventoryService) CountData(ctx context.Context, objComponent models.ComponentServerSide) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountData", ctx, objComponent)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountData indicates an expected call of CountData.
func (mr *MockIInventoryServiceMockRecorder) CountData(ctx, objComponent any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountData", reflect.TypeOf((*MockIInventoryService)(nil).CountData), ctx, objComponent)
}

// DeleteInv mocks base method.
func (m *MockIInventoryService) DeleteInv(ctx context.Context, ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInv", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInv indicates an expected call of DeleteInv.
func (mr *MockIInventoryServiceMockRecorder) DeleteInv(ctx, ID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInv", reflect.TypeOf((*MockIInventoryService)(nil).DeleteInv), ctx, ID)
}

// GetAllInv mocks base method.
func (m *MockIInventoryService) GetAllInv(ctx context.Context, objComponent models.ComponentServerSide) (*[]models.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllInv", ctx, objComponent)
	ret0, _ := ret[0].(*[]models.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllInv indicates an expected call of GetAllInv.
func (mr *MockIInventoryServiceMockRecorder) GetAllInv(ctx, objComponent any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllInv", reflect.TypeOf((*MockIInventoryService)(nil).GetAllInv), ctx, objComponent)
}

// InsertFromExcel mocks base method.
func (m *MockIInventoryService) InsertFromExcel(ctx context.Context, filePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertFromExcel", ctx, filePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertFromExcel indicates an expected call of InsertFromExcel.
func (mr *MockIInventoryServiceMockRecorder) InsertFromExcel(ctx, filePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertFromExcel", reflect.TypeOf((*MockIInventoryService)(nil).InsertFromExcel), ctx, filePath)
}

// InsertInv mocks base method.
func (m *MockIInventoryService) InsertInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertInv", ctx, obj)
	ret0, _ := ret[0].(*models.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertInv indicates an expected call of InsertInv.
func (mr *MockIInventoryServiceMockRecorder) InsertInv(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertInv", reflect.TypeOf((*MockIInventoryService)(nil).InsertInv), ctx, obj)
}

// UpdateInv mocks base method.
func (m *MockIInventoryService) UpdateInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInv", ctx, obj)
	ret0, _ := ret[0].(*models.Inventory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInv indicates an expected call of UpdateInv.
func (mr *MockIInventoryServiceMockRecorder) UpdateInv(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInv", reflect.TypeOf((*MockIInventoryService)(nil).UpdateInv), ctx, obj)
}

// MockIInventoryHandler is a mock of IInventoryHandler interface.
type MockIInventoryHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIInventoryHandlerMockRecorder
	isgomock struct{}
}

// MockIInventoryHandlerMockRecorder is the mock recorder for MockIInventoryHandler.
type MockIInventoryHandlerMockRecorder struct {
	mock *MockIInventoryHandler
}

// NewMockIInventoryHandler creates a new mock instance.
func NewMockIInventoryHandler(ctrl *gomock.Controller) *MockIInventoryHandler {
	mock := &MockIInventoryHandler{ctrl: ctrl}
	mock.recorder = &MockIInventoryHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIInventoryHandler) EXPECT() *MockIInventoryHandlerMockRecorder {
	return m.recorder
}

// DeleteInv mocks base method.
func (m *MockIInventoryHandler) DeleteInv(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteInv", c)
}

// DeleteInv indicates an expected call of DeleteInv.
func (mr *MockIInventoryHandlerMockRecorder) DeleteInv(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInv", reflect.TypeOf((*MockIInventoryHandler)(nil).DeleteInv), c)
}

// GetAllInv mocks base method.
func (m *MockIInventoryHandler) GetAllInv(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllInv", c)
}

// GetAllInv indicates an expected call of GetAllInv.
func (mr *MockIInventoryHandlerMockRecorder) GetAllInv(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllInv", reflect.TypeOf((*MockIInventoryHandler)(nil).GetAllInv), c)
}

// GetTemplate mocks base method.
func (m *MockIInventoryHandler) GetTemplate(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetTemplate", c)
}

// GetTemplate indicates an expected call of GetTemplate.
func (mr *MockIInventoryHandlerMockRecorder) GetTemplate(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplate", reflect.TypeOf((*MockIInventoryHandler)(nil).GetTemplate), c)
}

// InsertInv mocks base method.
func (m *MockIInventoryHandler) InsertInv(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertInv", c)
}

// InsertInv indicates an expected call of InsertInv.
func (mr *MockIInventoryHandlerMockRecorder) InsertInv(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertInv", reflect.TypeOf((*MockIInventoryHandler)(nil).InsertInv), c)
}

// UpdateInv mocks base method.
func (m *MockIInventoryHandler) UpdateInv(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateInv", c)
}

// UpdateInv indicates an expected call of UpdateInv.
func (mr *MockIInventoryHandlerMockRecorder) UpdateInv(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInv", reflect.TypeOf((*MockIInventoryHandler)(nil).UpdateInv), c)
}

// UploadExcel mocks base method.
func (m *MockIInventoryHandler) UploadExcel(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UploadExcel", c)
}

// UploadExcel indicates an expected call of UploadExcel.
func (mr *MockIInventoryHandlerMockRecorder) UploadExcel(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadExcel", reflect.TypeOf((*MockIInventoryHandler)(nil).UploadExcel), c)
}
