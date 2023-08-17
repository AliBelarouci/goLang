package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	ID              uint   `gorm:"primaryKey"`
	CallTime        string `gorm:"not null"`
	CallDisposition string `gorm:"not null"`
	Phone           string `gorm:"not null"`
	FirstName       string `gorm:"not null"`
	LastName        string
	Address1        string
	Address2        string
	City            string
	State           string
	Zip             string
}

type File struct {
	gorm.Model
	ID               uint   `gorm:"primaryKey"`
	ImportedDate     string `gorm:"not null"`
	Filename         string `gorm:"not null" validate:"required"`
	TotalRowCount    int64  `gorm:"not null"`
	ImportedRowCount int64  `gorm:"not null"`
	Hash             string `gorm:"not null"`
}
