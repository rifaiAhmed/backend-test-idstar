package services

import (
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"context"
	"fmt"
	"strconv"
	"strings"
)

type IngredientService struct {
	RepoIngredient interfaces.IIngredientRepository
}

func (s *IngredientService) InsertIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error) {
	data, err := s.RepoIngredient.InsertIngredient(ctx, obj)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *IngredientService) UpdateIngredient(ctx context.Context, obj *models.Ingredient) (*models.Ingredient, error) {
	// get old data
	oldData, err := s.RepoIngredient.FindByID(ctx, int(obj.ID))
	if err != nil {
		return nil, err
	}

	// perbaharui data
	oldData.InventoryID = obj.InventoryID
	oldData.RecipeID = obj.RecipeID
	oldData.Quantity = obj.Quantity

	// update
	newData, err := s.RepoIngredient.UpdateIngredient(ctx, oldData)
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (s *IngredientService) DeleteIngredient(ctx context.Context, ID int) error {
	return s.RepoIngredient.DeleteIngredient(ctx, ID)
}

func (s *IngredientService) GetAllIngredient(ctx context.Context, param string) (*[]models.Ingredient, error) {
	objs, err := s.RepoIngredient.GetAllIngredient(ctx, param)
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (s *IngredientService) GetRecipeIncludeIngredients(ctx context.Context, ID int) (models.RecipeFormat, *[]models.IngredientCustom, error) {
	recipe, ingredients, err := s.RepoIngredient.GetRecipeIncludeIngredients(ctx, ID)
	if err != nil {
		return recipe, ingredients, err
	}

	return recipe, ingredients, nil
}

func (s *IngredientService) MultipleCreateUpdate(ctx context.Context, data models.MultipleIngredients) error {
	// maaping data
	objs, err := ParseIngredients(data.Data, data.RecipeID)
	if err != nil {
		return err
	}

	return s.RepoIngredient.MultipleCreateUpdate(ctx, objs)
}

func ParseIngredients(data string, RecipeID int) ([]models.Ingredient, error) {
	var ingredients []models.Ingredient

	rows := strings.Split(data, ",")

	for _, row := range rows {
		fields := strings.Split(row, "|")
		if len(fields) != 3 {
			return nil, fmt.Errorf("format data tidak valid: %s", row)
		}

		id, err := strconv.ParseUint(fields[0], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("ID tidak valid: %s", fields[0])
		}

		inventoryID, err := strconv.ParseUint(fields[1], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("InventoryID tidak valid: %s", fields[1])
		}

		quantity, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return nil, fmt.Errorf("quantity tidak valid: %s", fields[2])
		}

		ingredients = append(ingredients, models.Ingredient{
			ID:          uint(id),
			RecipeID:    uint(RecipeID),
			InventoryID: uint(inventoryID),
			Quantity:    quantity,
		})
	}

	return ingredients, nil
}
