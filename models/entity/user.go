package entity

import "time"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"type:varchar;unique;not null"`
	Email     string `gorm:"type:varchar;unique;not null"`
	Password  string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
