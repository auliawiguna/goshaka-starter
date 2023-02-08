package models

import (
	"time"
)

type File struct {
	// gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `gorm:"not null" json:"user_id"`
	Filename  string    `gorm:"not null;field:file_name" json:"filename"`
	Mimetype  string    `gorm:"not null;field:mime_type" json:"mimetype"`
	Path      string    `gorm:"not null" json:"path"`
	Size      int       `gorm:"not null" json:"size"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
