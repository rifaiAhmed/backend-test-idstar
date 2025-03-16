package repository

import (
	"backend-test/internal/models"
	"context"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRecipeRepository_InsertRecipe(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	// now := time.Now()
	ctx := context.Background()

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
			name: "success",
			args: args{
				ctx: ctx,
				obj: &models.Recipe{
					Name: "Nasi Goreng",
				},
			},
			want: &models.Recipe{
				ID:   1,
				Name: "Nasi Goreng",
				SKU:  "SKU-250316-00001",
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT count(*) FROM "recipes"`)).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "recipes" ("name","sku","created_at","updated_at") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
					WithArgs(args.obj.Name, "SKU-250316-00001", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error - failed to generate SKU",
			args: args{
				ctx: ctx,
				obj: &models.Recipe{Name: "Nasi Goreng"},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT count(*) FROM "recipes"`)).
					WillReturnError(assert.AnError)

			},
		},
		{
			name: "error - failed to insert recipe",
			args: args{
				ctx: ctx,
				obj: &models.Recipe{Name: "Nasi Goreng"},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT count(*) FROM "recipes"`)).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "recipes" ("name","sku","created_at","updated_at") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
					WithArgs(args.obj.Name, "SKU-250316-00001", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			r := &RecipeRepository{
				DB: gormDB,
			}

			got, err := r.InsertRecipe(tt.args.ctx, tt.args.obj)

			if (err != nil) != tt.wantErr {
				t.Errorf("RecipeRepository.InsertRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, tt.want.ID, got.ID)
				assert.Equal(t, tt.want.Name, got.Name)
				assert.Equal(t, tt.want.SKU, got.SKU)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRecipeRepository_UpdateRecipe(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &RecipeRepository{
				DB: gormDB,
			}
			got, err := r.UpdateRecipe(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecipeRepository.UpdateRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecipeRepository.UpdateRecipe() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
