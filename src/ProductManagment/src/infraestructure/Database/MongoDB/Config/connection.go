package config

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
    Client *mongo.Client
}

func (db *MongoDB) Connect() error {
    // Establece un contexto
    ctx := context.TODO()

    // Configura las opciones de cliente
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/mydatabase")

    // Conéctate a MongoDB
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Comprueba la conexión
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    db.Client = client

    fmt.Println("¡Conectado a MongoDB!")
    return nil
}