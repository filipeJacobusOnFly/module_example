# Module Example

Este é um exemplo de projeto em Go que utiliza RabbitMQ para publicar e consumir registros. O projeto inclui um servidor HTTP que pode manipular requisições para gerar e gerenciar registros.

## Estrutura do Projeto
```
module_example/
├── cmd
│   ├── consume.go
│   ├── publish.go
│   ├── rabbitmq.go
│   ├── root.go
│   └── serve.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── logs
│   └── app.log
├── main.go
├── README.md
├── src
│   ├── database
│   │   ├── conection.go
│   │   └── rabbitMQConection.go
│   ├── http
│   │   ├── cache
│   │   │   └── cache.go
│   │   ├── controllers
│   │   │   ├── authController.go
│   │   │   ├── pdfConvertController.go
│   │   │   └── recordController.go
│   │   ├── middleware
│   │   │   └── endpoints.go
│   │   ├── models
│   │   │   ├── authModel.go
│   │   │   └── recordModel.go
│   │   └── repository
│   │       ├── authRepository.go
│   │       ├── queueRepository.go
│   │       └── recordsRepository.go
│   ├── logger
│   │   └── logger.go
│   └── workers
│       ├── consumer.go
│       └── producer.go
└── tests
└── unit
├── authController_test.go
├── authRepository_test.go
├── pdfConvertController_test.go
├── queueRepository_test.go
├── recordController_test.go
└── recordsRepository_test.go
```

## Pré-requisitos

- Go (versão 1.16 ou superior)
- Docker e Docker Compose (para executar o RabbitMQ)

## Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu_usuario/module_example.git
   cd module_example

2. Instale as dependências:
    ```bash
     go mod tidy

2. Inicie o RabbitMQ usando Docker:
    ```bash
    docker-compose up -d

# Uso
## Iniciar o Servidor HTTP

Para iniciar o servidor HTTP, execute o seguinte comando:

    docker-compose up -d

## Publicar Registros no RabbitMQ

Para publicar 10.000 registros no RabbitMQ, execute o seguinte comando:
    
    go run main.go publish

## Consumir Registros do RabbitMQ

Para iniciar o consumidor RabbitMQ, execute o seguinte comando:

    go run main.go consume

# Endpoints

Endpoints

* `GET /pdf`: Manipula requisições para gerar PDFs.
* `POST /records`: Recebe registros e os processa.

# Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

# Licença

Este projeto está licenciado sob a MIT License - veja o arquivo LICENSE para mais detalhes.