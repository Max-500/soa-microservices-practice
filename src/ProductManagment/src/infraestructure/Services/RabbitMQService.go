package services

import (
	"encoding/json"
	"fmt"
	"log"
	services "order-managment/src/application/Services"
	"order-managment/src/domain/entities"
	"order-managment/src/domain/ports"
	"github.com/streadway/amqp"
)

type DataGetProduct struct {
	ProductUuid string `json:"uuid"`
}

type DataUpdateTracking struct {
	ProductUuid string `json:"uuid"`
	Amount      int    `json:"amount"`
}

type RabbitMQService struct {
	Connection        *amqp.Connection
	ProductRepository ports.IProduct
}

func NewRabbitMQService(conn *amqp.Connection, repository ports.IProduct) RabbitMQService {
	return RabbitMQService{
		Connection:        conn,
		ProductRepository: repository,
	}
}

func (r *RabbitMQService) SendMessage(queue string, message []entities.Product) error {
	channel, err := r.Connection.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return err
	}
	defer channel.Close()

	// Declara la cola
	_, err = channel.QueueDeclare(
		queue, // name
		true, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return err
	}

	// Convierte el mensaje a bytes
	body, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Failed to encode message: %v", err)
		return err
	}

	// Publica el mensaje en la cola
	err = channel.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
		return err
	}

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
			log.Printf("Received a message: %s from to queue %s", d.Body, queue)
			if queue == "update_stock_queue" {
				var data DataUpdateTracking
				err := json.Unmarshal(d.Body, &data)
				if err != nil {
					log.Printf("Error decoding JSON: %s", err)
					continue
				}
				r.ProductRepository.UpdateTracking(data.ProductUuid, data.Amount)
			}

			if queue == "send_get_products_queue" {
				var data DataGetProduct
				err := json.Unmarshal(d.Body, &data)
				if err != nil {
					log.Printf("Error decoding JSON: %s", err)
					continue
				}
				result, err := r.ProductRepository.GetProduct(data.ProductUuid)
				println(result)
				if err != nil {
					panic(err)
				}
				r.SendMessage("receive_get_products_queue", result)
			}
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	return nil, nil
}
