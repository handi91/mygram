package response

import "time"

type PostSocialMedia struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type GetSocialMedia struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	SocialMediaUrl string          `json:"social_media_url"`
	UserID         int             `json:"user_id"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	User           SocialMediaUser `json:"User"`
}

type SocialMediaUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UpdateSocialMedia struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}
