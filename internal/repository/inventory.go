package repository

import (
	"backend-test/internal/models"
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type InventoryRepository struct {
	DB *gorm.DB
}

func (r *InventoryRepository) InsertInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error) {
	if err := r.DB.Create(&obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *InventoryRepository) UpdateInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error) {
	if err := r.DB.Save(&obj).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

func (r *InventoryRepository) DeleteInv(ctx context.Context, ID int) error {
	tx := r.DB.Begin()
	// hapus ingeredients
	if err := tx.Where("inventory_id = ?", ID).Delete(&models.Ingredient{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// hapus inventory
	if err := tx.Where("id = ?", ID).Delete(&models.Inventory{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (r *InventoryRepository) FindByID(ctx context.Context, ID int) (*models.Inventory, error) {
	var (
		obj *models.Inventory
	)
	if err := r.DB.Where("id = ?", ID).First(&obj).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

func (r *InventoryRepository) GetAllInv(ctx context.Context, objComponent models.ComponentServerSide) (*[]models.Inventory, error) {
	var obj *[]models.Inventory
	limit := objComponent.Limit
	isOrder := objComponent.SortBy + ` ` + objComponent.SortType
	if objComponent.Search != "" {
		err := r.DB.Where(`lower(item) like '%` + strings.ToLower(objComponent.Search) + `%'`).Order(isOrder).Limit(limit).Offset(objComponent.Skip).Find(&obj).Error
		if err != nil {
			return nil, err
		}
		return obj, nil
	}
	err := r.DB.Order(isOrder).Limit(limit).Offset(objComponent.Skip).Find(&obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *InventoryRepository) CountData(ctx context.Context, objComponent models.ComponentServerSide) (int64, error) {
	var count int64

	if objComponent.Search != "" {
		err := r.DB.Debug().Where(`lower(item) like '%` + strings.ToLower(objComponent.Search) + `%'`).Model(&models.Inventory{}).Count(&count).Error
		if err != nil {
			return count, err
		}
		return count, nil
	}
	err := r.DB.Debug().Model(&models.Inventory{}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *InventoryRepository) BatchInsert(ctx context.Context, objs []models.Inventory) error {
	tx := r.DB.Begin()
	if err := tx.Create(&objs).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert inventories: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
