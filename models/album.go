package models

type Album struct {
	ID     int `gorm:"type:int;primary_key"`
	Title  string `gorm:"type:varchar(255)"`
}