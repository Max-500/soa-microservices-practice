package main

import (
	"log"
	"net/http"
	routes "order-managment/src/infraestructure/Routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	dbTypeEnv := os.Getenv("DB_TYPE")

    var dbType string

	if dbTypeEnv != "" {
		dbType = dbTypeEnv
	}else{
		dbType = "MySQL"
	}

    r := routes.InitRoutes(dbType)
    log.Fatal(http.ListenAndServe(":8080", r))
}
