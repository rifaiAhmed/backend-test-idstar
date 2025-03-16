package services

import (
	"backend-test/external"
	"backend-test/helpers"
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RegisterService struct {
	UserRepo interfaces.IUserRepository
}

func (s *RegisterService) SubmitEmail(ctx context.Context, request models.User) (*models.User, string, error) {
	newUUID := uuid.New()
	var (
		now     = time.Now()
		objUser *models.User
		token   = ""
	)
	// find by email
	objUser, err := s.UserRepo.FindByEmail(ctx, request.Email)

	if err == gorm.ErrRecordNotFound {
		objUser, err = s.UserRepo.InsertNewUser(ctx, &request)
		if err != nil {
			return nil, token, err
		}
	}
	// generate token
	token, err = helpers.GenerateToken(ctx, objUser.ID, "token", objUser.Email, now)
	if err != nil {
		return nil, token, errors.Wrap(err, "failed to generate token")
	}
	// simpan token session
	userSession := &models.UserSession{
		UserID:       objUser.ID,
		Token:        token,
		Uuid:         newUUID,
		TokenExpired: now.Add(helpers.MapTypeToken["token"]),
	}
	err = s.UserRepo.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return nil, token, errors.Wrap(err, "failed to insert new session")
	}

	// send link
	var mail models.InternalNotificationRequest
	mail.Recipient = objUser.Email
	mail.UUID = userSession.Uuid
	err = s.SendEmail(ctx, mail)
	if err != nil {
		return nil, token, err
	}

	return objUser, token, nil
}

func (s *RegisterService) SendEmail(ctx context.Context, req models.InternalNotificationRequest) error {
	email := external.Email{
		To:      req.Recipient,
		Subject: "Submit Email Success",
		Body:    fmt.Sprintf(`Klik link berikut untuk mendapatkan token: <a href="%s">%s</a>`, "http://localhost:8080/auth/magic-link?uuid="+req.UUID.String(), "Klik di sini"),
	}
	err := email.SendEmail()
	if err != nil {
		return err
	}

	return nil
}

func (s *RegisterService) CekSessionByUUID(ctx context.Context, uuid uuid.UUID) (*models.UserSession, error) {
	data, err := s.UserRepo.CekSessionByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return data, nil
}
