package middleware

import (
	mainDB "module_example/src/database"
	cache "module_example/src/http/cache"
	"module_example/src/http/controllers"
	repositories "module_example/src/http/repository"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	cacheInstance *cache.TokenCache
)

func InitGin() {
	// Configurando o logger
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Erro ao carregar o arquivo .env")
	}

	db, err := mainDB.GetCon()
	if err != nil {
		logrus.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	cacheInstance = cache.NewTokenCache()
	dbWrapper := &repositories.DBWrapper{DB: db}
	tokenRepo := repositories.NewTokenRepository(dbWrapper, cacheInstance)
	recordRepo := repositories.NewRecordRepository(db)

	r := gin.Default()

	go repositories.StartBatchProcessing(recordRepo)
	r.Use(controllers.AuthMiddleware(tokenRepo))

	r.GET("/pdf", controllers.PdfHandler)
	r.POST("/records", controllers.RecordHandler(recordRepo))

	logrus.Info("Iniciando o servidor na porta 9051")
	if err := r.Run(":9051"); err != nil {
		logrus.Fatal("Erro ao iniciar o servidor:", err)
	}
}
