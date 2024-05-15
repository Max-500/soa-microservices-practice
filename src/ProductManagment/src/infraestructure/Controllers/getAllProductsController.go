package controllers

import (
	"encoding/json"
	"net/http"
	usecases "order-managment/src/application/UseCases"
)

type GetAllProductsController struct {
	UseCase usecases.GetAllProductsUseCase
}

func NewGetAllProductsController(useCase usecases.GetAllProductsUseCase) GetAllProductsController {
	return GetAllProductsController{
		UseCase: useCase,
	}
}

func (controller GetAllProductsController) Run(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := controller.UseCase.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(response) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No hay registros"})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"data": response})
	}
}


