package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	config "module_example/src/database"
	gin "module_example/src/http/middleware"
	"module_example/src/http/models"
	repositories "module_example/src/http/repository"
	"module_example/src/workers"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "init",
	Short: "Iniciar comandos",
	Long:  `Inicia comandos.`,
}

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Inicia o servidor HTTP e o RabbitMQ",
	Long:  `Este comando inicia o servidor HTTP.`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
		fmt.Println("Iniciando o servidor HTTP...")
	},
}

var InitRabbitMQ = &cobra.Command{
	Use:   "rabbitmq",
	Short: "Inicia o serviço RabbitMQ",
	Long:  `Este comando inicia o RabbitMQ em um contêiner Docker.`,
	Run: func(cmd *cobra.Command, args []string) {
		startRabbitMQ()
		fmt.Println("Iniciando o serviço RabbitMQ...")
	},
}

var ConsumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Inicia o consumidor RabbitMQ",
	Long:  `Este comando inicia o consumidor RabbitMQ que escuta mensagens na fila.`,
	Run: func(cmd *cobra.Command, args []string) {
		startRabbitMQConsumer()
		fmt.Println("Iniciando o consumidor RabbitMQ...")
	},
}

var PublishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publica 10.000 registros no RabbitMQ",
	Long:  `Este comando gera e publica 10.000 registros na fila do RabbitMQ.`,
	Run: func(cmd *cobra.Command, args []string) {
		publishRecords()
		fmt.Println("Publicando 10.000 registros...")
	},
}

func publishRecords() {
	cfg := config.LoadConfig()

	for i := 0; i < 10000; i++ {
		record := models.Record{
			RecordID: uint(i),
			Date:     time.Now(),
		}
		workers.PublishRecord(record, cfg)
	}
}

func init() {
	RootCmd.AddCommand(ServeCmd)
	RootCmd.AddCommand(InitRabbitMQ)
	RootCmd.AddCommand(ConsumeCmd)
	RootCmd.AddCommand(PublishCmd)
}

func startRabbitMQ() {
	cmd := exec.Command("docker-compose", "up", "-d")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao obter logs:", err)
		fmt.Println(string(output))
		return
	}
	fmt.Println(string(output))
}

func startRabbitMQConsumer() {
	db, err := config.GetCon()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %s", err)
	}
	defer db.Close()
	cfg := config.LoadConfig()

	repo := repositories.NewRecordRepository(db)

	workers.ConsumeRecords(repo, cfg)
}

func startServer() {

	gin.InitGin()
}
