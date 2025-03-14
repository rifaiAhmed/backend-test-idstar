package repository

import (
	"backend-test/helpers"
	"backend-test/internal/models"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type RecipeRepository struct {
	DB *gorm.DB
}

func (r *RecipeRepository) InsertRecipe(ctx context.Context, obj *models.Recipe) (*models.Recipe, error) {
	tx := r.DB.Begin()

	sku, err := GenerateSKU(tx)
	if err != nil {
		return nil, err
	}
	obj.SKU = sku
	if err := tx.Create(&obj).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return obj, nil
}

func (r *RecipeRepository) UpdateRecipe(ctx context.Context, obj *models.Recipe) (*models.Recipe, error) {
	if err := r.DB.Save(&obj).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

func (r *RecipeRepository) DeleteRecipe(ctx context.Context, ID int) error {
	tx := r.DB.Begin()
	// delete ingredient
	if err := tx.Where("recipe_id = ?", ID).Delete(&models.Ingredient{}).Error; err != nil {
		return err
	}

	// hapus receipe
	if err := tx.Where("id = ?", ID).Delete(&models.Recipe{}).Error; err != nil {
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

func (r *RecipeRepository) FindByID(ctx context.Context, ID int) (*models.Recipe, error) {
	var (
		obj *models.Recipe
	)
	if err := r.DB.Where("id = ?", ID).First(&obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *RecipeRepository) GetAllRecipe(ctx context.Context, param string) ([]helpers.Recipe, error) {
	var (
		objs            []models.Recipe
		objsIncludeCogs []helpers.Recipe
	)
	if err := r.DB.Preload("Ingredients").Find(&objs).Error; err != nil {
		return nil, err
	}
	for _, obj := range objs {
		var data helpers.Recipe
		data.ID = obj.ID
		data.Name = obj.Name
		data.SKU = obj.SKU
		cogs, err := r.GetCOGS(int(obj.ID))
		if err != nil {
			data.Cogs = 0
		}
		data.Cogs = *cogs
		objsIncludeCogs = append(objsIncludeCogs, data)
	}

	return objsIncludeCogs, nil
}

func GenerateSKU(tx *gorm.DB) (string, error) {
	var count int64
	if err := tx.Model(&models.Recipe{}).Count(&count).Error; err != nil {
		return "", err
	}
	currentDate := time.Now().Format("060102")

	transactionNumber := count + 1

	code := fmt.Sprintf("SKU-%s-%05d", currentDate, transactionNumber)
	return code, nil
}

func (r *RecipeRepository) GetCOGS(recipeID int) (*float64, error) {
	var cogs float64
	result := r.DB.Raw(`
		SELECT COALESCE(SUM(i.quantity * i2.price_per_qty), 0) AS cogs
		FROM ingredients i
		LEFT JOIN inventories i2 ON i.inventory_id = i2.id
		WHERE i.recipe_id = ?
	`, recipeID).Scan(&cogs)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch COGS: %w", result.Error)
	}

	return &cogs, nil
}
