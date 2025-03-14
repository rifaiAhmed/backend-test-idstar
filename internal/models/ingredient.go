package models

import "github.com/go-playground/validator/v10"

type Ingredient struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	RecipeID    uint    `json:"recipe_id" gorm:"not null" validate:"required"`
	InventoryID uint    `json:"inventory_id" gorm:"not null" validate:"required"`
	Quantity    float64 `json:"quantity" gorm:"not null" validate:"required"`

	Recipe    *Recipe    `json:"-" gorm:"foreignKey:RecipeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Inventory *Inventory `json:"-" gorm:"foreignKey:InventoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (*Ingredient) TableName() string {
	return "ingredients"
}

func (l Ingredient) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
