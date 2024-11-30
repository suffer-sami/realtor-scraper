package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/suffer-sami/realtor-scraper/internal/database"
)

const (
	defaultMaxConcurrency = 3
	defaultRequestTimeout = 30 * time.Second
)

type config struct {
	client             *http.Client
	db                 *sql.DB
	dbQueries          database.Queries
	logger             Logger
	platform           string
	jwtSecret          string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func configure(args []string) (*config, error) {
	godotenv.Load()
	maxConcurrency := defaultMaxConcurrency
	if len(args) > 0 {
		if newMaxConcurrency, err := strconv.Atoi(args[0]); err == nil {
			maxConcurrency = newMaxConcurrency
		} else {
			return nil, fmt.Errorf("invalid maxConcurrency: %v", err)
		}
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL must be set")
	}
	dbAuthToken := os.Getenv("DB_AUTH_TOKEN")
	if dbAuthToken == "" {
		return nil, fmt.Errorf("DB_AUTH_TOKEN must be set")
	}
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		return nil, fmt.Errorf("DB_FILE must be set")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		return nil, fmt.Errorf("PLATFORM must be set")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET must be set")
	}

	dbPath := dbFile
	if platform == "prod" {
		dbPath = fmt.Sprintf("%s?authToken=%s", dbURL, dbAuthToken)
	}

	db, err := sql.Open("libsql", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open db %s", err)
	}

	dbQueries := database.New(db)

	return &config{
		client:             &http.Client{Timeout: defaultRequestTimeout},
		db:                 db,
		dbQueries:          *dbQueries,
		logger:             stdDebugLogger{},
		platform:           platform,
		jwtSecret:          jwtSecret,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}
