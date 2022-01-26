package db

import (
	"car_informer/internal/app/model"
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

func InsertUser(user model.User) int64 {
	db := createConnection()
	defer db.Close()

	sqlQuery := `INSERT INTO users (ID, Email, Password, EncryptedPassword) VALUES($1, $2, $3, $4) RETURNING id`

	var id int64

	if err := db.QueryRow(sqlQuery, user.ID, user.Email, user.Password, user.EncryptedPassword).Scan(&id); err != nil {

	}

	return id
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
