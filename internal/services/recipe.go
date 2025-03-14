package services

import (
	"backend-test/helpers"
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

func (s *ReceipeService) DeleteRecipe(ctx context.Context, ID int) error {
	return s.RepoReceipe.DeleteRecipe(ctx, ID)
}

func (s *ReceipeService) GetAllRecipe(ctx context.Context, param string) ([]helpers.Recipe, error) {
	objs, err := s.RepoReceipe.GetAllRecipe(ctx, param)
	if err != nil {
		return nil, err
	}

	return objs, nil
}
