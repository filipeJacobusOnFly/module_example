package middleware

import (
	"log"
	mainDB "module_example/src/database"
	cache "module_example/src/http/cache"
	"module_example/src/http/controllers"
	repositories "module_example/src/http/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	cacheInstance *cache.TokenCache
)

func InitGin() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	db, _ := mainDB.GetCon()
	cacheInstance = cache.NewTokenCache()
	dbWrapper := &repositories.DBWrapper{DB: db}
	tokenRepo := repositories.NewTokenRepository(dbWrapper, cacheInstance)
	recordRepo := repositories.NewRecordRepository(db)

	r := gin.Default()

	go repositories.StartBatchProcessing(recordRepo)
	r.Use(controllers.AuthMiddleware(tokenRepo))

	r.GET("/pdf", controllers.PdfHandler)
	r.POST("/records", controllers.RecordHandler(recordRepo))

	if err := r.Run(":9051"); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
