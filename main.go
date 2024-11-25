package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	jwtSecret string
	client    *http.Client
}

func main() {
	godotenv.Load()
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set")
	}

	client := &http.Client{}

	cfg := &Config{
		jwtSecret: jwtSecret,
		client:    client,
	}
}
