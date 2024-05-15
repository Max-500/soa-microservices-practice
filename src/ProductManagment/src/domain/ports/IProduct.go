package ports

import (
	"order-managment/src/domain/entities"
	"github.com/google/uuid"
)

type IProduct interface {
	Create([]entities.Product) ([]entities.Product, error)
	Delete(id uuid.UUID) (string, error)
	UpdateTracking(id uuid.UUID) (string, error)
	GetAllProducts() ([]entities.Product, error)
}
