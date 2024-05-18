package services

import (
	"encoding/json"
	"fmt"
	"log"
	services "order-managment/src/application/Services"
	"order-managment/src/domain/ports"

	"github.com/streadway/amqp"
)

type DataUpdateTracking struct {
	ProductUuid string `json:"uuid"`
	Amount int `json:"amount"`
}

type RabbitMQService struct {
	Connection *amqp.Connection
	ProductRepository ports.IProduct
}

func NewRabbitMQService(conn *amqp.Connection, repository ports.IProduct) (RabbitMQService) {
	return RabbitMQService{
		Connection: conn,
		ProductRepository: repository,
	}
}

func (r *RabbitMQService) SendMessage(queue string, message services.Message) error {
	fmt.Println("Mensaje enviado a la cola", queue, ":", message)
	return nil
}

func (r *RabbitMQService) ReceiveMessage(queue string) (services.Message, error) {
	channel, err := r.Connection.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		panic(err)
	}

	msgs, err := channel.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		panic(err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var data DataUpdateTracking
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}
			log.Print(data)
			r.ProductRepository.UpdateTracking(data.ProductUuid, data.Amount)
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	return nil, nil
}

