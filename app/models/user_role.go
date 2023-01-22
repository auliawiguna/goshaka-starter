package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `gorm:"not null" json:"user_id"`
	RoleId    uint      `gorm:"not null" json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role      Role      `gorm:"foreignKey:RoleId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u UserRole) TableName() string {
	return "user_role"
}