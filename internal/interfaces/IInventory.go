package interfaces

import (
	"backend-test/internal/models"
	"context"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=IInventory.go -destination=../mocks/IInventory_mock.go -package=mocks
type IInventoryRepository interface {
	InsertInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error)
	UpdateInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error)
	DeleteInv(ctx context.Context, ID int) error
	FindByID(ctx context.Context, ID int) (*models.Inventory, error)
	GetAllInv(ctx context.Context, objComponent models.ComponentServerSide) (*[]models.Inventory, error)
	CountData(ctx context.Context, objComponent models.ComponentServerSide) (int64, error)
	BatchInsert(ctx context.Context, objs []models.Inventory) error
}

type IInventoryService interface {
	InsertInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error)
	UpdateInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error)
	DeleteInv(ctx context.Context, ID int) error
	GetAllInv(ctx context.Context, objComponent models.ComponentServerSide) (*[]models.Inventory, error)
	CountData(ctx context.Context, objComponent models.ComponentServerSide) (int64, error)
	InsertFromExcel(ctx context.Context, filePath string) error
}

type IInventoryHandler interface {
	InsertInv(c *gin.Context)
	UpdateInv(c *gin.Context)
	DeleteInv(c *gin.Context)
	GetAllInv(c *gin.Context)
	GetTemplate(c *gin.Context)
	UploadExcel(c *gin.Context)
}
