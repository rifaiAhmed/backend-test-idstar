package repository

import (
	"backend-test/internal/models"
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRepository_InsertNewUser(t *testing.T) {
	now := time.Now()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				user: &models.User{
					Email: "test@gmail.com",
				},
			},
			want: &models.User{
				ID:    1,
				Email: "test@gmail.com",
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "users" ("email","created_at","updated_at") VALUES ($1,$2,$3) RETURNING "id"`,
				)).
					WithArgs(
						args.user.Email,
						sqlmock.AnyArg(), // CreatedAt
						sqlmock.AnyArg(), // UpdatedAt
					).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &UserRepository{
				DB: gormDB,
			}
			got, err := r.InsertNewUser(tt.args.ctx, tt.args.user)

			// Cek error sesuai ekspektasi
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.InsertNewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Pastikan data selain timestamp sesuai
			assert.Equal(t, tt.want.ID, got.ID)
			assert.Equal(t, tt.want.Email, got.Email)

			// Validasi CreatedAt & UpdatedAt dengan toleransi waktu 1 detik
			assert.WithinDuration(t, now, got.CreatedAt, time.Second)
			assert.WithinDuration(t, now, got.UpdatedAt, time.Second)

			// Pastikan ekspektasi SQL mock terpenuhi
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
