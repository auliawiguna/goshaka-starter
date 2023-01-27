package models

import (
	"time"
)

type PermissionRole struct {
	// gorm.Model
	ID           uint       `gorm:"primaryKey" json:"id"`
	PermissionId uint       `gorm:"not null" json:"permission_id"`
	RoleId       uint       `gorm:"not null" json:"role_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Permission   Permission `gorm:"foreignKey:PermissionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role         Role       `gorm:"foreignKey:RoleId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u PermissionRole) TableName() string {
	return "permission_role"
}
