package entity

import "time"

type Comment struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	PhotoID   int
	Message   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
