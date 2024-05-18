package services

type Message map[string]interface{}

type IRabbitMQService interface {
    SendMessage(queue string, message Message) error
    ReceiveMessage(queue string) (Message, error)
}