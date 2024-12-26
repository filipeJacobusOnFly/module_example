package workers

import (
	"encoding/json"
	config "module_example/src/database"
	"module_example/src/http/models"
	repositories "module_example/src/http/repository"

	"github.com/sirupsen/logrus"

	"github.com/streadway/amqp"
)

func ConsumeRecords(repo *repositories.RecordRepository, cfg *config.Config) {
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logrus.Fatalf("Erro ao consumir mensagens: %v", err)
	}

	logrus.Info("Aguardando registros. Para sair pressione CTRL+C")

	for msg := range msgs {
		var record models.Record
		if err := json.Unmarshal(msg.Body, &record); err != nil {
			logrus.Warnf("Erro ao deserializar registro: %v", err)
			continue
		}

		if err := repo.CreateRecords([]models.Record{record}); err != nil {
			logrus.Errorf("Erro ao armazenar registro no banco de dados: %v", err)
		} else {
			logrus.Infof("Registro armazenado: %v", record)
		}
	}
}
