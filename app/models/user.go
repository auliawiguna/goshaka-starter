package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey" json:"id"`
	Username    string    `gorm:"size:50;not null;unique" json:"username"`
	Email       string    `gorm:"size:150;not null;unique" json:"email"`
	Password    string    `gorm:"size:150;not null" json:"password"`
	ValidatedAt time.Time `json:"validated_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	HashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	if u.Password == "" {
		return nil
	}

	u.Password = string(HashedPassword)

	return nil
}