package repository

import (
	"backend-test/internal/models"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) (*models.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var (
		obj *models.User
	)
	if err := r.DB.Where("email = ?", email).First(&obj).Error; err != nil {
		return obj, err
	}

	return obj, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, session *models.UserSession) error {
	return r.DB.Create(&session).Error
}

func (r *UserRepository) CekSessionByUUID(ctx context.Context, uuid uuid.UUID) (*models.UserSession, error) {
	var obj *models.UserSession
	if err := r.DB.Where("uuid = ?", uuid).First(&obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error) {
	var (
		session models.UserSession
		err     error
	)
	err = r.DB.Where("token = ?", token).First(&session).Error
	if err != nil {
		return session, err
	}
	if session.ID == 0 {
		return session, errors.New("session not found")
	}
	return session, nil
}
