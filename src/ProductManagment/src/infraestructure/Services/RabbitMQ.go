package services

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

func Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}
	fmt.Println("Conexión exitosa a RabbitMQ")
	return conn, nil
}