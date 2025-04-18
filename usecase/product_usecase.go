package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (p *ProductUsecase) GetProducts() ([]model.Product, error) {
	return p.repository.GetProducts()
}

func (p *ProductUsecase) CreateProduct(product model.Product) error {
	return p.repository.CreateProduct(product)
}

func (p *ProductUsecase) GetProductByID(id int) (*model.Product, error) {
	return p.repository.GetProductByID(id)
}
