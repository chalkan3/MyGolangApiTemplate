package database

import (
	linq "github.com/ahmetb/go-linq"
	"o2b.com.br/whatsAppApi/domain/helpers"

	"github.com/streadway/amqp"
)

// Queue rabbitmq queue
type Queue struct {
	Name             string `json:"name,omitempty"`
	Durable          bool   `json:"durable,omitempty"`
	DeleteWhenUnused bool   `json:"delete_when_unused,omitempty"`
	Exclusive        bool   `json:"exclusive,omitempty"`
	NoWait           bool   `json:"no_wait,omitempty"`
}

// RabbitMQ rabbit mq
type RabbitMQ struct {
	User     string  `json:"user,omitempty"`
	Password string  `json:"password,omitempty"`
	Port     string  `json:"Port,omitempty"`
	Host     string  `json:"Host,omitempty"`
	Queue    []Queue `json:"queue,omitempty"`
}

func (mq *RabbitMQ) connectionChannel() (*amqp.Channel, *amqp.Connection) {
	conn := mq.GetConnection()
	ch, err := conn.Channel()
	helpers.FailOnError(err, "Failed to connect to chanel")

	return ch, conn

}
func (mq *RabbitMQ) declareQueue() (*amqp.Queue, *amqp.Channel, *amqp.Connection) {
	processQueue := linq.From(mq.Queue).Where(func(c interface{}) bool {
		return c.(Queue).Name == "process"
	}).First().(Queue)

	ch, conn := mq.connectionChannel()

	q, err := ch.QueueDeclare(
		processQueue.Name,
		processQueue.Durable,
		processQueue.DeleteWhenUnused,
		processQueue.Exclusive,
		processQueue.NoWait,
		nil, // arguments
	)

	helpers.FailOnError(err, "Falied create queue")

	return &q, ch, conn

}
func (mq *RabbitMQ) formatAmpqURI() string {
	return "amqp://" + mq.User + mq.Password + "@" + mq.Host + ":" + mq.Port

}

// Publish publish to rabbitmq queue
func (mq *RabbitMQ) Publish(payload string) {

	q, ch, conn := mq.declareQueue()

	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		})

	if err == nil {
		conn.Close()
	}

	helpers.FailOnError(err, "Failed to publish a message")
}

// GetConnection get RabbitMQ connection
func (mq *RabbitMQ) GetConnection() *amqp.Connection {
	conn, err := amqp.Dial(mq.formatAmpqURI())
	helpers.FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

// NewRabbitMQ is my IoC
func NewRabbitMQ() *RabbitMQ {
	return GetManagerAppConfig().RabbitMQ
}
