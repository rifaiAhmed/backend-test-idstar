package api

import (
	"backend-test/helpers"
	"backend-test/internal/mocks"
	"backend-test/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestIUserHandler_SendMail(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIUserService(ctrlMock)
	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	validUser := &models.User{
		Email: "test@gmail.com",
	}

	tests := []struct {
		name       string
		request    any
		mockFn     func()
		wantStatus int
		wantBody   string
	}{
		{
			name:    "Success - Send Mail",
			request: validUser,
			mockFn: func() {
				mockService.EXPECT().
					SubmitEmail(gomock.Any(), gomock.Any()).
					Return(validUser, "token123", nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"success"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &IUserHandler{
				UserService: mockService,
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			reqBody, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest("POST", "/auth/submit-email", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tt.mockFn()

			api.SendMail(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestIUserHandler_CekSessionByUUID(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := mocks.NewMockIUserService(ctrlMock)
	helpers.Logger = logrus.New()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		request    string
		mockFn     func()
		wantStatus int
		wantBody   string
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &IUserHandler{
				UserService: mockService,
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req, _ := http.NewRequest("GET", "/auth/session/"+tt.request, nil)
			c.Params = []gin.Param{{Key: "uuid", Value: tt.request}}
			c.Request = req

			tt.mockFn()

			api.CekSessionByUUID(c)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}
