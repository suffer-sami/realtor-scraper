package main

import (
	"os"
	"time"

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
	cfg.addRequests(allRequests)

	count := 0
	for i := range allRequests {
		req := &allRequests[i]
		if req.processed {
			continue
		}
		request, err := cfg.getRequest(req.offset)
		if err != nil {
			cfg.logger.Fatalf("error: %v", err)
		}

		cfg.wg.Add(1)
		go cfg.processRequest(request)
		time.Sleep(1 * time.Second)

		count++
		if count >= 5 {
			break
		}
	}
	cfg.wg.Wait()
}
