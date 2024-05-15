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
	err := mongo.Connect()
	if err != nil {
		println(err)
	}

	mysql := &configMySQL.MySQL{}
	err = mysql.Connect()
	if err != nil {
		println(err)
	}
	productRepository := repositories.NewProductRepository(dbType, mysql, mongo)

	productUseCase := usecases.NewCreateProductUseCase(productRepository)
	productController := controllers.NewCreateProductController(productUseCase)

	r.HandleFunc("/", productController.Run).Methods("POST")

	return r
}
