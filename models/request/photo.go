package request

type PhotoRequest struct {
	Title    string `json:"title" valid:"required~Title is required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" valid:"required~Photo URL is required"`
}
