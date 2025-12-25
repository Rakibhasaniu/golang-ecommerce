package product

import (
	"main/rest/middleware"
	"net/http"
)

func (h *Handler) ProductRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Public routes (no authentication required)
	mux.HandleFunc("GET /products", h.GetProduct)
	mux.HandleFunc("GET /products/{id}", h.GetProductById)

	// Protected routes (authentication required)
	mux.Handle("POST /products/create", manager.With(http.HandlerFunc(h.CreateProduct), h.middleware.Auth))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProduct), h.middleware.Auth))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProduct), h.middleware.Auth))
}
