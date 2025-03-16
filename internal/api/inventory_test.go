package api

import (
	"backend-test/helpers"
	"backend-test/internal/mocks"
	"backend-test/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInventoryAPI_InsertInv(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIInventoryService(ctrl)
	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		body       models.Inventory
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		{
			name: "success - insert inventory",
			body: models.Inventory{
				Item:        "item 1",
				Qty:         2,
				Uom:         "Pcs",
				PricePerQty: 1000,
			},
			mockFn: func() {
				mockService.EXPECT().
					InsertInv(gomock.Any(), gomock.Any()).
					Return(&models.Inventory{
						ID:          1,
						Item:        "item 1",
						Qty:         2,
						Uom:         "Pcs",
						PricePerQty: 1000,
					}, nil).
					Times(1)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"message":"success"`,
		},
		{
			name:       "fail - invalid request body",
			body:       models.Inventory{},
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   `"message":"data tidak sesuai"`,
		},
		{
			name: "fail - service error",
			body: models.Inventory{
				Item:        "item 1",
				Qty:         1,
				Uom:         "Pcs",
				PricePerQty: 1000,
			},
			mockFn: func() {
				mockService.EXPECT().
					InsertInv(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("failed to create inventory")).
					Times(1)
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `"message":"data tidak sesuai"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &InventoryAPI{InventoryService: mockService}

			bodyBytes, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/inventory", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			tt.mockFn()

			api.InsertInv(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestInventoryAPI_UpdateInv(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIInventoryService(ctrl)
	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		body       models.Inventory
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		{
			name: "success - update inventory",
			body: models.Inventory{
				ID:          1,
				Item:        "item 1",
				Qty:         2,
				Uom:         "Pcs",
				PricePerQty: 1000,
			},
			mockFn: func() {
				mockService.EXPECT().
					UpdateInv(gomock.Any(), gomock.Any()).
					Return(&models.Inventory{
						ID:          1,
						Item:        "item 1",
						Qty:         2,
						Uom:         "Pcs",
						PricePerQty: 1000,
					}, nil).
					Times(1)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"message":"success"`,
		},
		{
			name:       "fail - invalid request body",
			body:       models.Inventory{},
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   `"message":"data tidak sesuai"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &InventoryAPI{InventoryService: mockService}

			bodyBytes, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("PUT", "/inventory/1", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{{Key: "id", Value: "1"}}
			c.Request = req

			tt.mockFn()

			api.UpdateInv(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestInventoryAPI_DeleteInv(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIInventoryService(ctrl)
	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		id         string
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		{
			name: "success - delete inventory",
			id:   "1",
			mockFn: func() {
				mockService.EXPECT().
					DeleteInv(gomock.Any(), 1).
					Return(nil).
					Times(1)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"message":"success"`,
		},
		{
			name:       "fail - invalid ID",
			id:         "abc",
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   `"message":"data tidak sesuai"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &InventoryAPI{InventoryService: mockService}

			req, _ := http.NewRequest("DELETE", "/inventory/"+tt.id, nil)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{{Key: "id", Value: tt.id}}
			c.Request = req

			tt.mockFn()

			api.DeleteInv(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestInventoryAPI_GetAllInv(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIInventoryService(ctrl)
	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		search     string
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		//
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &InventoryAPI{
				InventoryService: mockService,
			}

			req, _ := http.NewRequest("GET", "/ingredient?search="+tt.search, nil)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			tt.mockFn()

			api.GetAllInv(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}
