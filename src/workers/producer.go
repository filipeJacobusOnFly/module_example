package workers

import (
	"encoding/json"
	config "module_example/src/database"
	"module_example/src/http/models"

	"github.com/sirupsen/logrus"

	"github.com/streadway/amqp"
)

func PublishRecord(record models.Record, cfg *config.Config) {
	conn, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		logrus.Fatalf("Erro ao conectar ao RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logrus.Fatalf("Erro ao abrir canal: %v", err)
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
		logrus.Fatalf("Erro ao declarar fila: %v", err)
	}

	body, err := json.Marshal(record)
	if err != nil {
		logrus.Fatalf("Erro ao serializar registro: %v", err)
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
		logrus.Fatalf("Erro ao publicar mensagem: %v", err)
	}

	logrus.Infof("Registro enviado: %v", record)
}
