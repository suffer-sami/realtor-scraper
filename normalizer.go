package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/purell"
	"github.com/nyaruka/phonenumbers"
)

const (
	defaultCountryCode = "US"
)

// normalizeAgent normalizes the fields of an Agent struct.
func (cfg *config) normalizeAgent(agent *Agent) {

	if normalizedAgentHref, err := tryNormalizeURL(agent.Href); err != nil {
		cfg.logger.Warnf("error while normalizing agent href URL: %v", err)
	} else {
		agent.Href = normalizedAgentHref
	}

	if normalizedOfficeWebsite, err := tryNormalizeURL(agent.Office.Website); err != nil {
		cfg.logger.Warnf("error while normalizing agent office website URL: %v", err)
	} else {
		agent.Office.Website = normalizedOfficeWebsite
	}

	agent.Email = normalizeEmail(agent.Email)
	agent.Office.Email = normalizeEmail(agent.Office.Email)

	agentCountry := getCountryCode(agent.Address.Country, defaultCountryCode)
	officeCountry := getCountryCode(agent.Office.Address.Country, agentCountry)

	for i := range agent.Phones {
		if err := normalizePhone(&agent.Phones[i], agentCountry); err != nil {
			cfg.logger.Warnf("error while normalizing agent phone: %v", err)
		}
	}

	for i := range agent.Office.Phones {
		if err := normalizePhone(&agent.Office.Phones[i], officeCountry); err != nil {
			cfg.logger.Warnf("error while normalizing agent office phone: %v", err)
		}
	}

	for key, phone := range agent.Office.PhoneList {
		if err := normalizePhone(&phone, officeCountry); err != nil {
			cfg.logger.Warnf("error while normalizing agent office phone: %v", err)
			continue
		}
		agent.Office.PhoneList[key] = phone
	}

	for k, socialMedia := range agent.SocialMedias {
		if normalizedSocialMediaHref, err := tryNormalizeURL(socialMedia.Href); err != nil {
			cfg.logger.Warnf("error while normalizing agent social media URL: %v", err)
		} else {
			agent.SocialMedias[k] = SocialMedia{
				Href: normalizedSocialMediaHref,
				Type: strings.ToLower(socialMedia.Type),
			}
		}
	}
}

// normalizeEmail cleans and normalizes email.
func normalizeEmail(email string) string {
	if email == "" {
		return ""
	}
	return strings.ToLower(strings.TrimSpace(email))
}

// normalizePhone normalizes a phone number to the international format.
func normalizePhone(phone *Phone, regionCode string) error {
	phone.IsValid = false
	if phone.Number == "" {
		return nil
	}
	parsedNumber, err := phonenumbers.Parse(phone.Number, regionCode)
	if err != nil {
		return fmt.Errorf("failed to parse phone number '%s': %w", phone.Number, err)
	}

	if phonenumbers.IsValidNumber(parsedNumber) {
		phone.IsValid = true
	}
	phone.Number = phonenumbers.Format(parsedNumber, phonenumbers.INTERNATIONAL)
	return nil
}

// tryNormalizeURL attempts to normalize a URL and logs any errors.
func tryNormalizeURL(rawURL string) (string, error) {
	if rawURL == "" {
		return "", nil
	}
	normalized, err := normalizeURL(rawURL)
	if err != nil {
		return "", err
	}
	return normalized, nil
}

// normalizeURL cleans and normalizes a URL string.
func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(strings.TrimSpace(strings.ToLower(rawURL)))
	if err != nil {
		return "", fmt.Errorf("failed to parse URL '%s': %w", rawURL, err)
	}

	// Simplify URL if it has duplicate hostnames
	// e.g. https://twitter.com/http://twitter.com/username, http://www.facebook.com/http://facebook.com/username etc

	if strings.Contains(parsedURL.Path, strings.TrimPrefix(parsedURL.Host, "www.")) || strings.Contains(parsedURL.Path, ".") {
		parsedPathURL, err := url.Parse(strings.TrimPrefix(parsedURL.Path, "/"))
		if err != nil {
			return "", fmt.Errorf("failed to parse sub-path URL '%s': %w", parsedPathURL, err)
		}
		parsedURL.Path = parsedPathURL.Path
	}

	return purell.NormalizeURL(parsedURL, purell.FlagsUsuallySafeGreedy), nil
}

// getCountryCode extracts a 2-letter country code or returns a default value.
func getCountryCode(country string, defaultCode string) string {
	if len(country) >= 2 {
		country_code := strings.ToUpper(country[:2])
		if _, ok := phonenumbers.GetSupportedRegions()[country_code]; ok {
			return country_code
		}
	}
	return defaultCode
}
