package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);unique;not null"`
	Category string `gorm:"type:varchar(100);not null"`
	Price    int    `gorm:"type:int;not null"`
	Stock    int    `gorm:"type:int;not null"`
}
