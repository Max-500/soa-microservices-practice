package requests

type DeleteProductRequest struct {
	Product string `json:"product"`
}

type DeleteProductsRequest struct {
	Products []DeleteProductRequest `json:"products"`
}