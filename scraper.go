package main

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/suffer-sami/realtor-scraper/internal/database"
)

const (
	baseUrl               = "https://realtor.com"
	apiEndpoint           = baseUrl + "/realestateagents/api/v3/search"
	userAgent             = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
	defaultResultsPerPage = 20
	tokenTTL              = 1 * time.Minute
)

type Request struct {
	resultsPerPage int
	offset         int
	processed      bool
}

// getRequestParams constructs the search url parameters.
func getRequestParams(offset, resultsPerPage int) SearchRequestParams {
	return SearchRequestParams{
		Offset:              offset,
		Limit:               resultsPerPage,
		MarketingAreaCities: "_",
		Types:               "agent",
		Sort:                "agent_rating_high",
		FarOptOut:           "false",
		ClientID:            "FAR2.0",
		SeoUserType:         SeoUserType{IsBot: "false", DeviceType: "desktop"},
		IsCountySearch:      "false",
	}
}

// processRequest processes a given request and store the fetched agents
func (cfg *config) processRequest(request Request) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	cfg.logger.Infof("FETCHING: Agents (offset %d, limit %d)", request.offset, request.resultsPerPage)

	agents, err := cfg.getAgents(request.offset, request.resultsPerPage)
	if err != nil {
		cfg.logger.Errorf("error getting request (%d, %d): %v", request.offset, request.resultsPerPage, err)
		return
	}

	cfg.markRequestProcessed(request.offset)

	cfg.wg.Add(1)
	go func() {
		cfg.storeAgents(agents)
	}()
}

// getSearchResults fetches search results from the API.
func (cfg *config) getSearchResults(payload SearchRequestParams) (SearchRequestResponse, error) {
	parsedURL, _ := url.Parse(apiEndpoint)
	queryParams, err := buildQueryParams(payload)
	if err != nil {
		return SearchRequestResponse{}, fmt.Errorf("failed to build query params: %w", err)
	}
	parsedURL.RawQuery = queryParams.Encode()

	token, err := generateBearerToken(cfg.jwtSecret)
	if err != nil {
		return SearchRequestResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return SearchRequestResponse{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	setHeaders(req, token)

	resp, err := cfg.client.Do(req)
	if err != nil {
		return SearchRequestResponse{}, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return SearchRequestResponse{}, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearchRequestResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var response SearchRequestResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return SearchRequestResponse{}, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	return response, nil
}

// buildQueryParams converts the search payload into query parameters.
func buildQueryParams(payload SearchRequestParams) (url.Values, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload to JSON: %w", err)
	}

	var payloadMap map[string]interface{}
	if err := json.Unmarshal(payloadJSON, &payloadMap); err != nil {
		return nil, fmt.Errorf("error unmarshalling payload JSON: %w", err)
	}

	queryParams := url.Values{}
	for key, value := range payloadMap {
		queryParams.Add(key, fmt.Sprintf("%v", value))
	}

	return queryParams, nil
}

// setHeaders sets headers for the HTTP request.
func setHeaders(req *http.Request, token string) {
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Origin", baseUrl)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("User-Agent", userAgent)
}

// generateBearerToken creates a signed JWT token.
func generateBearerToken(secret string) (string, error) {
	claims := jwt.MapClaims{
		"exp": jwt.NewNumericDate(time.Now().UTC().Add(tokenTTL)),
		"sub": "find_a_realtor",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// getTotalResults retrieves the total number of matching rows.
func (cfg *config) getTotalResults() (int, error) {
	payload := getRequestParams(0, 0)

	response, err := cfg.getSearchResults(payload)
	if err != nil {
		return 0, fmt.Errorf("error getting search results: %w", err)
	}

	return response.MatchingRows, nil
}

// getAgents retrieves list of normalized agents matching the search criteria.
func (cfg *config) getAgents(offset, resultsPerPage int) ([]Agent, error) {
	payload := getRequestParams(offset, resultsPerPage)

	response, err := cfg.getSearchResults(payload)
	if err != nil {
		return nil, fmt.Errorf("error getting search results: %w", err)
	}

	for i := range response.Agents {
		cfg.normalizeAgent(&response.Agents[i])
	}

	if err := cfg.executeTransaction(context.Background(), func(ctx context.Context, qtx *database.Queries) error {
		_, err := qtx.CreateRequest(ctx, database.CreateRequestParams{
			Offset:         int64(offset),
			ResultsPerPage: int64(resultsPerPage),
		})
		return err
	}); err != nil {
		return nil, fmt.Errorf("error creating dbRequest: %w", err)
	}

	return response.Agents, nil
}

// getRequests calculates all requests needed to fetch `totalResults`, marking previously processed requests as true.
func (cfg *config) getRequests(totalResults int) ([]Request, error) {
	totalPages := int(math.Ceil(float64(totalResults) / float64(defaultResultsPerPage)))
	requestMap := initializeRequestMap(totalPages)

	prevRequests, err := cfg.fetchPreviousRequests()
	if err != nil {
		return nil, fmt.Errorf("error fetching previous dbRequest: %w", err)
	}

	markProcessedRequests(requestMap, prevRequests)
	return sortRequests(requestMap), nil
}

// initializeRequestMap creates an initial map of all potential requests.
func initializeRequestMap(totalPages int) map[int]Request {

	requestMap := make(map[int]Request, totalPages)
	for page := 0; page < totalPages; page++ {
		offset := page * defaultResultsPerPage
		requestMap[offset] = Request{
			resultsPerPage: defaultResultsPerPage,
			offset:         offset,
			processed:      false,
		}
	}
	return requestMap
}

// fetchPreviousRequests retrieves already processed requests from db.
func (cfg *config) fetchPreviousRequests() ([]database.GetRequestsRow, error) {
	previousRequests, err := cfg.dbQueries.GetRequests(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get requests from db: %w", err)
	}
	return previousRequests, nil
}

// markProcessedRequests updates the request map with processed request data.
func markProcessedRequests(requestMap map[int]Request, prevRequests []database.GetRequestsRow) {
	for _, prevReq := range prevRequests {
		offset := int(prevReq.Offset)
		requestMap[offset] = Request{
			resultsPerPage: int(prevReq.ResultsPerPage),
			offset:         offset,
			processed:      true,
		}
	}
}

// sortRequests converts the request map into a sorted slice of Request objects.
func sortRequests(requestMap map[int]Request) []Request {
	sortedKeys := make([]Request, 0, len(requestMap))
	for _, req := range requestMap {
		sortedKeys = append(sortedKeys, req)
	}

	sort.Slice(sortedKeys, func(i, j int) bool {
		return sortedKeys[i].offset < sortedKeys[j].offset
	})
	return sortedKeys
}
