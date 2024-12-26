package workers

import (
	"encoding/json"
	"log"

	config "module_example/src/database"
	"module_example/src/http/models"

	"github.com/streadway/amqp"
)

func PublishRecord(record models.Record, cfg *config.Config) {
	conn, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("Erro ao conectar ao RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Erro ao abrir canal: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"records_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Erro ao declarar fila: %s", err)
	}

	body, err := json.Marshal(record)
	if err != nil {
		log.Fatalf("Erro ao serializar registro: %s", err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Fatalf("Erro ao publicar mensagem: %s", err)
	}

	log.Printf("Registro enviado: %v", record)
}
