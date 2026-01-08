package product

import (
	"main/rest/middleware"
)

type Handler struct {
	middleware *middleware.Middlewares
	svc        Service
}

func NewHandler(middleware *middleware.Middlewares, svc Service) *Handler {
	return &Handler{middleware: middleware, svc: svc}
}
