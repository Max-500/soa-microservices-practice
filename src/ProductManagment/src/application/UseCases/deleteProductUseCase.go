package usecases

import (
	"order-managment/src/domain/ports"
	requests "order-managment/src/infraestructure/Controllers/DTO/Requests"
)

type DeleteProductUseCase struct {
	ProductRepository ports.IProduct
}

func NewDeleteProductUseCase(productRepository ports.IProduct) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		ProductRepository: productRepository,
	}
}

func (useCase *DeleteProductUseCase) Run (uuids []requests.DeleteProductRequest) (string, error) {
	return useCase.ProductRepository.Delete(uuids)
}