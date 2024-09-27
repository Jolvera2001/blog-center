package domain

import "time"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"uniqueIndex;size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
}

type IUserRepository interface {
	Create(user *User) error
	FindByID(id int) (*User, error)
	Update(id int) error
	Delete(id int) error
}
