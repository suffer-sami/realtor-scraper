package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/goccy/go-json"
)

// Converts a string to sql.NullString
func stringToNullString(s string) sql.NullString {
	s = strings.TrimSpace(s)
	return sql.NullString{String: s, Valid: s != ""}
}

// Converts any object to json string
func anyToJsonString(o any) (string, error) {
	data, err := json.Marshal(o)
	if err != nil {
		return "", fmt.Errorf("failed to encode json: %w", err)
	}
	return string(data), nil
}

// Converts an int to sql.NullInt64
func intToNullInt64(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: i != 0}
}

// Converts an int64 to sql.NullInt64
func int64ToNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: i != 0}
}

// Converts a bool to sql.NullBool
func boolToNullBool(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}

// Converts an numeric type to sql.NullInt64
func numericToNullInt(i NumericType) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: i != 0}
}

// Converts time.Time to sql.NullTime
func timeToNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: t, Valid: true}
}

// Converts a string w/ layout to sql.NullTime
func strToNullTime(t string, layout string) sql.NullTime {
	parsedTime, err := time.Parse(layout, t)
	if err != nil || t == "" || parsedTime.IsZero() {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: parsedTime, Valid: true}
}
