package services

import (
	"backend-test/internal/mocks"
	"backend-test/internal/models"
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"
)

func TestReceipeService_InsertRecipe(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIRecipeRepository(ctrlMock)

	type args struct {
		ctx context.Context
		obj *models.Recipe
	}

	tests := []struct {
		name    string
		args    args
		want    *models.Recipe
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - insert recipe",
			args: args{
				ctx: context.Background(),
				obj: &models.Recipe{},
			},
			want: &models.Recipe{
				ID:        1,
				Name:      "Nasi Goreng",
				SKU:       "SKU-250316-00001",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					InsertRecipe(args.ctx, args.obj).
					Return(&models.Recipe{
						ID:        1,
						Name:      args.obj.Name,
						SKU:       args.obj.SKU,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					}, nil)
			},
		},
		{
			name: "error - validation failed",
			args: args{
				ctx: context.Background(),
				obj: &models.Recipe{
					Name: "Nasi Goreng",
					SKU:  "SKU-250316-00001",
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					InsertRecipe(args.ctx, args.obj).
					Return(nil, errors.New("validation failed: name is required"))
			},
		},
		{
			name: "error - database failure",
			args: args{
				ctx: context.Background(),
				obj: &models.Recipe{
					Name: "Nasi Goreng",
					SKU:  "SKU-250316-00001",
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					InsertRecipe(args.ctx, args.obj).
					Return(nil, errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &ReceipeService{
				RepoReceipe: mockRepo,
			}
			got, err := s.InsertRecipe(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReceipeService.InsertRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Abaikan CreatedAt & UpdatedAt dalam perbandingan
			ignoreFields := cmpopts.IgnoreFields(models.Recipe{}, "CreatedAt", "UpdatedAt")
			if !cmp.Equal(got, tt.want, ignoreFields) {
				t.Errorf("ReceipeService.InsertRecipe() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestReceipeService_UpdateRecipe(t *testing.T) {
	now := time.Now()
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIRecipeRepository(ctrlMock)

	type args struct {
		ctx context.Context
		obj *models.Recipe
	}

	tests := []struct {
		name    string
		args    args
		want    *models.Recipe
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - update recipe",
			args: args{
				ctx: context.Background(),
				obj: &models.Recipe{
					ID:   1,
					Name: "Nasi Goreng",
					SKU:  "SKU-250316-00001",
				},
			},
			want: &models.Recipe{
				ID:        1,
				Name:      "Nasi Goreng",
				SKU:       "SKU-250316-00001",
				CreatedAt: now,
				UpdatedAt: now,
			},
			wantErr: false,
			mockFn: func(args args) {

				mockRepo.EXPECT().
					FindByID(args.ctx, int(args.obj.ID)).
					Return(&models.Recipe{
						ID:        1,
						Name:      "Nasi Goreng Original",
						SKU:       "SKU-250316-00001",
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)

				mockRepo.EXPECT().
					UpdateRecipe(args.ctx, gomock.Any()).
					Return(&models.Recipe{
						ID:          1,
						Name:        args.obj.Name,
						SKU:         args.obj.SKU,
						Ingredients: args.obj.Ingredients,
						CreatedAt:   now,
						UpdatedAt:   now,
					}, nil)
			},
		},
		{
			name: "error - recipe not found",
			args: args{
				ctx: context.Background(),
				obj: &models.Recipe{
					ID:   99,
					Name: "Non-existent Recipe",
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					FindByID(args.ctx, int(args.obj.ID)).
					Return(nil, errors.New("recipe not found"))
			},
		},
		{
			name: "error - database update failure",
			args: args{
				ctx: context.Background(),
				obj: &models.Recipe{
					ID:   1,
					Name: "Failed Update Recipe",
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				// Mock FindByID berhasil
				mockRepo.EXPECT().
					FindByID(args.ctx, int(args.obj.ID)).
					Return(&models.Recipe{
						ID:        1,
						Name:      "Nasi Goreng",
						SKU:       "SKU-250316-00001",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					}, nil)

				// Mock UpdateRecipe gagal
				mockRepo.EXPECT().
					UpdateRecipe(args.ctx, gomock.Any()).
					Return(nil, errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &ReceipeService{
				RepoReceipe: mockRepo,
			}
			got, err := s.UpdateRecipe(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReceipeService.UpdateRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ignoreFields := cmpopts.IgnoreFields(models.Recipe{}, "UpdatedAt")
			if !cmp.Equal(got, tt.want, ignoreFields) {
				t.Errorf("ReceipeService.UpdateRecipe() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestReceipeService_DeleteRecipe(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIRecipeRepository(ctrlMock)

	type args struct {
		ctx context.Context
		ID  int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - delete recipe",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					DeleteRecipe(args.ctx, args.ID).
					Return(nil)
			},
		},
		{
			name: "error - recipe not found",
			args: args{
				ctx: context.Background(),
				ID:  99,
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					DeleteRecipe(args.ctx, args.ID).
					Return(errors.New("recipe not found"))
			},
		},
		{
			name: "error - database error",
			args: args{
				ctx: context.Background(),
				ID:  2,
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					DeleteRecipe(args.ctx, args.ID).
					Return(errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &ReceipeService{
				RepoReceipe: mockRepo,
			}

			err := s.DeleteRecipe(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReceipeService.DeleteRecipe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReceipeService_GetAllRecipe(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIRecipeRepository(ctrlMock)

	type args struct {
		ctx          context.Context
		objComponent models.ComponentServerSide
		param        string
	}

	tests := []struct {
		name    string
		args    args
		want    []models.RecipeFormat
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - get all recipes",
			args: args{
				ctx:          context.Background(),
				objComponent: models.ComponentServerSide{Limit: 10, Offset: 0},
				param:        "chicken",
			},
			want: []models.RecipeFormat{
				{ID: 1, Name: "Chicken Curry", SKU: "SKU-001"},
				{ID: 2, Name: "Fried Chicken", SKU: "SKU-002"},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					GetAllRecipe(args.ctx, args.objComponent, args.param).
					Return([]models.RecipeFormat{
						{ID: 1, Name: "Chicken Curry", SKU: "SKU-001"},
						{ID: 2, Name: "Fried Chicken", SKU: "SKU-002"},
					}, nil)
			},
		},
		{
			name: "error - repository error",
			args: args{
				ctx:          context.Background(),
				objComponent: models.ComponentServerSide{Limit: 10, Offset: 0},
				param:        "beef",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					GetAllRecipe(args.ctx, args.objComponent, args.param).
					Return(nil, errors.New("database error"))
			},
		},
		{
			name: "success - empty recipes list",
			args: args{
				ctx:          context.Background(),
				objComponent: models.ComponentServerSide{Limit: 10, Offset: 0},
				param:        "vegetable",
			},
			want:    []models.RecipeFormat{},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					GetAllRecipe(args.ctx, args.objComponent, args.param).
					Return([]models.RecipeFormat{}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &ReceipeService{
				RepoReceipe: mockRepo,
			}

			got, err := s.GetAllRecipe(tt.args.ctx, tt.args.objComponent, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReceipeService.GetAllRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReceipeService.GetAllRecipe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReceipeService_CountData(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIRecipeRepository(ctrlMock)

	type args struct {
		ctx          context.Context
		objComponent models.ComponentServerSide
	}

	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - count data",
			args: args{
				ctx: context.Background(),
				objComponent: models.ComponentServerSide{
					Limit:  10,
					Offset: 0,
				},
			},
			want:    25,
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					CountData(args.ctx, args.objComponent).
					Return(int64(25), nil)
			},
		},
		{
			name: "error - repository error",
			args: args{
				ctx: context.Background(),
				objComponent: models.ComponentServerSide{
					Limit:  10,
					Offset: 0,
				},
			},
			want:    0,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					CountData(args.ctx, args.objComponent).
					Return(int64(0), errors.New("database error"))
			},
		},
		{
			name: "success - count is zero",
			args: args{
				ctx: context.Background(),
				objComponent: models.ComponentServerSide{
					Limit:  10,
					Offset: 0,
				},
			},
			want:    0,
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					CountData(args.ctx, args.objComponent).
					Return(int64(0), nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &ReceipeService{
				RepoReceipe: mockRepo,
			}

			got, err := s.CountData(tt.args.ctx, tt.args.objComponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReceipeService.CountData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReceipeService.CountData() = %v, want %v", got, tt.want)
			}
		})
	}
}
