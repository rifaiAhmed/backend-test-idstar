package repository

import (
	"backend-test/internal/models"
	"context"

	"gorm.io/gorm"
)

type IngredientRepository struct {
	DB *gorm.DB
}

func (r *IngredientRepository) InsertIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error) {
	if err := r.DB.Create(&obj).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

func (r *IngredientRepository) UpdateIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error) {
	if err := r.DB.Save(&obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *IngredientRepository) DeleteIngredient(ctx context.Context, ID int) error {
	return r.DB.Where("id = ?", ID).Delete(&models.Ingredient{}).Error
}

func (r *IngredientRepository) FindByID(ctx context.Context, ID int) (*models.Ingredient, error) {
	var (
		obj *models.Ingredient
	)
	if err := r.DB.Where("id = ?", ID).First(&obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *IngredientRepository) GetAllIngredient(ctx context.Context, param string) (*[]models.Ingredient, error) {
	var (
		objs *[]models.Ingredient
	)
	if err := r.DB.Find(&objs).Error; err != nil {
		return nil, err
	}

	return objs, nil
}
