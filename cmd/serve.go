package cmd

import (
	"module_example/src/http/middleware"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Inicia o servidor HTTP e o RabbitMQ",
	Long:  `Este comando inicia o servidor HTTP.`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
		logrus.Info("Iniciando o servidor HTTP...")
	},
}

func startServer() {
	middleware.InitGin()
	logrus.Info("Servidor HTTP iniciado.")
}
