package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string  `gorm:"not null"`
	Author string  `gorm:"not null"`
	Price  float64 `gorm:"not null"`
}