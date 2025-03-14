package interfaces

import (
	"backend-test/internal/models"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
	GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error)
	CekSessionByUUID(ctx context.Context, uuid uuid.UUID) (*models.UserSession, error)
}

type IUserService interface {
	SubmitEmail(ctx context.Context, request models.User) (*models.User, string, error)
	CekSessionByUUID(ctx context.Context, uuid uuid.UUID) (*models.UserSession, error)
}

type IUserHandler interface {
	SendMail(c *gin.Context)
	CekSessionByUUID(c *gin.Context)
}
