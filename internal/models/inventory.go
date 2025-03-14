package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Inventory struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Item        string    `json:"item" gorm:"type:varchar(255);not null" validate:"required"`
	Qty         float64   `json:"qty" gorm:"not null" validate:"required"`
	Uom         string    `json:"uom" gorm:"type:varchar(50);not null" validate:"required"`
	PricePerQty float64   `json:"price_per_qty" gorm:"not null" validate:"required"`
	CreatedAt   time.Time `json:"-" `
	UpdatedAt   time.Time `json:"-" `
}

func (*Inventory) TableName() string {
	return "inventories"
}
func (l Inventory) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
