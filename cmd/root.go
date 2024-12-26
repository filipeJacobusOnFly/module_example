package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "init",
	Short: "Iniciar comandos",
	Long:  `Inicia comandos.`,
}

func init() {
	RootCmd.AddCommand(ServeCmd)
	RootCmd.AddCommand(InitRabbitMQ)
	RootCmd.AddCommand(ConsumeCmd)
	RootCmd.AddCommand(PublishCmd)
}
