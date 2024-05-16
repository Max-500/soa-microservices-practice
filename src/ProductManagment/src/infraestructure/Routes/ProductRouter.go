package routes

import (
	usecases "order-managment/src/application/UseCases"
	controllers "order-managment/src/infraestructure/Controllers"
	configMongo "order-managment/src/infraestructure/Database/MongoDB/Config"
	configMySQL "order-managment/src/infraestructure/Database/MySQL/Config"
	repositories "order-managment/src/infraestructure/Repositories"
	"github.com/gorilla/mux"
)

func InitRoutes(dbType string, dbName string, dbHost string, dbPort string, dbUser string, dbPassword string) *mux.Router {
	r := mux.NewRouter()
	mongo := &configMongo.MongoDB{
		Name: dbName,
		Host: dbHost,
		Port: dbPort,
	}

	if dbType == "MongoDB" {
		err := mongo.Connect(dbType)
		if err != nil {
			println(err)
		}
	}
	
	mysql := &configMySQL.MySQL{
		User: dbUser,
		Password: dbPassword,
		Host: dbHost,
		Port: dbPort,
		Name: dbName,
	}
	
	if dbType == "MySQL" {
		err := mysql.Connect(dbType)
		if err != nil {
			println(err)
		}
	}

	productRepository := repositories.NewProductRepository(dbType, mysql, mongo, dbName)

	createProductUseCase := usecases.NewCreateProductUseCase(productRepository)
	createProductController := controllers.NewCreateProductController(createProductUseCase)

	deleteProductUseCase := usecases.NewDeleteProductUseCase(productRepository)
	deleteProductController := controllers.NewDeleteProductController(deleteProductUseCase)

	getAllProductsUseCase := usecases.NewGetAllProductsUseCase(productRepository)
	getAllProductsController := controllers.NewGetAllProductsController(getAllProductsUseCase)

	r.HandleFunc("/", createProductController.Run).Methods("POST")
	r.HandleFunc("/", deleteProductController.Run).Methods("DELETE")
	r.HandleFunc("/", getAllProductsController.Run).Methods("GET")

	return r
}
