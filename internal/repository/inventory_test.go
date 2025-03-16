package repository

import (
	"backend-test/internal/models"
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestInventoryRepository_InsertInv(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	now := time.Now()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)
	type args struct {
		ctx context.Context
		obj *models.Inventory
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Inventory
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				obj: &models.Inventory{
					Item:        "Plastik",
					Qty:         1,
					Uom:         "Kg",
					PricePerQty: 10000,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "inventories" ("item","qty","uom","price_per_qty","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`,
				)).
					WithArgs(
						args.obj.Item,
						args.obj.Qty,
						args.obj.Uom,
						args.obj.PricePerQty,
						args.obj.CreatedAt,
						args.obj.UpdatedAt,
					).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
			want: &models.Inventory{
				ID:          1,
				Item:        "Plastik",
				Qty:         1,
				Uom:         "Kg",
				PricePerQty: 10000,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				obj: &models.Inventory{
					Item:        "Plastik",
					Qty:         1,
					Uom:         "Kg",
					PricePerQty: 10000,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "inventories" ("item","qty","uom","price_per_qty","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`,
				)).
					WithArgs(
						args.obj.Item,
						args.obj.Qty,
						args.obj.Uom,
						args.obj.PricePerQty,
						args.obj.CreatedAt,
						args.obj.UpdatedAt,
					).
					WillReturnError(errors.New("insert failed"))

				mock.ExpectRollback()
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &InventoryRepository{
				DB: gormDB,
			}
			got, err := r.InsertInv(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryRepository.InsertInv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InventoryRepository.InsertInv() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestInventoryRepository_UpdateInv(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	now := time.Now()

	type args struct {
		ctx context.Context
		obj *models.Inventory
	}

	tests := []struct {
		name    string
		args    args
		want    *models.Inventory
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				obj: &models.Inventory{
					ID:          1,
					Item:        "Plastik",
					Qty:         1,
					Uom:         "Kg",
					PricePerQty: 10000,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			want: &models.Inventory{
				ID:          1,
				Item:        "Plastik",
				Qty:         1,
				Uom:         "Kg",
				PricePerQty: 10000,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(
					`UPDATE "inventories" SET "item"=$1,"qty"=$2,"uom"=$3,"price_per_qty"=$4,"created_at"=$5,"updated_at"=$6 WHERE "id" = $7`,
				)).
					WithArgs(
						args.obj.Item,
						args.obj.Qty,
						args.obj.Uom,
						args.obj.PricePerQty,
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.obj.ID,
					).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				obj: &models.Inventory{
					ID:          1,
					Item:        "Plastik",
					Qty:         1,
					Uom:         "Kg",
					PricePerQty: 10000,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(
					`UPDATE "inventories" SET "item"=$1,"qty"=$2,"uom"=$3,"price_per_qty"=$4,"created_at"=$5,"updated_at"=$6 WHERE "id" = $7`,
				)).
					WithArgs(
						args.obj.Item,
						args.obj.Qty,
						args.obj.Uom,
						args.obj.PricePerQty,
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.obj.ID,
					).
					WillReturnError(errors.New("failed to execute update query"))
				mock.ExpectRollback()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			r := &InventoryRepository{
				DB: gormDB,
			}

			got, err := r.UpdateInv(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryRepository.UpdateInv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil {
				assert.WithinDuration(t, tt.want.UpdatedAt, got.UpdatedAt, time.Millisecond)
				got.UpdatedAt = tt.want.UpdatedAt
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InventoryRepository.UpdateInv() = %v, want %v", got, tt.want)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestInventoryRepository_DeleteInv(t *testing.T) {
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
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "ingredients" WHERE inventory_id = $1`)).
					WithArgs(args.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "inventories" WHERE id = $1`)).
					WithArgs(args.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "error_delete_inventory",
			args: args{
				ctx: context.Background(),
				ID:  2,
			},
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "ingredients" WHERE inventory_id = $1`)).
					WithArgs(args.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "inventories" WHERE id = $1`)).
					WithArgs(args.ID).
					WillReturnError(errors.New("failed to delete inventory"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &InventoryRepository{
				DB: gormDB,
			}
			if err := r.DeleteInv(tt.args.ctx, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("InventoryRepository.DeleteInv() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestInventoryRepository_FindByID(t *testing.T) {
	now := time.Now()
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
		want    *models.Inventory
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: &models.Inventory{
				ID:          1,
				Item:        "Plastik",
				Qty:         1,
				Uom:         "Kg",
				PricePerQty: 10000,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: false,
			mockFn: func(args args) {

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "inventories" WHERE id = $1 ORDER BY "inventories"."id" LIMIT $2`,
				)).WithArgs(args.ID, 1).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "created_at", "updated_at", "item", "qty", "uom",
						"price_per_qty",
					}).AddRow(
						1, now, now, "Plastik", 1, "Kg",
						10000,
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
					`SELECT * FROM "inventories" WHERE id = $1 ORDER BY "inventories"."id" LIMIT $2`,
				)).WithArgs(args.ID, 1).
					WillReturnError(assert.AnError)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &InventoryRepository{
				DB: gormDB,
			}
			got, err := r.FindByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InventoryRepository.FindByID() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestInventoryRepository_GetAllInv(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	type args struct {
		ctx          context.Context
		objComponent models.ComponentServerSide
	}

	tests := []struct {
		name    string
		args    args
		want    *[]models.Inventory
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "Success - Get all inventory with search and sorting",
			args: args{
				ctx: context.Background(),
				objComponent: models.ComponentServerSide{
					Search:   "plastik",
					SortBy:   "id",
					SortType: "asc",
					Limit:    10,
					Skip:     0,
				},
			},
			want: &[]models.Inventory{
				{ID: 1, Item: "Plastik", Qty: 10, Uom: "Kg", PricePerQty: 10000},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "inventories" WHERE lower\(item\) like '%.*%' ORDER BY id asc LIMIT \$1`).
					WillReturnRows(sqlmock.NewRows([]string{"id", "item", "qty", "uom", "price_per_qty"}).
						AddRow(1, "Plastik", 10, "Kg", 10000))

			},
		},
		{
			name: "error - Get all inventory with search and sorting",
			args: args{
				ctx: context.Background(),
				objComponent: models.ComponentServerSide{
					Search:   "plastik",
					SortBy:   "id",
					SortType: "asc",
					Limit:    10,
					Skip:     0,
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "inventories" WHERE lower\(item\) like '%.*%' ORDER BY id asc LIMIT \$1`).
					WillReturnError(errors.New("query syntax mismatch"))

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &InventoryRepository{
				DB: gormDB,
			}
			got, err := r.GetAllInv(tt.args.ctx, tt.args.objComponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryRepository.GetAllInv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InventoryRepository.GetAllInv() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestInventoryRepository_CountData(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

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
			name: "Success - Count all data",
			args: args{
				ctx: context.Background(),
				objComponent: models.ComponentServerSide{
					Search: "",
				},
			},
			want:    5,
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT count\(\*\) FROM "inventories"`).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))
			},
		},
		{
			name: "Success - Count data with search",
			args: args{
				ctx: context.Background(),
				objComponent: models.ComponentServerSide{
					Search: "plastik",
				},
			},
			want:    3,
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT count\(\*\) FROM "inventories" WHERE lower\(item\) like '%plastik%'`).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))
			},
		},
		{
			name: "Error - Database failure",
			args: args{
				ctx: context.Background(),
				objComponent: models.ComponentServerSide{
					Search: "plastik",
				},
			},
			want:    0,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT count\(\*\) FROM "inventories" WHERE lower\(item\) like '%plastik%'`).
					WillReturnError(errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &InventoryRepository{
				DB: gormDB,
			}
			got, err := r.CountData(tt.args.ctx, tt.args.objComponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryRepository.CountData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InventoryRepository.CountData() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestInventoryRepository_BatchInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)
	type args struct {
		ctx  context.Context
		objs []models.Inventory
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InventoryRepository{
				DB: gormDB,
			}
			if err := r.BatchInsert(tt.args.ctx, tt.args.objs); (err != nil) != tt.wantErr {
				t.Errorf("InventoryRepository.BatchInsert() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
