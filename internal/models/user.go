package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `gorm:"size:100;not null"`
	Email    string  `gorm:"size:255;not null;uniqueIndex"`
	Password string  `gorm:"size:255;not null"`
	Profile  Profile `gorm:"foreignKey:UserID"`
	Roles    []*Role `gorm:"many2many:user_roles;"`
}

type Profile struct {
	gorm.Model
	UserID      uint   `gorm:"not null;uniqueIndex"`
	Avatar      string `gorm:"size:255"`
	Address     string `gorm:"size:255"`
	PhoneNumber string `gorm:"size:20"`
	DateOfBirth string `gorm:"type:date"`
}

type Role struct {
	gorm.Model
	Name        string  `gorm:"size:50;not null;uniqueIndex"`
	Description string  `gorm:"size:255"`
	Users       []*User `gorm:"many2many:user_roles;"`
}
