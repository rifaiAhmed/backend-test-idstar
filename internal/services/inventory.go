package services

import (
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type InventoryService struct {
	InventoryRepo interfaces.IInventoryRepository
}

func (s *InventoryService) InsertInv(ctx context.Context, obj *models.Inventory) (*models.Inventory, error) {
	data, err := s.InventoryRepo.InsertInv(ctx, obj)
	if err != nil {
		return nil, err
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

func (s *InventoryService) InsertFromExcel(ctx context.Context, filePath string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return fmt.Errorf("failed to read sheet: %v", err)
	}

	var inventories []models.Inventory

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) < 4 {
			logrus.Printf("Warning: Row %d has missing columns, skipping...", i+1)
			continue
		}

		qty, err := strconv.ParseFloat(strings.TrimSpace(row[1]), 64)
		if err != nil {
			logrus.Printf("Warning: Invalid quantity in row %d: %v", i+1, err)
			continue
		}

		price, err := strconv.ParseFloat(strings.TrimSpace(row[3]), 64)
		if err != nil {
			logrus.Printf("Warning: Invalid price in row %d: %v", i+1, err)
			continue
		}

		inventories = append(inventories, models.Inventory{
			Item:        row[0],
			Qty:         qty,
			Uom:         row[2],
			PricePerQty: price,
		})
	}
	if len(inventories) == 0 {
		return fmt.Errorf("no valid inventories found in the file")
	}
	err = s.InventoryRepo.BatchInsert(ctx, inventories)
	if err != nil {
		return err
	}
	return nil
}
