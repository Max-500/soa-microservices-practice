package main

import (
	"log"
	"net/http"
	routes "order-managment/src/infraestructure/Routes"
)

func main() {
    dbType := "MySQL" // Es MongoDB o MySQL
    r := routes.InitRoutes(dbType)
    log.Fatal(http.ListenAndServe(":8080", r))
}
