package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `gorm:"size:100;not null"`
	Description string   `gorm:"type:text"`
	SKU         string   `gorm:"size:50;uniqueIndex"`
	Price       float64  `gorm:"type:decimal(10,2);not null"`
	CategoryID  uint     `gorm:"index"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	Inventory   Inventory
}

type Category struct {
	gorm.Model
	Name        string    `gorm:"size:50;not null;uniqueIndex"`
	Description string    `gorm:"size:255"`
	Products    []Product `gorm:"foreignKey:CategoryID"`
}

type Inventory struct {
	gorm.Model
	ProductID    uint `gorm:"uniqueIndex"`
	Quantity     int  `gorm:"default:0"`
	ReorderLevel int  `gorm:"default:10"`
}
