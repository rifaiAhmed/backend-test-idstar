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

type RecipeFormat struct {
	ID   uint    `json:"id"`
	Name string  `json:"name"`
	SKU  string  `json:"sku"`
	Cogs float64 `json:"cogs"`
}

type IngredientCustom struct {
	ID          int     `json:"id"`
	InventoryID int     `json:"inventory_id"`
	Quantity    float64 `json:"quantity"`
	Item        string  `json:"item"`
	Uom         string  `json:"uom"`
}
