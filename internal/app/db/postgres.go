package db

import (
	"crypto/sha256"
	"fmt"
)

const (
	salt = "3DKJH^&%&^DRjhKSFD^$%RSFHG"
)

//func createConnection() *sql.DB {
//
//}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
