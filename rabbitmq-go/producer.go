package main

import (
	"log"
	"os"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "No se pudo conectar a RabbitMQ")
	defer connection.Close()

	canal, err := connection.Channel()
	failOnError(err, "Error al abrir un canal")
	defer canal.Close()

	q, err := canal.QueueDeclare(
		"IS-910", // name
		true,         // durable
		false,        // Elimine automaticamente al usarla
		false,        // exclusiva
		false,        // timeout
		nil,          // arguments
	)
	failOnError(err, "Error al declarar una cola o queue")

	body := bodyFrom(os.Args)
	err = canal.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Error al publicar un mensaje")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}