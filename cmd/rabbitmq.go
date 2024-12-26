package cmd

import (
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var InitRabbitMQ = &cobra.Command{
	Use:   "rabbitmq",
	Short: "Inicia o serviço RabbitMQ",
	Long:  `Este comando inicia o RabbitMQ em um contêiner Docker.`,
	Run: func(cmd *cobra.Command, args []string) {
		startRabbitMQ()
		logrus.Info("Iniciando o serviço RabbitMQ...")
	},
}

func startRabbitMQ() {
	cmd := exec.Command("docker-compose", "up", "-d")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.WithError(err).Error("Erro ao iniciar o RabbitMQ")
		logrus.Debug(string(output))
		return
	}
	logrus.Info("RabbitMQ iniciado com sucesso.")
	logrus.Debug(string(output))
}
