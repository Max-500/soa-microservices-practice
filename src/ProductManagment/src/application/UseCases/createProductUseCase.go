package usecases

import (
	"order-managment/src/domain/entities"
	"order-managment/src/domain/ports"
)

type CreateProductUseCase struct {
	ProductRepository ports.IProduct
}

func NewCreateProductUseCase(productRepository ports.IProduct) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: productRepository,
	}
}

func (useCase *CreateProductUseCase) Run(data []entities.Product) ([]entities.Product, error) {
	return useCase.ProductRepository.Create(data)
}
