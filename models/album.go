package models

import (
    "gorm.io/gorm"
)

type Album struct {
	gorm.Model
	ID     int `gorm:"type:int;primary_key"`
	Title  string `gorm:"type:varchar(255)"`
	Songs []Song `gorm:"foreignKey:ID"`
}