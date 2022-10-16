package controller

import (
	"mygram-api/service"
)

type Controller struct {
	service service.Service
}

func New(s service.Service) Controller {
	return Controller{
		service: s,
	}
}
