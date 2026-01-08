package product

import "main/domain"

type service struct {
	prdcRepo ProductRepo
}

func NewService(prdcRepo ProductRepo) Service {
	return &service{
		prdcRepo: prdcRepo,
	}
}

func (svc *service) CreateProduct(p domain.Product) (*domain.Product, error) {
	return svc.prdcRepo.CreateProduct(p)
}

func (svc *service) GetProducts(page int, limit int) ([]domain.Product, error) {
	return svc.prdcRepo.GetProducts(page, limit)
}

func (svc *service) CountProducts() (int, error) {
	return svc.prdcRepo.CountProducts()
}

func (svc *service) GetProductById(id int) (*domain.Product, error) {
	return svc.prdcRepo.GetProductById(id)
}

func (svc *service) UpdateProduct(p domain.Product) (*domain.Product, error) {
	return svc.prdcRepo.UpdateProduct(p)
}

func (svc *service) DeleteProduct(id int) (*domain.Product, error) {
	return svc.prdcRepo.DeleteProduct(id)
}
