package services

import (
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"context"

	"gorm.io/gorm"
)

type InventoryService struct {
	InventoryRepo interfaces.IInventoryRepository
}

func (s *InventoryService) InsertInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error) {
	data, err := s.InventoryRepo.InsertInv(ctx, obj)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *InventoryService) UpdateInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error) {
	// find by ID
	old, err := s.InventoryRepo.FindByID(ctx, int(obj.ID))
	if err != nil || err == gorm.ErrRecordNotFound {
		return nil, err
	}
	// perbarui data
	old.Item = obj.Item
	old.PricePerQty = obj.PricePerQty
	old.Qty = obj.Qty
	old.Uom = obj.Uom

	// update
	newData, err := s.InventoryRepo.UpdateInv(ctx, old)
	if err != nil {
		return nil, err
	}

	return newData, nil
}

func (s *InventoryService) DeleteInv(ctx context.Context, ID int) error {
	return s.InventoryRepo.DeleteInv(ctx, ID)
}

func (s *InventoryService) GetAllInv(ctx context.Context, objComponent models.ComponentServerSide) (*[]models.Inventory, error) {
	data, err := s.InventoryRepo.GetAllInv(ctx, objComponent)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *InventoryService) CountData(ctx context.Context, objComponent models.ComponentServerSide) (int64, error) {
	count, err := s.InventoryRepo.CountData(ctx, objComponent)

	if err != nil {
		return count, err
	}

	return count, nil
}
