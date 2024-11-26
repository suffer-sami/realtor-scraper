package main

import (
	"fmt"
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

	cfg := &Config{
		jwtSecret: jwtSecret,
		client:    &http.Client{Timeout: defaultRequestTimeout},
	}

	totalResults, err := cfg.getTotalResults()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(totalResults)
}
