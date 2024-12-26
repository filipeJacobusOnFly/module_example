package workers

import (
	"encoding/json"
	"log"
	config "module_example/src/database"
	"module_example/src/http/models"
	repositories "module_example/src/http/repository"

	"github.com/streadway/amqp"
)

func ConsumeRecords(repo *repositories.RecordRepository, cfg *config.Config) {
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
		log.Fatalf("Erro ao consumir mensagens: %s", err)
	}

	log.Println("Aguardando registros. Para sair pressione CTRL+C")

	for msg := range msgs {
		var record models.Record
		if err := json.Unmarshal(msg.Body, &record); err != nil {
			log.Printf("Erro ao deserializar registro: %s", err)
			continue
		}

		if err := repo.CreateRecords([]models.Record{record}); err != nil {
			log.Printf("Erro ao armazenar registro no banco de dados: %s", err)
		} else {
			log.Printf("Registro armazenado: %v", record)
		}
	}
}
