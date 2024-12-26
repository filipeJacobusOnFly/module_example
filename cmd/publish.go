package cmd

import (
	"module_example/src/database"
	"module_example/src/http/models"
	"module_example/src/workers"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var PublishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publica 10.000 registros no RabbitMQ",
	Long:  `Este comando gera e publica 10.000 registros na fila do RabbitMQ.`,
	Run: func(cmd *cobra.Command, args []string) {
		publishRecords()
		logrus.Info("Publicando 10.000 registros...")
	},
}

func publishRecords() {
	cfg := database.LoadConfig()

	for i := 0; i < 10000; i++ {
		record := models.Record{
			RecordID: uint(i),
			Date:     time.Now(),
		}
		workers.PublishRecord(record, cfg)
		logrus.Debugf("Registro publicado: %d", i)
	}
}
