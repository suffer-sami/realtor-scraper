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

	for {
		remainingReqs, isComplete := cfg.getRemainingRequests()
		cfg.logger.Infof("STATS: (Total Agents: %d, Remaining Agents: %d)", cfg.getRequestCount(), len(remainingReqs))
		if isComplete {
			cfg.logger.Infof("========== COMPLETE ==========")
			return
		}
		count := 0

		for _, reqKey := range remainingReqs {
			req, err := cfg.getRequest(reqKey)
			if err != nil {
				cfg.logger.Fatalf("error getting remaining requests (key: %d): %v", reqKey, err)
			}
			if req.processed {
				continue
			}
			request, err := cfg.getRequest(req.offset)
			if err != nil {
				cfg.logger.Fatalf("error: %v", err)
			}

			cfg.wg.Add(1)
			go cfg.processRequest(request)

			count++
			if count%cfg.throttleRequestLimit == 0 {
				cfg.logger.Infof(
					"COOLDOWN: trottling requests For %v after %d requests",
					defaultThrottleTimeout,
					cfg.throttleRequestLimit,
				)
				cfg.client.CloseIdleConnections()
				time.Sleep(defaultThrottleTimeout)
			}

			if cfg.platform == "dev" && count >= cfg.throttleRequestLimit {
				break
			}

		}
		cfg.wg.Wait()
		if cfg.platform == "dev" {
			break
		}
	}
}
