package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	Phone     string `gorm:"size:255;unique" json:"phone"`
	Password  string `gorm:"size:255;not null" json:"password"`
	FcmToken  string `gorm:"size:255" json:"fcm_token"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
