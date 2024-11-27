package main

import (
	"encoding/json"
	"time"
)

type SearchRequestParams struct {
	Offset                  int         `json:"offset"`
	Limit                   int         `json:"limit"`
	MarketingAreaCities     string      `json:"marketing_area_cities"`
	PostalCode              string      `json:"postal_code"`
	IsPostalSearch          string      `json:"is_postal_search"`
	Name                    string      `json:"name"`
	Types                   string      `json:"types"`
	Sort                    string      `json:"sort"`
	FarOptOut               string      `json:"far_opt_out"`
	ClientID                string      `json:"client_id"`
	RecommendationsCountMin string      `json:"recommendations_count_min"`
	AgentRatingMin          string      `json:"agent_rating_min"`
	Languages               string      `json:"languages"`
	AgentType               string      `json:"agent_type"`
	PriceMin                string      `json:"price_min"`
	PriceMax                string      `json:"price_max"`
	Designations            string      `json:"designations"`
	Photo                   string      `json:"photo"`
	SeoUserType             SeoUserType `json:"seoUserType"`
	IsCountySearch          string      `json:"is_county_search"`
	County                  string      `json:"county"`
}

type SeoUserType struct {
	IsBot      string `json:"isBot"`
	DeviceType string `json:"deviceType"`
}

type SearchRequestResponse struct {
	Agents       []Agent `json:"agents"`
	MatchingRows int     `json:"matching_rows"`
}

type Agent struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	NickName   string `json:"nick_name"`
	PersonName string `json:"person_name"`
	Title      string `json:"title"`
	Slogan     string `json:"slogan"`

	Email   string  `json:"email"`
	Phones  []Phone `json:"phones"`
	Address Address `json:"address"`

	Photo           Photo `json:"photo"`
	BackgroundPhoto Photo `json:"background_photo"`
	HasPhoto        bool  `json:"has_photo"`

	Role             string           `json:"role"`
	AgentType        []string         `json:"agent_type"`
	Types            string           `json:"types"`
	AgentRating      int              `json:"agent_rating"`
	AgentTeamDetails AgentTeamDetails `json:"agent_team_details"`

	Broker Broker `json:"broker"`
	Office Office `json:"office"`

	Description          string           `json:"description"`
	Designations         []Designation    `json:"designations"`
	Specializations      []Specialization `json:"specializations"`
	RecommendationsCount int              `json:"recommendations_count"`
	ReviewCount          int              `json:"review_count"`

	ServedAreas         []Area   `json:"served_areas"`
	MarketingAreaCities []Area   `json:"marketing_area_cities"`
	Zips                []string `json:"zips"`
	Languages           []string `json:"languages"`
	UserLanguages       []string `json:"user_languages"`

	Mls             []MLS        `json:"mls"`
	MlsHistory      []MlsHistory `json:"mls_history"`
	MlsMonetization bool         `json:"mls_monetization"`

	AdvertiserID int    `json:"advertiser_id"`
	PartyID      int    `json:"party_id"`
	NrdsID       string `json:"nrds_id"`
	NarOnly      int    `json:"nar_only"`

	FeedLicenses []FeedLicense `json:"feed_licenses"`
	IsRealtor    bool          `json:"is_realtor"`
	LastUpdated  string        `json:"last_updated"`
	FirstMonth   int           `json:"first_month"`
	FirstYear    json.Number   `json:"first_year"`

	SocialMedias map[string]SocialMedia `json:"social_media"`
	Video        string                 `json:"video"`
	WebURL       string                 `json:"web_url"`

	RecentlySold RecentlySold `json:"recently_sold"`
	ForSalePrice ForSalePrice `json:"for_sale_price"`

	Settings Settings `json:"settings"`
	Href     string   `json:"href"`
}

type Address struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line       string `json:"line"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
	StateCode  string `json:"state_code"`
}

type Photo struct {
	Href     string `json:"href"`
	IsZoomed bool   `json:"is_zoomed"`
}

type Broker struct {
	AccentColor   string        `json:"accent_color"`
	Designations  []Designation `json:"designations"`
	FulfillmentID int           `json:"fulfillment_id"`
	Name          string        `json:"name"`
	Photo         Photo         `json:"photo"`
	Video         string        `json:"video"`
}

type Designation struct {
	Name string `json:"name"`
}

type Specialization struct {
	Name string `json:"name"`
}

type MLS struct {
	Abbreviation  string `json:"abbreviation"`
	LicenseNumber string `json:"license_number"`
	Primary       bool   `json:"primary"`
	Type          string `json:"type"`
	MemberID      string `json:"member.id"`
}

type Area struct {
	CityState string `json:"city_state"`
	Name      string `json:"name"`
	StateCode string `json:"state_code"`
}

type AgentTeamDetails struct {
	IsTeamMember bool `json:"is_team_member"`
}

type FeedLicense struct {
	Country       string `json:"country"`
	LicenseNumber string `json:"license_number"`
	StateCode     string `json:"state_code"`
}

type PriceStats struct {
	Count int `json:"count"`
	Max   int `json:"max"`
	Min   int `json:"min"`
}

type ForSalePrice struct {
	LastListingDate time.Time `json:"last_listing_date"`
	PriceStats
}

type RecentlySold struct {
	LastSoldDate string `json:"last_sold_date"`
	PriceStats
}

type MlsHistory struct {
	Abbreviation     string    `json:"abbreviation"`
	InactivationDate time.Time `json:"inactivation_date"`
	LicenseNumber    string    `json:"license_number"`
	Member           Member    `json:"member"`
	Primary          bool      `json:"primary"`
	Type             string    `json:"type"`
}

type Member struct {
	ID string `json:"id"`
}

type Phone struct {
	Ext     string `json:"ext"`
	Number  string `json:"number"`
	Type    string `json:"type"`
	IsValid bool   `json:"is_valid"`
}

type Office struct {
	Name         string           `json:"name"`
	Address      Address          `json:"address"`
	Phones       []Phone          `json:"phones"`
	PhoneList    map[string]Phone `json:"phone_list"`
	Photo        Photo            `json:"photo"`
	Website      string           `json:"website"`
	FeedLicenses []FeedLicense    `json:"feed_licenses"`
}

type SocialMedia struct {
	Href string `json:"href"`
	Type string `json:"type"`
}

type Settings struct {
	BrokerDataFeedOptOut  bool `json:"broker_data_feed_opt_out"`
	DisplayListings       bool `json:"display_listings"`
	DisplayPriceRange     bool `json:"display_price_range"`
	DisplayRatings        bool `json:"display_ratings"`
	DisplaySoldListings   bool `json:"display_sold_listings"`
	FarOverride           bool `json:"far_override"`
	FullAccess            bool `json:"full_access"`
	HasDotrealtor         bool `json:"has_dotrealtor"`
	LoadedFromSb          bool `json:"loaded_from_sb"`
	NewFeaturePopupClosed struct {
		AgentLeftNavAvatarToProfile bool `json:"agent_left_nav_avatar_to_profile"`
	} `json:"new_feature_popup_closed"`
	ProfileWizard struct {
		RealsatisfiedOptOut bool `json:"realsatisfied_opt_out"`
		TtOptOut            bool `json:"tt_opt_out"`
	} `json:"profile_wizard"`
	Recommendations struct {
		Realsatisfied struct {
			ID      string `json:"id"`
			Linked  string `json:"linked"`
			Updated string `json:"updated"`
			User    string `json:"user"`
		} `json:"realsatisfied"`
	} `json:"recommendations"`
	ReviewsOptOut struct {
		Rdc bool `json:"rdc"`
		Rs  bool `json:"rs"`
	} `json:"reviews_opt_out"`
	ShareContacts bool `json:"share_contacts"`
	ShowStream    bool `json:"show_stream"`
	TermsOfUse    bool `json:"terms_of_use"`
	Unsubscribe   struct {
		AccountNotify bool `json:"account_notify"`
		Autorecs      bool `json:"autorecs"`
		Recapprove    bool `json:"recapprove"`
	} `json:"unsubscribe"`
	UseDotRealtorForSearchEngines bool `json:"use_dot_realtor_for_search_engines"`
}
