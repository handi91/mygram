package request

type PostComment struct {
	Message string `json:"message" valid:"required~message is required"`
	PhotoID int    `json:"photo_id" valid:"required~photo id is required"`
}

type UpdateComment struct {
	Message string `json:"message" valid:"required~message is required"`
}
