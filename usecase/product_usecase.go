package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) GetProductById(product_id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(product_id)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) UpdateProduct(product model.Product) (model.Product, error) {
	product, err := pu.repository.UpdateProduct(product)

	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (pu *ProductUseCase) DeleteProductById(product_id int) error {
	err := pu.repository.DeleteProductById(product_id)

	if err != nil {
		return err
	}
	return nil
}
