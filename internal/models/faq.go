package models

import (
	"time"
	"gorm.io/gorm"
)

type Faq struct {
	ID uint `gorm:"primarykey" json:"id"`
	Question string `gorm:"size:255;not null" json:"question"`
	Answer string `gorm:"size:255;not null" json:"answer"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}