package db

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func createConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connPQ := os.Getenv("POSTGRES_URL")

	db, err := sql.Open("postgres", connPQ)

	if err != nil {
		log.Fatalf("could not open connection: %v", err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("could not ping connection: %v", err)
	}

	fmt.Println("successfully established connection!")

	return db
}

func generatePasswordHash(password string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	salt := os.Getenv("SALT_PHRASE")
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
