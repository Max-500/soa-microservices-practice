package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"order-managment/src/application/UseCases"
	"order-managment/src/domain/entities"
)

type CreateProductController struct {
	UseCase *usecases.CreateProductUseCase
}

func NewCreateProductController(useCase *usecases.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		UseCase: useCase,
	}
}

func (controller *CreateProductController) Run(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Error reading request body",
		})		
		return
	}
	
	var products entities.Products
	err = json.Unmarshal(body, &products)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Error parsing request body",
		})
		return
	}
	
	var data []entities.Product
	for _, product := range products.Products {
		if product.Name == "" || product.Price == 0 || product.Stock == 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "They are sending incomplete data.",
			})
			return
		}

		product := entities.Product{
			Name: product.Name,
			Price: product.Price,
			Stock: product.Stock,
		}
		data = append(data, product)
	}

	if len(data) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "No data was sent.",
		})
		return
	}


	response, err := controller.UseCase.Run(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": response,
	})
}