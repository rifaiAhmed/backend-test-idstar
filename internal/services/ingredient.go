package services

import (
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"context"
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
