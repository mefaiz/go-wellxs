package models

import (
	"time"
	"gorm.io/gorm"

)

type UserSubscription struct {
	ID 				uint `gorm:"primarykey" json:"id"`
	UserID 			uint `gorm:"not null"`
	User 			User `gorm:"foreignKey:UserID"`
	SubscriptionID 	uint `gorm:"not null"`
	Subscription 	Subscription `gorm:"foreignKey:SubscriptionID"`
	StartDate 		time.Time `gorm:"autoCreateTime" json:"start_date"`
	EndDate 		time.Time `gorm:"autoCreateTime" json:"end_date"`
	CreatedAt 		time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 		time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}