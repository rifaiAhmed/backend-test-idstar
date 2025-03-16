package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"

	"backend-test/internal/mocks"
	"backend-test/internal/models"
)

func TestRegisterService_CekSessionByUUID(t *testing.T) {
	uuidD := uuid.New()
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrlMock)

	type args struct {
		ctx  context.Context
		uuid uuid.UUID
	}

	tests := []struct {
		name    string
		args    args
		want    *models.UserSession
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - session found",
			args: args{
				ctx:  context.Background(),
				uuid: uuidD,
			},
			want: &models.UserSession{
				UserID: 1,
				Uuid:   uuidD,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					CekSessionByUUID(args.ctx, args.uuid).
					Return(&models.UserSession{
						UserID:       1,
						Uuid:         args.uuid,
						CreatedAt:    time.Now(),
						TokenExpired: time.Now().Add(24 * time.Hour),
					}, nil)
			},
		},
		{
			name: "error - session not found",
			args: args{
				ctx:  context.Background(),
				uuid: uuidD,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					CekSessionByUUID(args.ctx, args.uuid).
					Return(nil, gorm.ErrRecordNotFound)
			},
		},
		{
			name: "error - database failure",
			args: args{
				ctx:  context.Background(),
				uuid: uuidD,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					CekSessionByUUID(args.ctx, args.uuid).
					Return(nil, errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &RegisterService{
				UserRepo: mockRepo,
			}
			got, err := s.CekSessionByUUID(tt.args.ctx, tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterService.CekSessionByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ignoreFields := cmpopts.IgnoreFields(models.UserSession{}, "CreatedAt", "TokenExpired")
			if !cmp.Equal(got, tt.want, ignoreFields) {
				t.Errorf("RegisterService.CekSessionByUUID() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
