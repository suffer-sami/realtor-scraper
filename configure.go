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
	requests           map[int]Request
	logger             Logger
	platform           string
	jwtSecret          string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) getRemainingRequests() (remaining []int, isComplete bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	remainingKeys := []int{}

	for key, req := range cfg.requests {
		if !req.processed {
			remainingKeys = append(remainingKeys, key)
		}
	}
	return remainingKeys, len(remainingKeys) == 0
}

func (cfg *config) addRequests(requests []Request) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	for i := range requests {
		request := requests[i]
		cfg.requests[request.offset] = request
	}
}

func (cfg *config) markRequestProcessed(key int) error {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	request, ok := cfg.requests[key]
	if !ok {
		return fmt.Errorf("invalid key: %d", key)
	}
	request.processed = true
	cfg.requests[key] = request
	return nil
}

func (cfg *config) getRequest(key int) (Request, error) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	request, ok := cfg.requests[key]
	if !ok {
		return Request{}, fmt.Errorf("invalid key: %d", key)
	}

	return request, nil
}

func (cfg *config) getRequestCount() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.requests)
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
		client: &http.Client{Timeout: defaultRequestTimeout, Transport: &http.Transport{
			DisableKeepAlives:   true,
			MaxIdleConnsPerHost: defaultMaxConcurrency,
			IdleConnTimeout:     10 * time.Second,
		}},
		db:                 db,
		dbQueries:          *dbQueries,
		requests:           make(map[int]Request),
		logger:             stdDebugLogger{},
		platform:           platform,
		jwtSecret:          jwtSecret,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}
