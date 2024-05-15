package routes

import (
	usecases "order-managment/src/application/UseCases"
	controllers "order-managment/src/infraestructure/Controllers"
	configMongo "order-managment/src/infraestructure/Database/MongoDB/Config"
	configMySQL "order-managment/src/infraestructure/Database/MySQL/Config"
	repositories "order-managment/src/infraestructure/Repositories"

	"github.com/gorilla/mux"
)

func InitRoutes(dbType string) *mux.Router {
	r := mux.NewRouter()

	mongo := &configMongo.MongoDB{}
	err := mongo.Connect(dbType)
	if err != nil {
		println(err)
	}

	mysql := &configMySQL.MySQL{}
	err = mysql.Connect(dbType)
	if err != nil {
		println(err)
	}
	productRepository := repositories.NewProductRepository(dbType, mysql, mongo)

	createProductUseCase := usecases.NewCreateProductUseCase(productRepository)
	createProductController := controllers.NewCreateProductController(createProductUseCase)

	deleteProductUseCase := usecases.NewDeleteProductUseCase(productRepository)
	deleteProductController := controllers.NewDeleteProductController(deleteProductUseCase)

	r.HandleFunc("/", createProductController.Run).Methods("POST")

	r.HandleFunc("/", deleteProductController.Run).Methods("DELETE")

	return r
}
