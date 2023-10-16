package models

import (
    "gorm.io/gorm"
)

type Song struct {
	gorm.Model
	ID     int `gorm:"type:int;primary_key"`
	Title string
}