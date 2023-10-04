package controller

import "github.com/ImranZahoor/blog-api/internal/service"

type (
	Controller struct {
		service service.Service
	}
)

func NewController(svc service.Service) Controller {
	return Controller{service: svc}
}
