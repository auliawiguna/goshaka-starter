package models

import (
	"time"
)

type ChangeEmail struct {
	// gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `gorm:"not null" json:"user_id"`
	Token     string    `gorm:"size:150;not null" json:"token"`
	OldEmail  string    `gorm:"size:150;not null" json:"old_email"`
	NewEmail  string    `gorm:"size:150;not null" json:"new_email"`
	ExpiredAt time.Time `gorm:"null" json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
