package database

import (
	"database/sql"
	"log"

	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetCon() (*sql.DB, error) {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %s", err)
	}
	once.Do(func() {
		dsn := "host=" + os.Getenv("DB_HOST") +
			" port=" + os.Getenv("DB_PORT") +
			" user=" + os.Getenv("DB_USER") +
			" password=" + os.Getenv("DB_PASSWORD") +
			" dbname=" + os.Getenv("DB_NAME") +
			" sslmode=disable"

		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Println("Erro ao conectar ao banco de dados:", err)
			return
		}
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(0)

		if err = db.Ping(); err != nil {
			log.Println("Erro ao pingar o banco de dados:", err)
		}
	})

	return db, err
}
