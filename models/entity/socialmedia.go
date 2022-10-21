package entity

import "time"

type SocialMedia struct {
	ID             int    `gorm:"primaryKey"`
	Name           string `gorm:"type:varchar;not null"`
	SocialMediaUrl string `gorm:"not null"`
	UserID         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
