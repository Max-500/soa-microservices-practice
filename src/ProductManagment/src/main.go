package main

import (
	"log"
	"net/http"
	routes "order-managment/src/infraestructure/Routes"
)

func main() {
    dbType := "MongoDB" // Es MongoDB o MySQL
    r := routes.InitRoutes(dbType)
    log.Fatal(http.ListenAndServe(":8080", r))
}
