package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/joho/godotenv"
	"github.com/suffer-sami/realtor-scraper/internal/database"
	"github.com/suffer-sami/realtor-scraper/internal/logger"
)

const (
	defaultMaxConcurrency       = 3
	defaultRequestTimeout       = 30 * time.Second
	defaultThrottleRequestLimit = 5
	defaultThrottleTimeout      = 5 * time.Second
	defaultUseDbLocal           = true
	defaultSaveRawAgents        = false
	defaultLoggerPrefix         = "realtor-scraper"
	defaultLogLevel             = "INFO"
)

type config struct {
	client               *http.Client
	db                   *sql.DB
	dbQueries            database.Queries
	requests             map[int]Request
	saveRawAgents        bool
	throttleRequestLimit int
	logger               logger.Logger
	platform             string
	jwtSecret            string
	mu                   *sync.Mutex
	agentCount           atomic.Int32
	activeRequestCount   atomic.Int32
	maxThreadCount       int
	concurrencyControl   chan struct{}
	wg                   *sync.WaitGroup
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
	sort.Ints(remainingKeys)
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
	// Check if godotenv.Load() returns an error and handle it
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}
	maxConcurrency := defaultMaxConcurrency
	if len(args) > 0 {
		if newMaxConcurrency, err := strconv.Atoi(args[0]); err == nil {
			maxConcurrency = newMaxConcurrency
		} else {
			return nil, fmt.Errorf("invalid maxConcurrency: %w", err)
		}
	}

	platform := os.Getenv("PLATFORM")
	if platform == "" {
		return nil, fmt.Errorf("PLATFORM must be set")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET must be set")
	}

	saveRawAgents := defaultSaveRawAgents
	if saveRawAgentsStr := os.Getenv("SAVE_RAW_AGENTS"); saveRawAgentsStr != "" {
		if newsaveRawAgents, err := strconv.ParseBool(saveRawAgentsStr); err == nil {
			saveRawAgents = newsaveRawAgents
		}
	}
	logLevel := defaultLogLevel
	newLogLevel := os.Getenv("LOG_LEVEL")
	if newLogLevel != "" {
		logLevel = newLogLevel
	}
	throttleRequestLimit := defaultThrottleRequestLimit
	if newThrottleReqLimitStr := os.Getenv("THROTTLE_REQUEST_LIMIT"); newThrottleReqLimitStr != "" {
		newLimit, err := strconv.Atoi(newThrottleReqLimitStr)
		if err == nil && newLimit > 0 {
			throttleRequestLimit = newLimit
		} else {
			return nil, fmt.Errorf("invalid THROTTLE_REQUEST_LIMIT: %w", err)
		}
	}

	useDbLocal := defaultUseDbLocal
	if useDbLocalStr := os.Getenv("USE_DB_LOCAL"); useDbLocalStr != "" {
		if newUseDbLocalStr, err := strconv.ParseBool(useDbLocalStr); err == nil {
			useDbLocal = newUseDbLocalStr
		}
	}

	var dbPath string
	if useDbLocal {
		dbFile := os.Getenv("DB_FILE")
		if dbFile == "" {
			return nil, fmt.Errorf("DB_FILE must be set")
		}
		dbPath = dbFile
	} else {
		dbURL := os.Getenv("DB_URL")
		if dbURL == "" {
			return nil, fmt.Errorf("DB_URL must be set")
		}
		dbAuthToken := os.Getenv("DB_AUTH_TOKEN")
		if dbAuthToken == "" {
			return nil, fmt.Errorf("DB_AUTH_TOKEN must be set")
		}

		if !strings.HasPrefix(dbURL, "libsql") {
			return nil, fmt.Errorf("invalid database URL format")
		}
		dbPath = fmt.Sprintf("%s?authToken=%s", dbURL, dbAuthToken)
	}

	db, err := sql.Open("libsql", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	dbQueries := database.New(db)

	return &config{
		client: &http.Client{Timeout: defaultRequestTimeout, Transport: &http.Transport{
			MaxIdleConnsPerHost: defaultMaxConcurrency,
			IdleConnTimeout:     10 * time.Second,
		}},
		db:                   db,
		dbQueries:            *dbQueries,
		requests:             make(map[int]Request),
		throttleRequestLimit: throttleRequestLimit,
		saveRawAgents:        saveRawAgents,
		logger:               logger.NewLogger(defaultLoggerPrefix, logLevel),
		platform:             platform,
		jwtSecret:            jwtSecret,
		mu:                   &sync.Mutex{},
		maxThreadCount:       maxConcurrency,
		concurrencyControl:   make(chan struct{}, maxConcurrency),
		wg:                   &sync.WaitGroup{},
	}, nil
}
