package services

import (
	"backend-test/internal/mocks"
	"backend-test/internal/models"
	"context"
	"errors"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestIngredientService_InsertIngredient(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIIngredientRepository(ctrlMock)

	type args struct {
		ctx context.Context
		obj *models.Ingredient
	}

	tests := []struct {
		name    string
		args    args
		want    *models.Ingredient
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - insert ingredient",
			args: args{
				ctx: context.Background(),
				obj: &models.Ingredient{
					ID:          1,
					RecipeID:    1,
					InventoryID: 1,
					Quantity:    2,
				},
			},
			want: &models.Ingredient{
				ID:          1,
				RecipeID:    1,
				InventoryID: 1,
				Quantity:    2,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					InsertIngredient(args.ctx, args.obj).
					Return(args.obj, nil).
					Times(1)
			},
		},
		{
			name: "fail - repository returns error",
			args: args{
				ctx: context.Background(),
				obj: &models.Ingredient{
					ID:          2,
					RecipeID:    1,
					InventoryID: 1,
					Quantity:    2,
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					InsertIngredient(args.ctx, args.obj).
					Return(nil, errors.New("database error")).
					Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &IngredientService{
				RepoIngredient: mockRepo,
			}

			got, err := s.InsertIngredient(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientService.InsertIngredient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngredientService.InsertIngredient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngredientService_UpdateIngredient(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIIngredientRepository(ctrlMock)

	type args struct {
		ctx context.Context
		obj *models.Ingredient
	}

	tests := []struct {
		name    string
		args    args
		want    *models.Ingredient
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - update ingredient",
			args: args{
				ctx: context.Background(),
				obj: &models.Ingredient{
					ID:          1,
					InventoryID: 100,
					RecipeID:    200,
					Quantity:    50,
				},
			},
			want: &models.Ingredient{
				ID:          1,
				InventoryID: 100,
				RecipeID:    200,
				Quantity:    50,
			},
			wantErr: false,
			mockFn: func(args args) {

				mockRepo.EXPECT().
					FindByID(args.ctx, int(args.obj.ID)).
					Return(&models.Ingredient{ID: 1, InventoryID: 101, RecipeID: 201, Quantity: 30}, nil).
					Times(1)

				mockRepo.EXPECT().
					UpdateIngredient(args.ctx, &models.Ingredient{ID: 1, InventoryID: 100, RecipeID: 200, Quantity: 50}).
					Return(args.obj, nil).
					Times(1)
			},
		},
		{
			name: "fail - ingredient not found",
			args: args{
				ctx: context.Background(),
				obj: &models.Ingredient{
					ID:          2,
					InventoryID: 300,
					RecipeID:    400,
					Quantity:    10,
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {

				mockRepo.EXPECT().
					FindByID(args.ctx, int(args.obj.ID)).
					Return(nil, errors.New("ingredient not found")).
					Times(1)
			},
		},
		{
			name: "fail - repository update error",
			args: args{
				ctx: context.Background(),
				obj: &models.Ingredient{
					ID:          3,
					InventoryID: 500,
					RecipeID:    600,
					Quantity:    20,
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {

				mockRepo.EXPECT().
					FindByID(args.ctx, int(args.obj.ID)).
					Return(&models.Ingredient{ID: 3, InventoryID: 501, RecipeID: 601, Quantity: 15}, nil).
					Times(1)

				mockRepo.EXPECT().
					UpdateIngredient(args.ctx, &models.Ingredient{ID: 3, InventoryID: 500, RecipeID: 600, Quantity: 20}).
					Return(nil, errors.New("update failed")).
					Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &IngredientService{
				RepoIngredient: mockRepo,
			}

			got, err := s.UpdateIngredient(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientService.UpdateIngredient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngredientService.UpdateIngredient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngredientService_DeleteIngredient(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIIngredientRepository(ctrlMock)

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
			name: "success - delete ingredient",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			wantErr: false,
			mockFn: func(args args) {

				mockRepo.EXPECT().
					DeleteIngredient(args.ctx, args.ID).
					Return(nil).
					Times(1)
			},
		},
		{
			name: "fail - ingredient not found",
			args: args{
				ctx: context.Background(),
				ID:  2,
			},
			wantErr: true,
			mockFn: func(args args) {

				mockRepo.EXPECT().
					DeleteIngredient(args.ctx, args.ID).
					Return(errors.New("ingredient not found")).
					Times(1)
			},
		},
		{
			name: "fail - database error",
			args: args{
				ctx: context.Background(),
				ID:  3,
			},
			wantErr: true,
			mockFn: func(args args) {

				mockRepo.EXPECT().
					DeleteIngredient(args.ctx, args.ID).
					Return(errors.New("database error")).
					Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &IngredientService{
				RepoIngredient: mockRepo,
			}

			err := s.DeleteIngredient(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientService.DeleteIngredient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIngredientService_GetAllIngredient(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIIngredientRepository(ctrlMock)

	type args struct {
		ctx   context.Context
		param string
	}

	mockIngredients := []models.Ingredient{
		{ID: 1, InventoryID: 101, RecipeID: 201, Quantity: 2.5},
		{ID: 2, InventoryID: 102, RecipeID: 202, Quantity: 3.0},
	}

	tests := []struct {
		name    string
		args    args
		want    *[]models.Ingredient
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - get all ingredients",
			args: args{
				ctx:   context.Background(),
				param: "recipe_201",
			},
			want:    &mockIngredients,
			wantErr: false,
			mockFn: func(args args) {

				mockRepo.EXPECT().
					GetAllIngredient(args.ctx, args.param).
					Return(&mockIngredients, nil).
					Times(1)
			},
		},
		{
			name: "fail - no ingredients found",
			args: args{
				ctx:   context.Background(),
				param: "recipe_999",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {

				mockRepo.EXPECT().
					GetAllIngredient(args.ctx, args.param).
					Return(nil, errors.New("no ingredients found")).
					Times(1)
			},
		},
		{
			name: "fail - database error",
			args: args{
				ctx:   context.Background(),
				param: "recipe_500",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {

				mockRepo.EXPECT().
					GetAllIngredient(args.ctx, args.param).
					Return(nil, errors.New("database error")).
					Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &IngredientService{
				RepoIngredient: mockRepo,
			}

			got, err := s.GetAllIngredient(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientService.GetAllIngredient() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngredientService.GetAllIngredient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngredientService_GetRecipeIncludeIngredients(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIIngredientRepository(ctrlMock)
	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		args    args
		want    models.RecipeFormat
		want1   *[]models.IngredientCustom
		wantErr bool
		mockFn  func(args args)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &IngredientService{
				RepoIngredient: mockRepo,
			}
			got, got1, err := s.GetRecipeIncludeIngredients(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientService.GetRecipeIncludeIngredients() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngredientService.GetRecipeIncludeIngredients() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("IngredientService.GetRecipeIncludeIngredients() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIngredientService_MultipleCreateUpdate(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIIngredientRepository(ctrlMock)

	type args struct {
		ctx  context.Context
		data models.MultipleIngredients
	}

	validData := "1|101|2.5,2|102|3.0"
	invalidData := "1|101,2|102|xyz"

	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - multiple ingredients created/updated",
			args: args{
				ctx: context.Background(),
				data: models.MultipleIngredients{
					Data:     validData,
					RecipeID: 10,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					MultipleCreateUpdate(args.ctx, gomock.Any()).
					Return(nil).
					Times(1)
			},
		},
		{
			name: "fail - invalid data format",
			args: args{
				ctx: context.Background(),
				data: models.MultipleIngredients{
					Data:     invalidData,
					RecipeID: 10,
				},
			},
			wantErr: true,
			mockFn: func(args args) {

			},
		},
		{
			name: "fail - repository error",
			args: args{
				ctx: context.Background(),
				data: models.MultipleIngredients{
					Data:     validData,
					RecipeID: 10,
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					MultipleCreateUpdate(args.ctx, gomock.Any()).
					Return(errors.New("database error")).
					Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &IngredientService{
				RepoIngredient: mockRepo,
			}

			err := s.MultipleCreateUpdate(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientService.MultipleCreateUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
