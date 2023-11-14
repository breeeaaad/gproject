package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Jwtconfig() []byte {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return []byte(os.Getenv("SecretKey"))
}
