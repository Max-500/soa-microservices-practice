package usecases

import (
	"order-managment/src/domain/entities"
	"order-managment/src/domain/ports"
)

type GetAllProductsUseCase struct {
	ProductRepository ports.IProduct
}

func NewGetAllProductsUseCase(productRepository ports.IProduct) GetAllProductsUseCase {
	return GetAllProductsUseCase{
		ProductRepository: productRepository,
	}
}

func (useCase GetAllProductsUseCase) Run() ([]entities.Product, error) {
	return useCase.ProductRepository.GetAllProducts()
}
