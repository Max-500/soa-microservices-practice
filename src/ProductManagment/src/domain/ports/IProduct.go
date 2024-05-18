package ports

import (
	"order-managment/src/domain/entities"
	requests "order-managment/src/infraestructure/Controllers/DTO/Requests"
)

type IProduct interface {
	Create([]entities.Product) ([]entities.Product, error)
	Delete(uuid []requests.DeleteProductRequest) (string, error)
	UpdateTracking(uuid string, amount int) (string, error)
	GetAllProducts() ([]entities.Product, error)
}