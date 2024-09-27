package domain

import "time"

type User struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"uniqueIndex;size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
}

type IUserRepository interface {
	Create(user *User) error
	FindByID(uuid string) (*User, error)
	Update(user *User) error
	Delete(uuid string) error
}
