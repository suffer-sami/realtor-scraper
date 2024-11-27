package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
)

const (
	baseUrl               = "https://realtor.com"
	apiEndpoint           = baseUrl + "/realestateagents/api/v3/search"
	userAgent             = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
	defaultResultsPerPage = 20
	defaultRequestTimeout = 30 * time.Second
	tokenTTL              = 1 * time.Minute
)

// getRequestParams constructs the search url parameters.
func getRequestParams(offset, limit int) SearchRequestParams {
	return SearchRequestParams{
		Offset:              offset,
		Limit:               limit,
		MarketingAreaCities: "_",
		Types:               "agent",
		Sort:                "agent_rating_high",
		FarOptOut:           "false",
		ClientID:            "FAR2.0",
		SeoUserType:         SeoUserType{IsBot: "false", DeviceType: "desktop"},
		IsCountySearch:      "false",
	}
}

// getSearchResults fetches search results from the API.
func (cfg *Config) getSearchResults(payload SearchRequestParams) (SearchRequestResponse, error) {
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
func (cfg *Config) getTotalResults() (int, error) {
	payload := getRequestParams(0, 0)

	response, err := cfg.getSearchResults(payload)
	if err != nil {
		return 0, err
	}

	return response.MatchingRows, nil
}

// getAgents retrieves list of agents matching the search criteria.
func (cfg *Config) getAgents(offset, limit int) ([]Agent, error) {
	payload := getRequestParams(offset, limit)

	response, err := cfg.getSearchResults(payload)
	if err != nil {
		return nil, err
	}

	for i := range response.Agents {
		normalizeAgent(&response.Agents[i])
	}

	return response.Agents, nil
}
