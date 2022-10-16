package entity

import "time"

type Photo struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Caption   string
	PhotoUrl  string `gorm:"not null"`
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
