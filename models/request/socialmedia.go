package request

type SocialMediaRequest struct {
	Name           string `json:"name" valid:"required~name is required"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~social media url is required"`
}
