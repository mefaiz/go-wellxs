package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID        uint   `gorm:"primarykey"`
	AddressLine1    string `gorm:"size:255"`
	AddressLine2    string `gorm:"size:255"`
	AddressLine3    string `gorm:"size:255"`
	City      string `gorm:"size:100"`
	State     string `gorm:"size:100"`
	Country   string `gorm:"size:100"`
	PostCode  string `gorm:"size:20"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
