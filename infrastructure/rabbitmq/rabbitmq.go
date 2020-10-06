package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	User     string
	Password string
	port     string
	url      string
}

func (mq *RabbitMQ) Publish(payload string) {

}

func (mq *RabbitMQ) CreateConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	mq.failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
}

func (mq *RabbitMQ) failOnError(err error, msg string) {
	if err != nil {
		log.Println("%s: %s", msg, err)
	}
}

func NewRabbitMQ() *RabbitMQ {
	return &RabbitMQ{}
}
