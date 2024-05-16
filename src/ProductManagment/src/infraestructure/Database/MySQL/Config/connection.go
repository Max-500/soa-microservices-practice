package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

type MySQL struct {
    DB *sql.DB
    User string
    Password string
    Host string
    Port string
    Name string
}

func (db *MySQL) Connect(dbType string) error {
    // Configura las opciones de cliente
    url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.User, db.Password, db.Host, db.Port, db.Name)
    mysqlDB, err := sql.Open("mysql", url)
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