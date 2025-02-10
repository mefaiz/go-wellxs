package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID             uint   `gorm:"primarykey"`
	UserID         uint   `gorm:"not null"`
	User           User   `gorm:"foreignKey:UserID"`
	FwCode         string `gorm:"size:50"`
	IcNo           string `gorm:"size:50"`
	PassportExpiry *time.Time
	Country        string `gorm:"size:100"`
	BirthDate      *time.Time
	Height         float32 `gorm:"type:decimal(5,2)"`
	Weight         float32 `gorm:"type:decimal(5,2)"`
	AddressID      *uint   // Optional relationship
	Address        Address `gorm:"foreignKey:AddressID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
