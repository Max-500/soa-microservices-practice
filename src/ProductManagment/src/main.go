package main

import (
	"log"
	"net/http"
	routes "order-managment/src/infraestructure/Routes"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	dbTypeEnv := os.Getenv("DB_TYPE")
	
	dbNameEnv := os.Getenv("DB_NAME")

    var dbName string

	if dbNameEnv != "" {
		dbName = dbNameEnv
	}else{
		dbName = "mydatabase"
	}

    var dbType string

	if dbTypeEnv != "" {
		dbType = dbTypeEnv
	}else{
		dbType = "MySQL"
	}

	var port string

	portEnv := os.Getenv("PRODUCT_PORT")
	if portEnv != ""{
		port = portEnv
	}else{
		port = "8080"
	}

	var dbHost string
	hostEnv := os.Getenv("DB_HOST")
	if hostEnv != "" {
		dbHost = hostEnv
	} else {
		dbHost = "localhost"
	}

	var dbPort string
	portDbEnv := os.Getenv("DB_PORT")
	if portDbEnv != "" {
		dbPort = portDbEnv
	} else {
		if dbType == "MySQL" {
			dbPort = "3306"
		}else{
			dbPort = "27017"
		}
	}

	var dbUser string
	dbUserEnv := os.Getenv("DB_USER")
	if dbUserEnv != "" {
		dbUser = dbUserEnv
	} else {
		dbUser = "root"
	}

	var dbPassword string
	dbPasswordEnv := os.Getenv("DB_PASSWORD")
	if dbPasswordEnv != "" {
		dbPassword = dbPasswordEnv
	} else {
		dbPassword = ""
	}

	url := fmt.Sprintf(":%s", port)
    r := routes.InitRoutes(dbType, dbName, dbHost, dbPort, dbUser, dbPassword)
    log.Fatal(http.ListenAndServe(url, r))
}
