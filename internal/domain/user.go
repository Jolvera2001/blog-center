package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"uniqueIndex;size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}

type IUserRepository interface {
	Create(user *User) error
	FindByID(uuid string) (*User, error)
	Update(user *User) error
	Delete(uuid string) error
}
