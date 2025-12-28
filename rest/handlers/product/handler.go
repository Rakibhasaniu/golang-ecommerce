package product

import (
	"main/repo"
	"main/rest/middleware"
)

type Handler struct {
	middleware  *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(middleware *middleware.Middlewares, productRepo repo.ProductRepo) *Handler {
	return &Handler{middleware: middleware, productRepo: productRepo}
}
