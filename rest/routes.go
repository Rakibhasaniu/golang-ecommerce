package rest

import (
	handler "main/rest/handlers"
	"main/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.HandleFunc("GET /products", handler.GetProduct)
	mux.HandleFunc("POST /products/create", handler.CreateProduct)
	mux.HandleFunc("GET /products/{id}", handler.GetProductById)
}
