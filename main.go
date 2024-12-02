package main

import (
	"os"

	_ "github.com/tursodatabase/go-libsql"
)

func main() {
	args := os.Args[1:]

	cfg, err := configure(args)
	if err != nil {
		cfg.logger.Fatalf("error while configuration: %v", err)
	}
	defer cfg.db.Close()

	totalResults, err := cfg.getTotalResults()
	if err != nil {
		cfg.logger.Fatalf("error getting total results: %v", err)
	}
	cfg.logger.Infof("Total Agents: %d\n", totalResults)

	allRequests, err := cfg.getRequests(totalResults)
	if err != nil {
		cfg.logger.Fatalf("error getting search requests: %v", err)
	}

	for i := range allRequests {
		request := &allRequests[i]
		if request.processed {
			continue
		}

		cfg.logger.Infof("Getting Agents (Offset: %d, ResultsPerPage: %d)...\n", request.offset, request.resultsPerPage)
		agents, err := cfg.getAgents(request.offset, request.resultsPerPage)
		if err != nil {
			cfg.logger.Fatalf("error getting request (%d, %d): %v", request.offset, request.resultsPerPage, err)
		}

		for _, agent := range agents {
			cfg.wg.Add(1)
			go func() {
				if err := cfg.storeAgent(agent); err != nil {
					cfg.logger.Errorf("error storing agent (ID: %s): %v", agent.ID, err)
				}
			}()
		}
		break
	}
	cfg.wg.Wait()
}
