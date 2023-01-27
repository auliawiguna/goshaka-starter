package models

import (
	"time"
)

type Permission struct {
	// gorm.Model
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"size:50;not null;unique" json:"name"`
	Display        string    `gorm:"size:150;not null" json:"display"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	PermissionRole []PermissionRole
}
