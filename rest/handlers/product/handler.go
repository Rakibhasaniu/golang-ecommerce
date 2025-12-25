package product

import "main/rest/middleware"

type Handler struct {
	middleware *middleware.Middlewares
}

func NewHandler(middleware *middleware.Middlewares) *Handler {
	return &Handler{middleware: middleware}
}
