package services

import "order-managment/src/domain/entities"

type Message map[string]interface{}

type IRabbitMQService interface {
	SendMessage(queue string, message []entities.Product) error
	ReceiveMessage(queue string) (Message, error)
}