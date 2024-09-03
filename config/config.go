package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if  err := godotenv.Load(); err != nil {
		log.Fatalf("Could not load .env file : %v", err)
	}
}