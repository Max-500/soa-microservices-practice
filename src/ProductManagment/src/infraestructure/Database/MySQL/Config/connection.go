package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

type MySQL struct {
    DB *sql.DB
}

func (db *MySQL) Connect(dbType string) error {
    // Configura las opciones de cliente
    mysqlDB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mydatabase")
    if err != nil {
        log.Fatal(err)
    }

    // Comprueba la conexión
    err = mysqlDB.Ping()
    if err != nil {
        log.Fatal(err)
    }

    db.DB = mysqlDB

    if dbType == "MySQL"{
        fmt.Println("¡Conectado a MySQL!")
    }
    return nil
}