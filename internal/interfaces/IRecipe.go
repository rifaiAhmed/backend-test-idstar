package interfaces

import (
	"backend-test/helpers"
	"backend-test/internal/models"
	"context"

	"github.com/gin-gonic/gin"
)

type IRecipeRepository interface {
	InsertRecipe(ctx context.Context, obj *models.Recipe) (*models.Recipe, error)
	DeleteRecipe(ctx context.Context, ID int) error
	FindByID(ctx context.Context, ID int) (*models.Recipe, error)
	GetAllRecipe(ctx context.Context, param string) ([]helpers.Recipe, error)
}

type IRecipeService interface {
	InsertRecipe(ctx context.Context, obj *models.Recipe) (*models.Recipe, error)
	DeleteRecipe(ctx context.Context, ID int) error
	GetAllRecipe(ctx context.Context, param string) ([]helpers.Recipe, error)
}

type IRecipeHandler interface {
	InsertRecipe(c *gin.Context)
	DeleteRecipe(c *gin.Context)
	GetAllRecipe(c *gin.Context)
}
