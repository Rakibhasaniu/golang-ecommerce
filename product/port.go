package product

import (
	"main/domain"
	prdHandler "main/rest/handlers/product"
)

type Service interface {
	// CreateProduct(p domain.Product) (*domain.Product, error)
	// GetProducts() ([]domain.Product, error)
	// GetProductById(id int) (*domain.Product, error)
	// UpdateProduct(p domain.Product) (*domain.Product, error)
	// DeleteProduct(id int) (*domain.Product, error)
	prdHandler.Service
}

type ProductRepo interface {
	CreateProduct(p domain.Product) (*domain.Product, error)
	GetProducts(page int, limit int) ([]domain.Product, error)
	CountProducts() (int, error)
	GetProductById(id int) (*domain.Product, error)
	UpdateProduct(p domain.Product) (*domain.Product, error)
	DeleteProduct(id int) (*domain.Product, error)
}
