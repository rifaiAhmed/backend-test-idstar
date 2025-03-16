package repository

import (
	"backend-test/internal/models"
	"context"
	"fmt"
	"sync"

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

func (r *IngredientRepository) CustomIngerdient(ctx context.Context, ID int) ([]models.IngredientCustom, error) {
	var (
		obj []models.IngredientCustom
	)
	query := `
	select 
	i.id,
	i.inventory_id,
	i.quantity,
	i2.item,
	i2.uom
	from ingredients i 
	left join inventories i2 on i.inventory_id = i2.id
	where i.recipe_id = ` + fmt.Sprint(ID)
	if err := r.DB.Raw(query).Scan(&obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *IngredientRepository) FindByIDRecipe(ctx context.Context, ID int) (models.RecipeFormat, error) {
	var (
		obj models.RecipeFormat
	)
	query := `
	select 
	r.id,
	r."name",
	r.sku,
	(SELECT COALESCE(SUM(i.quantity * i2.price_per_qty), 0) AS cogs
			FROM ingredients i
			LEFT JOIN inventories i2 ON i.inventory_id = i2.id
			WHERE i.recipe_id = ` + fmt.Sprint(ID) + `) as cogs
	from recipes r
	where r.id = ` + fmt.Sprint(ID) + `
	`
	if err := r.DB.Raw(query).Scan(&obj).Error; err != nil {
		return obj, err
	}

	return obj, nil
}

func (r *IngredientRepository) GetRecipeIncludeIngredients(ctx context.Context, ID int) (models.RecipeFormat, *[]models.IngredientCustom, error) {
	var (
		obj  models.RecipeFormat
		objs *[]models.IngredientCustom
		errs []error
	)

	objChan := make(chan models.RecipeFormat, 1)
	objsChan := make(chan *[]models.IngredientCustom, 1)
	errChan := make(chan error, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		obj, err := r.FindByIDRecipe(ctx, ID)
		if err != nil {
			errChan <- err
			return
		}
		objChan <- obj
	}()

	go func() {
		defer wg.Done()
		all, err := r.CustomIngerdient(ctx, ID)
		if err != nil {
			errChan <- err
			return
		}
		objsChan <- &all
	}()

	wg.Wait()
	close(objChan)
	close(objsChan)
	close(errChan)

	if result, ok := <-objChan; ok {
		obj = result
	}
	if result, ok := <-objsChan; ok {
		objs = result
	}

	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return obj, objs, fmt.Errorf("error occurred: %v", errs)
	}

	return obj, objs, nil
}

func (r *IngredientRepository) MultipleCreateUpdate(ctx context.Context, objs []models.Ingredient) error {
	tx := r.DB.Begin()
	for _, obj := range objs {
		if err := tx.Save(&obj).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
