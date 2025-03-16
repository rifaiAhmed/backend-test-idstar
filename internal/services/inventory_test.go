package services

import (
	"backend-test/internal/mocks"
	"backend-test/internal/models"
	"context"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/xuri/excelize/v2"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestInventoryService_InsertInv(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIInventoryRepository(ctrlMock)

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
			name: "success - insert new inventory",
			args: args{
				ctx: context.Background(),
				obj: &models.Inventory{
					ID:   1,
					Item: "Beras Premium",
					Qty:  100,
				},
			},
			want: &models.Inventory{
				ID:   1,
				Item: "Beras Premium",
				Qty:  100,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					InsertInv(args.ctx, args.obj).
					Return(args.obj, nil)
			},
		},
		{
			name: "error - repository error",
			args: args{
				ctx: context.Background(),
				obj: &models.Inventory{
					ID:   2,
					Item: "Gula Pasir",
					Qty:  50,
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					InsertInv(args.ctx, args.obj).
					Return(nil, errors.New("database error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &InventoryService{
				InventoryRepo: mockRepo,
			}
			got, err := s.InsertInv(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.InsertInv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InventoryService.InsertInv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInventoryService_UpdateInv(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIInventoryRepository(ctrlMock)
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
			name: "success - update inventory",
			args: args{
				ctx: context.Background(),
				obj: &models.Inventory{
					ID:          1,
					Item:        "Beras Premium",
					PricePerQty: 12000,
					Qty:         50,
					Uom:         "kg",
				},
			},
			want: &models.Inventory{
				ID:          1,
				Item:        "Beras Premium",
				PricePerQty: 12000,
				Qty:         50,
				Uom:         "kg",
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					FindByID(args.ctx, int(args.obj.ID)).
					Return(&models.Inventory{
						ID:          1,
						Item:        "Beras Medium",
						PricePerQty: 10000,
						Qty:         40,
						Uom:         "kg",
					}, nil)

				mockRepo.EXPECT().
					UpdateInv(args.ctx, gomock.Any()).
					Return(args.obj, nil)
			},
		},
		{
			name: "error - inventory not found",
			args: args{
				ctx: context.Background(),
				obj: &models.Inventory{
					ID:          99,
					Item:        "Minyak Goreng",
					PricePerQty: 20000,
					Qty:         20,
					Uom:         "liter",
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					FindByID(args.ctx, int(args.obj.ID)).
					Return(nil, gorm.ErrRecordNotFound)
			},
		},
		{
			name: "error - failed to update inventory",
			args: args{
				ctx: context.Background(),
				obj: &models.Inventory{
					ID:          2,
					Item:        "Gula Pasir",
					PricePerQty: 15000,
					Qty:         30,
					Uom:         "kg",
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					FindByID(args.ctx, int(args.obj.ID)).
					Return(&models.Inventory{
						ID:          2,
						Item:        "Gula Pasir",
						PricePerQty: 14000,
						Qty:         25,
						Uom:         "kg",
					}, nil)

				mockRepo.EXPECT().
					UpdateInv(args.ctx, gomock.Any()).
					Return(nil, errors.New("database error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &InventoryService{
				InventoryRepo: mockRepo,
			}
			got, err := s.UpdateInv(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.UpdateInv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InventoryService.UpdateInv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInventoryService_DeleteInv(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIInventoryRepository(ctrlMock)
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
			name: "success - delete inventory",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					DeleteInv(args.ctx, args.ID).
					Return(nil)
			},
		},
		{
			name: "fail - inventory not found",
			args: args{
				ctx: context.Background(),
				ID:  999,
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					DeleteInv(args.ctx, args.ID).
					Return(errors.New("inventory not found"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &InventoryService{
				InventoryRepo: mockRepo,
			}
			if err := s.DeleteInv(tt.args.ctx, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.DeleteInv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInventoryService_GetAllInv(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIInventoryRepository(ctrlMock)

	type args struct {
		ctx          context.Context
		objComponent models.ComponentServerSide
	}

	mockData := []models.Inventory{
		{ID: 1, Item: "Item A", Qty: 10, Uom: "pcs", PricePerQty: 1000},
		{ID: 2, Item: "Item B", Qty: 5, Uom: "box", PricePerQty: 2000},
	}

	tests := []struct {
		name    string
		args    args
		want    *[]models.Inventory
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - get all inventory",
			args: args{
				ctx:          context.Background(),
				objComponent: models.ComponentServerSide{Limit: 10, Offset: 0},
			},
			want:    &mockData,
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					GetAllInv(args.ctx, args.objComponent).
					Return(&mockData, nil)
			},
		},
		{
			name: "fail - database error",
			args: args{
				ctx:          context.Background(),
				objComponent: models.ComponentServerSide{Limit: 10, Offset: 0},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					GetAllInv(args.ctx, args.objComponent).
					Return(nil, errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockFn(tt.args)

			s := &InventoryService{
				InventoryRepo: mockRepo,
			}

			got, err := s.GetAllInv(tt.args.ctx, tt.args.objComponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.GetAllInv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InventoryService.GetAllInv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInventoryService_CountData(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIInventoryRepository(ctrlMock)

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
			name: "success - count inventory data",
			args: args{
				ctx:          context.Background(),
				objComponent: models.ComponentServerSide{Limit: 10, Offset: 0},
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
			name: "fail - database error",
			args: args{
				ctx:          context.Background(),
				objComponent: models.ComponentServerSide{Limit: 10, Offset: 0},
			},
			want:    0,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					CountData(args.ctx, args.objComponent).
					Return(int64(0), errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockFn(tt.args)

			s := &InventoryService{
				InventoryRepo: mockRepo,
			}

			got, err := s.CountData(tt.args.ctx, tt.args.objComponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.CountData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InventoryService.CountData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInventoryService_InsertFromExcel(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIInventoryRepository(ctrlMock)

	type args struct {
		ctx      context.Context
		filePath string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - insert from valid Excel file",
			args: args{
				ctx:      context.Background(),
				filePath: "test_inventory.xlsx",
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					BatchInsert(args.ctx, gomock.Any()).
					Return(nil).
					Times(1)
			},
		},
		{
			name: "fail - file not found",
			args: args{
				ctx:      context.Background(),
				filePath: "invalid_file.xlsx",
			},
			wantErr: true,
			mockFn: func(args args) {
				// not found
			},
		},
		{
			name: "fail - BatchInsert returns error",
			args: args{
				ctx:      context.Background(),
				filePath: "test_inventory.xlsx",
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					BatchInsert(args.ctx, gomock.Any()).
					Return(errors.New("database error")).
					Times(1)
			},
		},
	}

	testFilePath := "test_inventory.xlsx"
	createTestExcelFile(testFilePath)

	defer os.Remove(testFilePath)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &InventoryService{
				InventoryRepo: mockRepo,
			}

			err := s.InsertFromExcel(tt.args.ctx, tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.InsertFromExcel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func createTestExcelFile(filePath string) {
	f := excelize.NewFile()
	sheetName := "Sheet1"
	f.SetSheetName(f.GetSheetName(1), sheetName)

	_ = f.SetCellValue(sheetName, "A1", "Item")
	_ = f.SetCellValue(sheetName, "B1", "Qty")
	_ = f.SetCellValue(sheetName, "C1", "UOM")
	_ = f.SetCellValue(sheetName, "D1", "PricePerQty")

	_ = f.SetCellValue(sheetName, "A2", "Laptop")
	_ = f.SetCellValue(sheetName, "B2", 10)
	_ = f.SetCellValue(sheetName, "C2", "pcs")
	_ = f.SetCellValue(sheetName, "D2", 15000.5)

	_ = f.SaveAs(filePath)
}
