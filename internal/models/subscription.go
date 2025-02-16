package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Plan      string `gorm:"size:255;not null" json:"plan"`
	OUM  	  string `gorm:"size:255;not null" json:"oum"`
	Price     float64 `gorm:"not null" json:"price"`
	Role      string `gorm:"size:255;not null" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}