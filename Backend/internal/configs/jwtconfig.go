package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func JwtPubKey() []byte {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return []byte(os.Getenv("publicKey"))
}

func JwtPrKey() []byte {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return []byte(os.Getenv("privateKey"))
}
