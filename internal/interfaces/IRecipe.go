package interfaces

import (
	"backend-test/internal/models"
	"context"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=IRecipe.go -destination=../mocks/IRecipe_mock.go -package=mocks
type IRecipeRepository interface {
	InsertRecipe(ctx context.Context, obj *models.Recipe) (*models.Recipe, error)
	UpdateRecipe(ctx context.Context, obj *models.Recipe) (*models.Recipe, error)
	DeleteRecipe(ctx context.Context, ID int) error
	FindByID(ctx context.Context, ID int) (*models.Recipe, error)
	GetAllRecipe(ctx context.Context, objComponent models.ComponentServerSide, param string) ([]models.RecipeFormat, error)
	CountData(ctx context.Context, objComponent models.ComponentServerSide) (int64, error)
}

type IRecipeService interface {
	InsertRecipe(ctx context.Context, obj *models.Recipe) (*models.Recipe, error)
	UpdateRecipe(ctx context.Context, obj *models.Recipe) (*models.Recipe, error)
	DeleteRecipe(ctx context.Context, ID int) error
	GetAllRecipe(ctx context.Context, objComponent models.ComponentServerSide, param string) ([]models.RecipeFormat, error)
	CountData(ctx context.Context, objComponent models.ComponentServerSide) (int64, error)
}

type IRecipeHandler interface {
	InsertRecipe(c *gin.Context)
	UpdateRecipe(c *gin.Context)
	DeleteRecipe(c *gin.Context)
	GetAllRecipe(c *gin.Context)
}
