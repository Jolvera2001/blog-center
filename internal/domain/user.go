package domain

import "time"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"uniqueIndex;size:255"`
	CreatedAt time.Time
}
