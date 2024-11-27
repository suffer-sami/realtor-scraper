package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/suffer-sami/realtor-scraper/internal/database"
	_ "github.com/tursodatabase/go-libsql"
)

type Config struct {
	client    *http.Client
	db        database.Queries
	platform  string
	jwtSecret string
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	dbAuthToken := os.Getenv("DB_AUTH_TOKEN")
	if dbAuthToken == "" {
		log.Fatal("DB_AUTH_TOKEN must be set")
	}
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		log.Fatal("DB_FILE must be set")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set")
	}

	dbPath := dbFile
	if platform != "dev" {
		dbPath = fmt.Sprintf("%s?authToken=%s", dbURL, dbAuthToken)
	}

	db, err := sql.Open("libsql", dbPath)
	if err != nil {
		log.Fatalf("failed to open db %s", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	cfg := &Config{
		client:    &http.Client{Timeout: defaultRequestTimeout},
		db:        *dbQueries,
		platform:  platform,
		jwtSecret: jwtSecret,
	}

	totalResults, err := cfg.getTotalResults()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total Agents: %d\n", totalResults)

	allRequests, err := cfg.getRequests(totalResults)
	if err != nil {
		log.Fatal(err)
	}

	count := 0

	for i := range allRequests {
		request := &allRequests[i]
		if request.processed {
			continue
		}

		log.Printf("Getting Agents (Offset: %d, ResultsPerPage: %d)...\n", request.offset, request.resultsPerPage)
		agents, err := cfg.getAgents(request.offset, request.resultsPerPage)
		if err != nil {
			log.Fatal(err)
		}
		request.processed = true

		for i, agent := range agents {
			fmt.Printf("%d. Agent: %s\n", i+1, agent.FullName)
			dbAgent := database.CreateAgentParams{
				ID:                   agent.ID,
				FirstName:            toNullString(agent.FirstName),
				LastName:             toNullString(agent.LastName),
				NickName:             toNullString(agent.NickName),
				PersonName:           toNullString(agent.PersonName),
				Title:                toNullString(agent.Title),
				Slogan:               toNullString(agent.Slogan),
				Email:                toNullString(agent.Email),
				AgentRating:          toNullInt(agent.AgentRating),
				Description:          toNullString(agent.Description),
				RecommendationsCount: toNullInt(agent.RecommendationsCount),
				ReviewCount:          toNullInt(agent.ReviewCount),
				LastUpdated:          strToNullTime(agent.LastUpdated, time.RFC1123),
				FirstMonth:           toNullInt(int(agent.FirstMonth)),
				FirstYear:            toNullInt(int(agent.AgentRating)),
				Video:                toNullString(agent.Video),
				WebUrl:               toNullString(agent.WebURL),
				Href:                 toNullString(agent.Href),
			}
			_, err = cfg.db.CreateAgent(context.Background(), dbAgent)
			if err != nil {
				log.Print(err)
				continue
			}
		}

		count++
		if count >= 5 {
			break
		}
	}
}
