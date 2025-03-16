package repository

import (
	"backend-test/internal/models"
	"context"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestIngredientRepository_InsertIngredient(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

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
			name: "success",
			args: args{
				ctx: context.Background(),
				obj: &models.Ingredient{
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
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "ingredients" ("recipe_id","inventory_id","quantity") VALUES ($1,$2,$3) RETURNING "id"`,
				)).
					WithArgs(
						args.obj.RecipeID,
						args.obj.InventoryID,
						args.obj.Quantity,
					).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				obj: &models.Ingredient{
					RecipeID:    1,
					InventoryID: 1,
					Quantity:    2,
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "ingredients" ("recipe_id","inventory_id","quantity") VALUES ($1,$2,$3) RETURNING "id"`,
				)).
					WithArgs(
						args.obj.RecipeID,
						args.obj.InventoryID,
						args.obj.Quantity,
					).
					WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &IngredientRepository{
				DB: gormDB,
			}
			got, err := r.InsertIngredient(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientRepository.InsertIngredient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngredientRepository.InsertIngredient() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestIngredientRepository_UpdateIngredient(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

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
			name: "success",
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
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(
					`UPDATE "ingredients" SET "recipe_id"=$1,"inventory_id"=$2,"quantity"=$3 WHERE "id" = $4`,
				)).
					WithArgs(
						args.obj.RecipeID,
						args.obj.InventoryID,
						args.obj.Quantity,
						args.obj.ID,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				obj: &models.Ingredient{
					ID:          1,
					RecipeID:    1,
					InventoryID: 1,
					Quantity:    2,
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(
					`UPDATE "ingredients" SET "recipe_id"=$1,"inventory_id"=$2,"quantity"=$3 WHERE "id" = $4`,
				)).
					WithArgs(
						args.obj.RecipeID,
						args.obj.InventoryID,
						args.obj.Quantity,
						args.obj.ID,
					).
					WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			r := &IngredientRepository{
				DB: gormDB,
			}
			got, err := r.UpdateIngredient(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientRepository.UpdateIngredient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngredientRepository.UpdateIngredient() = %v, want %v", got, tt.want)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestIngredientRepository_DeleteIngredient(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)
	type args struct {
		ctx    context.Context
		argsID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				argsID: 1,
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "ingredients" WHERE id = $1`)).
					WithArgs(args.argsID).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				ctx:    context.Background(),
				argsID: 1,
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "ingredients" WHERE id = $1`)).
					WithArgs(args.argsID).
					WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &IngredientRepository{
				DB: gormDB,
			}
			if err := r.DeleteIngredient(tt.args.ctx, tt.args.argsID); (err != nil) != tt.wantErr {
				t.Errorf("IngredientRepository.DeleteIngredient() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestIngredientRepository_FindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)
	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Ingredient
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: &models.Ingredient{
				ID:          1,
				RecipeID:    1,
				InventoryID: 1,
				Quantity:    2,
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "ingredients" WHERE id = $1 ORDER BY "ingredients"."id" LIMIT $2`,
				)).WithArgs(args.ID, 1).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "recipe_id", "inventory_id", "quantity",
					}).AddRow(
						1, 1, 1, 2,
					))
			},
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "ingredients" WHERE id = $1 ORDER BY "ingredients"."id" LIMIT $2`,
				)).WithArgs(args.ID, 1).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &IngredientRepository{
				DB: gormDB,
			}
			got, err := r.FindByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngredientRepository.FindByID() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestIngredientRepository_GetAllIngredient(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	type args struct {
		ctx   context.Context
		param string
	}

	tests := []struct {
		name    string
		args    args
		want    *[]models.Ingredient
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx:   context.Background(),
				param: "",
			},
			want: &[]models.Ingredient{
				{ID: 1, RecipeID: 1, InventoryID: 1, Quantity: 2},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "ingredients"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "recipe_id", "inventory_id", "quantity"}).
						AddRow(1, 1, 1, 2))
			},
		},
		{
			name: "error",
			args: args{
				ctx:   context.Background(),
				param: "",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "ingredients"`)).
					WillReturnError(assert.AnError)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &IngredientRepository{
				DB: gormDB,
			}
			got, err := r.GetAllIngredient(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientRepository.GetAllIngredient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngredientRepository.GetAllIngredient() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestIngredientRepository_CustomIngerdient(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)
	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		args    args
		want    []models.IngredientCustom
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: []models.IngredientCustom{
				{ID: 1, InventoryID: 1, Quantity: 2, Item: "Sugar"},
			},
			wantErr: false,
			mockFn: func(args args) {
				query := regexp.QuoteMeta(`
					select 
					i.id,
					i.inventory_id,
					i.quantity,
					i2.item,
					i2.uom
					from ingredients i 
					left join inventories i2 on i.inventory_id = i2.id
					where i.recipe_id = ` + fmt.Sprint(args.ID))

				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"id", "inventory_id", "quantity", "item"}).
						AddRow(1, 1, 2, "Sugar"))
			},
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				query := regexp.QuoteMeta(`
					select 
					i.id,
					i.inventory_id,
					i.quantity,
					i2.item,
					i2.uom
					from ingredients i 
					left join inventories i2 on i.inventory_id = i2.id
					where i.recipe_id = ` + fmt.Sprint(args.ID))

				mock.ExpectQuery(query).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &IngredientRepository{
				DB: gormDB,
			}
			got, err := r.CustomIngerdient(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientRepository.CustomIngerdient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngredientRepository.CustomIngerdient() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestIngredientRepository_FindByIDRecipe(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)
	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		args    args
		want    models.RecipeFormat
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: models.RecipeFormat{
				ID:   1,
				Name: "Pasta Carbonara",
				SKU:  "PASTA-123",
				Cogs: 15000.0,
			},
			wantErr: false,
			mockFn: func(args args) {
				query := regexp.QuoteMeta(`
					select 
					r.id,
					r."name",
					r.sku,
					(SELECT COALESCE(SUM(i.quantity * i2.price_per_qty), 0) AS cogs
							FROM ingredients i
							LEFT JOIN inventories i2 ON i.inventory_id = i2.id
							WHERE i.recipe_id = ` + fmt.Sprint(args.ID) + `) as cogs
					from recipes r
					where r.id = ` + fmt.Sprint(args.ID))

				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "sku", "cogs"}).
						AddRow(1, "Pasta Carbonara", "PASTA-123", 15000.0))
			},
		},
		{
			name: "error - database failure",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want:    models.RecipeFormat{},
			wantErr: true,
			mockFn: func(args args) {
				query := regexp.QuoteMeta(`
					select 
					r.id,
					r."name",
					r.sku,
					(SELECT COALESCE(SUM(i.quantity * i2.price_per_qty), 0) AS cogs
							FROM ingredients i
							LEFT JOIN inventories i2 ON i.inventory_id = i2.id
							WHERE i.recipe_id = ` + fmt.Sprint(args.ID) + `) as cogs
					from recipes r
					where r.id = ` + fmt.Sprint(args.ID))

				mock.ExpectQuery(query).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &IngredientRepository{
				DB: gormDB,
			}
			got, err := r.FindByIDRecipe(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("IngredientRepository.FindByIDRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngredientRepository.FindByIDRecipe() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
