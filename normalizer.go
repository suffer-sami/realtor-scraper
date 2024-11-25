package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/purell"
	"github.com/nyaruka/phonenumbers"
)

const (
	defaultCountryCode = "US"
)

// normalizeAgent normalizes the fields of an Agent struct.
func normalizeAgent(agent *Agent) {
	agent.Href = tryNormalizeURL(agent.Href)
	agent.Office.Website = tryNormalizeURL(agent.Office.Website)

	agentCountry := getCountryCode(agent.Address.Country, defaultCountryCode)
	officeCountry := getCountryCode(agent.Office.Address.Country, agentCountry)

	normalizePhoneList(agent.Phones, agentCountry)
	normalizePhoneList(agent.Office.Phones, officeCountry)

	for key, phone := range agent.Office.PhoneList {
		normalizePhone(&phone, officeCountry)
		agent.Office.PhoneList[key] = phone
	}

	for k, socialMedia := range agent.SocialMedias {
		agent.SocialMedias[k] = SocialMedia{
			Href: tryNormalizeURL(socialMedia.Href),
			Type: socialMedia.Type,
		}
	}
}

// normalizePhoneList normalizes all phone numbers in a list.
func normalizePhoneList(phones []Phone, regionCode string) {
	for i := range phones {
		normalizePhone(&phones[i], regionCode)
	}
}

// normalizePhone normalizes a phone number to the international format.
func normalizePhone(phone *Phone, regionCode string) {
	parsedNumber, err := phonenumbers.Parse(phone.Number, regionCode)
	if err != nil {
		log.Printf("Failed to parse phone number '%s': %v", phone.Number, err)
		return
	}
	phone.Number = phonenumbers.Format(parsedNumber, phonenumbers.INTERNATIONAL)
}

// tryNormalizeURL attempts to normalize a URL and logs any errors.
func tryNormalizeURL(rawURL string) string {
	if rawURL == "" {
		return ""
	}
	normalized, err := normalizeURL(rawURL)
	if err != nil {
		log.Printf("Failed to normalize URL '%s': %v", rawURL, err)
		return rawURL
	}
	return normalized
}

// normalizeURL cleans and normalizes a URL string.
func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(strings.TrimSpace(strings.ToLower(rawURL)))
	if err != nil {
		return "", err
	}

	// Simplify URL if it has duplicate hostnames
	// e.g. https://twitter.com/http://twitter.com/username, http://www.facebook.com/http://facebook.com/username etc

	if strings.Contains(parsedURL.Path, strings.TrimPrefix(parsedURL.Host, "www.")) || strings.Contains(parsedURL.Path, ".") {
		parsedPathURL, err := url.Parse(strings.TrimPrefix(parsedURL.Path, "/"))
		if err != nil {
			return "", err
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
