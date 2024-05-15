package requests

import "order-managment/src/domain/entities"

type CreateProductRequest struct {
	Products []entities.Product `json:"products"`
}