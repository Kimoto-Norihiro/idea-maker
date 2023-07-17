package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UID	 string `json:"uid" gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email"`
}