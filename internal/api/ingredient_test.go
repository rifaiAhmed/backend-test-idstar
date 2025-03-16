package api

import (
	"backend-test/constants"
	"backend-test/helpers"
	"backend-test/internal/mocks"
	"backend-test/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestIngredientAPI_InsertIngredient(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIIngredientService(ctrlMock)

	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	validIngredient := models.Ingredient{
		ID:          1,
		RecipeID:    2,
		InventoryID: 3,
		Quantity:    100.5,
	}

	tests := []struct {
		name       string
		request    any
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		{
			name:    "success - insert ingredient",
			request: validIngredient,
			mockFn: func() {
				mockService.EXPECT().
					InsertIngredient(gomock.Any(), gomock.Any()).
					Return(&validIngredient, nil).
					Times(1)
			},
			wantStatus: http.StatusOK,
			wantBody:   constants.SuccessMessage,
		},
		{
			name:       "fail - invalid JSON",
			request:    `{"invalid":}`,
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   constants.ErrFailedBadRequest,
		},
		{
			name: "fail - validation error",
			request: models.Ingredient{
				RecipeID:    0,
				InventoryID: 3,
				Quantity:    100.5,
			},
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   constants.ErrFailedBadRequest,
		},
		{
			name:    "fail - service returns error",
			request: validIngredient,
			mockFn: func() {
				mockService.EXPECT().
					InsertIngredient(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("failed to insert ingredient")).
					Times(1)
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "failed to insert ingredient",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &IngredientAPI{
				SvcIngredient: mockService,
			}

			var reqBody []byte
			if str, ok := tt.request.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, _ = json.Marshal(tt.request)
			}

			req, _ := http.NewRequest("POST", "/ingredient", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			tt.mockFn()

			api.InsertIngredient(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestIngredientAPI_UpdateIngredient(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIIngredientService(ctrlMock)

	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	validIngredient := models.Ingredient{
		ID:          1,
		RecipeID:    2,
		InventoryID: 3,
		Quantity:    100.5,
	}

	tests := []struct {
		name       string
		request    any
		uriID      string
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		{
			name:    "success - update ingredient",
			request: validIngredient,
			uriID:   "1",
			mockFn: func() {
				mockService.EXPECT().
					UpdateIngredient(gomock.Any(), gomock.Any()).
					Return(&validIngredient, nil).
					Times(1)
			},
			wantStatus: http.StatusOK,
			wantBody:   constants.SuccessMessage,
		},
		{
			name:       "fail - invalid JSON",
			request:    `{"invalid":}`,
			uriID:      "1",
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   constants.ErrFailedBadRequest,
		},
		{
			name: "fail - validation error",
			request: models.Ingredient{
				RecipeID:    0,
				InventoryID: 3,
				Quantity:    100.5,
			},
			uriID:      "1",
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   constants.ErrFailedBadRequest,
		},
		{
			name:       "fail - invalid URI ID",
			request:    validIngredient,
			uriID:      "abc",
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   constants.ErrFailedBadRequest,
		},
		{
			name:    "fail - service returns error",
			request: validIngredient,
			uriID:   "1",
			mockFn: func() {
				mockService.EXPECT().
					UpdateIngredient(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("failed to update ingredient")).
					Times(1)
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "failed to update ingredient",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &IngredientAPI{
				SvcIngredient: mockService,
			}

			var reqBody []byte
			if str, ok := tt.request.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, _ = json.Marshal(tt.request)
			}

			req, _ := http.NewRequest("PUT", "/ingredient/"+tt.uriID, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = append(c.Params, gin.Param{Key: "id", Value: tt.uriID})
			c.Request = req

			tt.mockFn()

			api.UpdateIngredient(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestIngredientAPI_DeleteIngredient(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIIngredientService(ctrlMock)

	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		uriID      string
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		{
			name:  "success - delete ingredient",
			uriID: "1",
			mockFn: func() {
				mockService.EXPECT().
					DeleteIngredient(gomock.Any(), 1).
					Return(nil).
					Times(1)
			},
			wantStatus: http.StatusOK,
			wantBody:   constants.SuccessMessage,
		},
		{
			name:       "fail - invalid URI ID",
			uriID:      "abc",
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   constants.ErrFailedBadRequest,
		},
		{
			name:  "fail - service returns error",
			uriID: "2",
			mockFn: func() {
				mockService.EXPECT().
					DeleteIngredient(gomock.Any(), 2).
					Return(errors.New("failed to delete ingredient")).
					Times(1)
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "failed to delete ingredient",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &IngredientAPI{
				SvcIngredient: mockService,
			}

			req, _ := http.NewRequest("DELETE", "/ingredient/"+tt.uriID, nil)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = append(c.Params, gin.Param{Key: "id", Value: tt.uriID})
			c.Request = req

			tt.mockFn()

			api.DeleteIngredient(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestIngredientAPI_GetAllIngredient(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIIngredientService(ctrlMock)

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
			api := &IngredientAPI{
				SvcIngredient: mockService,
			}

			req, _ := http.NewRequest("GET", "/ingredient?search="+tt.search, nil)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			tt.mockFn()

			api.GetAllIngredient(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestIngredientAPI_GetRecipeIncludeIngredients(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIIngredientService(ctrlMock)

	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		queryID    string
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		{
			name:    "success - get recipe with ingredients",
			queryID: "1",
			mockFn: func() {
				mockService.EXPECT().
					GetRecipeIncludeIngredients(gomock.Any(), 1).
					Return(models.RecipeFormat{
						ID:   1,
						Name: "Pasta",
					}, &[]models.IngredientCustom{
						{ID: 1, InventoryID: 1, Quantity: 2, Item: "item 1"},
						{ID: 2, InventoryID: 2, Quantity: 2, Item: "Item 2"},
					}, nil).
					Times(1)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"message":"Successfully get data"`,
		},
		{
			name:       "fail - invalid id parameter",
			queryID:    "abc",
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   `"message":"data tidak sesuai"`, // Disesuaikan dengan API
		},
		{
			name:    "fail - service returns error",
			queryID: "1",
			mockFn: func() {
				mockService.EXPECT().
					GetRecipeIncludeIngredients(gomock.Any(), 1).
					Return(models.RecipeFormat{}, nil, errors.New("recipe not found")).
					Times(1)
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `"message":"data tidak sesuai"`, // Disesuaikan dengan API
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &IngredientAPI{
				SvcIngredient: mockService,
			}

			req, _ := http.NewRequest("GET", "/ingredient/recipe?id="+tt.queryID, nil)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			tt.mockFn()

			api.GetRecipeIncludeIngredients(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody) // Disesuaikan dengan API
		})
	}
}

func TestIngredientAPI_MultipleCreateUpdate(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIIngredientService(ctrlMock)

	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name        string
		requestBody string
		mockFn      func()
		wantStatus  int
		wantBody    string
	}{
		{
			name: "success - multiple create/update ingredients",
			requestBody: `{
				"id": 1,
				"data": "0|1|2,0|1|3"
			}`,
			mockFn: func() {
				mockService.EXPECT().
					MultipleCreateUpdate(gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"message":"success"`,
		},
		{
			name:        "fail - missing required fields",
			requestBody: `{}`,
			mockFn:      func() {},
			wantStatus:  http.StatusBadRequest,
			wantBody:    `"message":"data tidak sesuai"`,
		},
		{
			name: "fail - service returns error",
			requestBody: `{
				"id": 1,
				"data": "0|1|2,0|1|3"
			}`,
			mockFn: func() {
				mockService.EXPECT().
					MultipleCreateUpdate(gomock.Any(), gomock.Any()).
					Return(errors.New("failed to create ingredient")).
					Times(1)
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `"message":"data tidak sesuai"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &IngredientAPI{
				SvcIngredient: mockService,
			}

			req, _ := http.NewRequest("POST", "/ingredient/multiple", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			tt.mockFn()

			api.MultipleCreateUpdate(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}
