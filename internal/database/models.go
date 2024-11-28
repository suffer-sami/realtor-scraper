// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"
)

type Agent struct {
	ID                   string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	FirstName            sql.NullString
	LastName             sql.NullString
	NickName             sql.NullString
	PersonName           sql.NullString
	Title                sql.NullString
	Slogan               sql.NullString
	Email                sql.NullString
	AgentRating          sql.NullInt64
	Description          sql.NullString
	RecommendationsCount sql.NullInt64
	ReviewCount          sql.NullInt64
	LastUpdated          sql.NullTime
	FirstMonth           sql.NullInt64
	FirstYear            sql.NullInt64
	Video                sql.NullString
	WebUrl               sql.NullString
	Href                 sql.NullString
}

type AgentLanguage struct {
	AgentID    sql.NullString
	LanguageID sql.NullInt64
}

type AgentMultipleListingService struct {
	AgentID                  sql.NullString
	MultipleListingServiceID sql.NullInt64
}

type AgentUserLanguage struct {
	AgentID    sql.NullString
	LanguageID sql.NullInt64
}

type FeedLicense struct {
	ID            int64
	Country       sql.NullString
	LicenseNumber sql.NullString
	StateCode     sql.NullString
	AgentID       sql.NullString
}

type Language struct {
	ID   int64
	Name sql.NullString
}

type ListingsDatum struct {
	ID              int64
	Count           sql.NullInt64
	Min             sql.NullInt64
	Max             sql.NullInt64
	LastListingDate sql.NullTime
	AgentID         sql.NullString
}

type MultipleListingService struct {
	ID               int64
	Abbreviation     sql.NullString
	InactivationDate sql.NullTime
	LicenseNumber    sql.NullString
	MemberID         sql.NullString
	Type             sql.NullString
	IsPrimary        sql.NullBool
}

type Request struct {
	ID             int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Offset         int64
	ResultsPerPage int64
}

type SalesDatum struct {
	ID           int64
	Count        sql.NullInt64
	Min          sql.NullInt64
	Max          sql.NullInt64
	LastSoldDate sql.NullTime
	AgentID      sql.NullString
}

type SocialMedia struct {
	ID      int64
	Type    sql.NullString
	Href    sql.NullString
	AgentID sql.NullString
}
