package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadENV() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
