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
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestReceipeAPi_InsertRecipe(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIRecipeService(ctrlMock)
	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	validRecipe := models.Recipe{
		ID:   1,
		Name: "recipe 1",
	}

	tests := []struct {
		name       string
		request    any
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		{
			name:    "Success - Insert Recipe",
			request: validRecipe,
			mockFn: func() {
				mockService.EXPECT().InsertRecipe(gomock.Any(), gomock.Any()).Return(&validRecipe, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   constants.SuccessMessage,
		},
		{
			name:       "Fail - Invalid Request Body",
			request:    "invalid-json",
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   constants.ErrFailedBadRequest,
		},
		{
			name:    "Fail - Service Error",
			request: validRecipe,
			mockFn: func() {
				mockService.EXPECT().InsertRecipe(gomock.Any(), gomock.Any()).Return(nil, errors.New("service error"))
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "service error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &ReceipeAPi{
				SvcRecipe: mockService,
			}
			var reqBody []byte
			if str, ok := tt.request.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, _ = json.Marshal(tt.request)
			}

			req, _ := http.NewRequest("POST", "/recipe", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			tt.mockFn()

			api.InsertRecipe(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestReceipeAPi_UpdateRecipe(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIRecipeService(ctrlMock)

	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	validRecipe := models.Recipe{
		ID:   1,
		Name: "Updated Recipe",
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
			name:    "Success Update Recipe",
			request: validRecipe,
			uriID:   "1",
			mockFn: func() {
				mockService.EXPECT().UpdateRecipe(gomock.Any(), gomock.Any()).Return(&validRecipe, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   "success",
		},
		{
			name:       "Invalid ID in URI",
			request:    validRecipe,
			uriID:      "abc",
			mockFn:     func() {}, // Tidak perlu mock karena akan gagal parsing ID
			wantStatus: http.StatusBadRequest,
			wantBody:   constants.ErrFailedBadRequest,
		},
		{
			name:    "Service Error on Update",
			request: validRecipe,
			uriID:   "1",
			mockFn: func() {
				mockService.EXPECT().UpdateRecipe(gomock.Any(), gomock.Any()).Return(nil, errors.New("failed to update recipe"))
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "failed to update recipe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &ReceipeAPi{SvcRecipe: mockService}

			var reqBody []byte
			reqBody, _ = json.Marshal(tt.request)

			req, _ := http.NewRequest("PUT", "/recipe/"+tt.uriID, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = append(c.Params, gin.Param{Key: "id", Value: tt.uriID})
			c.Request = req

			tt.mockFn()

			api.UpdateRecipe(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestReceipeAPi_DeleteRecipe(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIRecipeService(ctrlMock)
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
			name:  "Success Delete Recipe",
			uriID: "1",
			mockFn: func() {
				mockService.EXPECT().DeleteRecipe(gomock.Any(), 1).Return(nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   constants.SuccessMessage,
		},
		{
			name:       "Invalid ID",
			uriID:      "abc",
			mockFn:     func() {},
			wantStatus: http.StatusBadRequest,
			wantBody:   constants.ErrFailedBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &ReceipeAPi{SvcRecipe: mockService}

			req, _ := http.NewRequest("DELETE", "/recipe/"+tt.uriID, nil)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = append(c.Params, gin.Param{Key: "id", Value: tt.uriID})
			c.Request = req

			tt.mockFn()

			api.DeleteRecipe(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestReceipeAPi_GetAllRecipe(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIRecipeService(ctrlMock)

	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	mockRecipes := []models.RecipeFormat{
		{ID: 1, Name: "Recipe 1", SKU: "SKU-123"},
		{ID: 2, Name: "Recipe 2", SKU: "SKU-456"},
	}

	tests := []struct {
		name       string
		query      string
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		{
			name:  "Success Get All Recipes",
			query: "?search=",
			mockFn: func() {
				mockService.EXPECT().CountData(gomock.Any(), gomock.Any()).Return(int64(2), nil)

				mockService.EXPECT().GetAllRecipe(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockRecipes, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   "Succesfully Get Data!",
		},
		{
			name:  "Service Error on GetAllRecipe",
			query: "?search=",
			mockFn: func() {
				mockService.EXPECT().CountData(gomock.Any(), gomock.Any()).Return(int64(2), nil)

				mockService.EXPECT().GetAllRecipe(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("failed to get recipes"))
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "failed to get recipes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &ReceipeAPi{SvcRecipe: mockService}

			req, _ := http.NewRequest("GET", "/recipe"+tt.query, nil)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			tt.mockFn()

			api.GetAllRecipe(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}
