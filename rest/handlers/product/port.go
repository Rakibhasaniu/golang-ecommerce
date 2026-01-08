package product

import "main/domain"

type Service interface {
	CreateProduct(domain.Product) (*domain.Product, error)
	GetProducts(page int, limit int) ([]domain.Product, error)
	CountProducts() (int, error)
	GetProductById(id int) (*domain.Product, error)
	UpdateProduct(domain.Product) (*domain.Product, error)
	DeleteProduct(id int) (*domain.Product, error)
}
