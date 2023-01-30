package models

import (
	"time"
)

type Note struct {
	// gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	SubTitle  string    `json:"subtitle"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
