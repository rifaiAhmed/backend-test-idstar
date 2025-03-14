package interfaces

import (
	"backend-test/internal/models"
	"context"

	"github.com/gin-gonic/gin"
)

type IIngredientRepository interface {
	InsertIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error)
	UpdateIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error)
	DeleteIngredient(ctx context.Context, ID int) error
	FindByID(ctx context.Context, ID int) (*models.Ingredient, error)
	GetAllIngredient(ctx context.Context, param string) (*[]models.Ingredient, error)
}

type IIngredientService interface {
	InsertIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error)
	UpdateIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error)
	DeleteIngredient(ctx context.Context, ID int) error
	GetAllIngredient(ctx context.Context, param string) (*[]models.Ingredient, error)
}

type IIngredientHandler interface {
	InsertIngredient(c *gin.Context)
	UpdateIngredient(c *gin.Context)
	DeleteIngredient(c *gin.Context)
	GetAllIngredient(c *gin.Context)
}
