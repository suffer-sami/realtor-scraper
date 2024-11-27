package main

import (
	"database/sql"
	"time"
)

// Converts a string to sql.NullString
func toNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// Converts an int to sql.NullInt64
func toNullInt(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: i != 0}
}

// Converts an int64 to sql.NullInt64
func toNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: i != 0}
}

// Converts a string w/ layout to sql.NullTime
func toStrNullTime(t string, layout string) sql.NullTime {
	parsedTime, err := time.Parse(layout, t)
	if err != nil || t == "" {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: parsedTime, Valid: true}
}
