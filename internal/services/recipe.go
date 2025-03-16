package services

import (
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"context"
)

type ReceipeService struct {
	RepoReceipe interfaces.IRecipeRepository
}

func (s *ReceipeService) InsertRecipe(ctx context.Context, obj *models.Recipe) (*models.Recipe, error) {
	data, err := s.RepoReceipe.InsertRecipe(ctx, obj)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *ReceipeService) UpdateRecipe(ctx context.Context, obj *models.Recipe) (*models.Recipe, error) {
	oldData, err := s.RepoReceipe.FindByID(ctx, int(obj.ID))
	if err != nil {
		return nil, err
	}

	// maaping data
	oldData.Name = obj.Name

	// update
	newData, err := s.RepoReceipe.UpdateRecipe(ctx, oldData)
	if err != nil {
		return nil, err
	}

	return newData, nil
}

func (s *ReceipeService) DeleteRecipe(ctx context.Context, ID int) error {
	return s.RepoReceipe.DeleteRecipe(ctx, ID)
}

func (s *ReceipeService) GetAllRecipe(ctx context.Context, objComponent models.ComponentServerSide, param string) ([]models.RecipeFormat, error) {
	objs, err := s.RepoReceipe.GetAllRecipe(ctx, objComponent, param)
	if err != nil {
		return nil, err
	}

	return objs, nil
}

func (s *ReceipeService) CountData(ctx context.Context, objComponent models.ComponentServerSide) (int64, error) {
	count, err := s.RepoReceipe.CountData(ctx, objComponent)

	if err != nil {
		return count, err
	}

	return count, nil
}
