package infraestructure

import (
	"log"
	"github.com/streadway/amqp"
	"fmt"
	"encoding/json"

)

type RabbitMQPublisher struct {
	Channel *amqp.Channel
}

func NewRabbitMQPublisher() (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial("amqp://guest:guest@3.225.46.249:5672/") 
	if err != nil {
		return nil, fmt.Errorf("error conectando a RabbitMQ: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("error creando canal de RabbitMQ: %v", err)
	}

	return &RabbitMQPublisher{Channel: channel}, nil
}

func (r *RabbitMQPublisher) Publish(event string, data interface{}) error {
	_, err := r.Channel.QueueDeclare(
		"flight_created",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("error declarando la cola: %v", err)
	}

	// Convertir a JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error convirtiendo a JSON: %v", err)
	}

	// Publicar
	err = r.Channel.Publish(
		"",
		"flight_created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBytes,
		},
	)
	if err != nil {
		return fmt.Errorf("error publicando el mensaje: %v", err)
	}

	log.Println("Mensaje publicado a RabbitMQ:", string(jsonBytes))
	return nil
}
