package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Recipe struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"type:varchar(50);unique;not null" validate:"required"`
	SKU         string       `json:"sku" gorm:"type:varchar(50);unique;not null"`
	Ingredients []Ingredient `json:"ingredients" gorm:"foreignKey:RecipeID"`
	CreatedAt   time.Time    `json:"-" `
	UpdatedAt   time.Time    `json:"-" `
}

func (*Recipe) TableName() string {
	return "recipes"
}

func (l Recipe) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
