package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Title     string    `json:"title"`
	SubTitle  string    `json:"subtitle"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
