package interfaces

import (
	"backend-test/internal/models"
	"context"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=IIngredient.go -destination=../mocks/IIngredient_mock.go -package=mocks
type IIngredientRepository interface {
	InsertIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error)
	UpdateIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error)
	DeleteIngredient(ctx context.Context, ID int) error
	FindByID(ctx context.Context, ID int) (*models.Ingredient, error)
	GetAllIngredient(ctx context.Context, param string) (*[]models.Ingredient, error)
	GetRecipeIncludeIngredients(ctx context.Context, ID int) (models.RecipeFormat, *[]models.IngredientCustom, error)
	MultipleCreateUpdate(ctx context.Context, objs []models.Ingredient) error
}

type IIngredientService interface {
	InsertIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error)
	UpdateIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error)
	DeleteIngredient(ctx context.Context, ID int) error
	GetAllIngredient(ctx context.Context, param string) (*[]models.Ingredient, error)
	GetRecipeIncludeIngredients(ctx context.Context, ID int) (models.RecipeFormat, *[]models.IngredientCustom, error)
	MultipleCreateUpdate(ctx context.Context, data models.MultipleIngredients) error
}

type IIngredientHandler interface {
	InsertIngredient(c *gin.Context)
	UpdateIngredient(c *gin.Context)
	DeleteIngredient(c *gin.Context)
	GetAllIngredient(c *gin.Context)
	GetRecipeIncludeIngredients(c *gin.Context)
	MultipleCreateUpdate(c *gin.Context)
}
