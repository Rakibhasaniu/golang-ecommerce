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
	mux.HandleFunc("PUT /products/{id}", handler.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", handler.DeleteProduct)
	mux.HandleFunc("POST /users/create", handler.CreateUser)
	mux.HandleFunc("POST /users/login", handler.Login)
}
