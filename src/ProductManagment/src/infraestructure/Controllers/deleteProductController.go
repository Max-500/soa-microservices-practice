package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	usecases "order-managment/src/application/UseCases"
	requests "order-managment/src/infraestructure/Controllers/DTO/Requests"
	"strings"
)

type DeleteProductController struct {
	UseCase *usecases.DeleteProductUseCase
}

func NewDeleteProductController(useCase *usecases.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{
		UseCase: useCase,
	}
}

func (controller *DeleteProductController) Run(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Error reading request body",
		})
		return
	}

	var uuids requests.DeleteProductsRequest
	err = json.Unmarshal(body, &uuids)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Error parsing request body",
		})
		return
	}

	var data []requests.DeleteProductRequest
	for _, product := range uuids.Products {
		if product.Product == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "They are sending incomplete data.",
			})
			return
		}
		uuid := requests.DeleteProductRequest{
			Product: product.Product,
		}
		data = append(data, uuid)
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

	if strings.Contains(response, "No se pudieron eliminar los productos") {
		w.WriteHeader(http.StatusNotFound)
	} else if response == "Algo salio mal" || response == "Tipo de base de datos no soportado" {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": response,
	})
}

