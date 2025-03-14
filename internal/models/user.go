package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID        int       `json:"id" `
	Email     string    `json:"email" gorm:"colum:email;type:varchar(100)" validate:"required"`
	CreatedAt time.Time `json:"-" `
	UpdatedAt time.Time `json:"-" `
}

func (*User) TableName() string {
	return "users"
}
func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID           int `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       int       `json:"user_id" gorm:"type:int" validate:"required"`
	Token        string    `json:"token" gorm:"type:text" validate:"required"`
	Uuid         uuid.UUID `json:"uuid" gorm:"type:uuid;unique" validate:"required"`
	TokenExpired time.Time `json:"-" validate:"required"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
