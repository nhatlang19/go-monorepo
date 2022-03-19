package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique, not null"`
	Password string `json:"-"`
}