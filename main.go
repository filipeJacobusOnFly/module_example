package main

import (
	"database/sql"
	"log"
	"module_example/cache"
	"module_example/controllers"
	"module_example/repositories"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	once          sync.Once
	db            *sql.DB
	err           error
	cacheInstance *cache.TokenCache
)

func GetCon() *sql.DB {
	once.Do(func() {
		dsn := "host=" + os.Getenv("DB_HOST") +
			" port=" + os.Getenv("DB_PORT") +
			" user=" + os.Getenv("DB_USER") +
			" password=" + os.Getenv("DB_PASSWORD") +
			" dbname=" + os.Getenv("DB_NAME") +
			" sslmode=disable"

		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Fatal("Erro ao conectar ao banco de dados:", err)
		}
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(0)
		cacheInstance = cache.NewTokenCache()
		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}
	})
	return db
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	db := GetCon()

	tokenRepo := repositories.NewTokenRepository(db, cacheInstance)
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
