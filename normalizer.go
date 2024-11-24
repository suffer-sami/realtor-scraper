package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/purell"
	"github.com/nyaruka/phonenumbers"
)

func normalizeAgent(agent *Agent) {
	if agent.Href != "" {
		if normalizedHref, err := normalizeUrl(agent.Href); err == nil {
			agent.Href = normalizedHref
		}
	}

	if agent.Office.Website != "" {
		if normalizedHref, err := normalizeUrl(agent.Office.Website); err == nil {
			agent.Office.Website = normalizedHref
		}
	}

	agentAddressCountry := "US"
	if agent.Address.Country != "" {
		agentAddressCountry = agent.Address.Country[:2]
	}

	for i := range agent.Phones {
		phone := &agent.Phones[i]
		normalizePhone(phone, agentAddressCountry)
	}

	officeCountry := agentAddressCountry
	if agent.Office.Address.Country != "" {
		officeCountry = agent.Office.Address.Country[:2]
	}

	for i := range agent.Office.Phones {
		phone := &agent.Office.Phones[i]
		normalizePhone(phone, officeCountry)
	}
	for key, phone := range agent.Office.PhoneList {
		normalizePhone(&phone, officeCountry)
		agent.Office.PhoneList[key] = phone
	}

	for k, v := range agent.SocialMedias {
		normalizedHref, err := normalizeUrl(v.Href)
		if err != nil {
			log.Printf("Error Normalizing Url: %v", err)
			continue
		}

		agent.SocialMedias[k] = SocialMedia{
			Href: normalizedHref,
			Type: v.Type,
		}
	}

}

func normalizePhone(phone *Phone, regionCode string) {
	parsedNumber, err := phonenumbers.Parse(phone.Number, regionCode)
	if err != nil {
		log.Printf("error parsing phone number: %v", err)
		return
	}

	phone.Number = phonenumbers.Format(parsedNumber, phonenumbers.INTERNATIONAL)
}

func normalizeUrl(rawURL string) (string, error) {
	parsedURL, err := url.Parse(strings.ToLower(rawURL))
	if err != nil {
		return "", err
	}

	if strings.Contains(parsedURL.Path, strings.TrimPrefix(parsedURL.Host, "www.")) || strings.Contains(parsedURL.Path, ".") {
		parsedPathURL, err := url.Parse(strings.TrimPrefix(parsedURL.Path, "/"))
		if err != nil {
			return "", err
		}
		parsedURL.Path = parsedPathURL.Path

	}

	normalizedURL := purell.NormalizeURL(parsedURL, purell.FlagsUsuallySafeGreedy)

	return normalizedURL, nil
}
