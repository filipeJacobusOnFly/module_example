package cmd

import (
	"module_example/src/database"
	repositories "module_example/src/http/repository"
	"module_example/src/workers"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ConsumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Inicia o consumidor RabbitMQ",
	Long:  `Este comando inicia o consumidor RabbitMQ que escuta mensagens na fila.`,
	Run: func(cmd *cobra.Command, args []string) {
		startRabbitMQConsumer()
		logrus.Info("Iniciando o consumidor RabbitMQ...")
	},
}

func startRabbitMQConsumer() {
	db, err := database.GetCon()
	if err != nil {
		logrus.WithError(err).Fatal("Erro ao conectar ao banco de dados")
	}
	defer db.Close()
	cfg := database.LoadConfig()

	repo := repositories.NewRecordRepository(db)

	workers.ConsumeRecords(repo, cfg)
}
