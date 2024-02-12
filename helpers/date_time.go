// Package helpers provides utility functions for date and time formatting.
package helpers

import "time"

const (
	DateTimeFormatLocal     = "2006-01-02 15:04:05 MST"
	NameMonthDateTimeFormat = "Jan 2, 2006 15:04 MST"
	Time                    = "15:04 MST"
)

// FormatLocalDateTime formats a Unix timestamp into a human-readable date and time string
// in the local time zone according to the given format string.
func FormatLocalDateTime(unixTime int64, format string) string {
	return time.Unix(unixTime, 0).Local().Format(format)
}

// FormatUTCDateTime formats a Unix timestamp into a human-readable date and time string
// in the UTC time zone according to the given format string.
func FormatUTCDateTime(unixTime int64, format string) string {
	return time.Unix(unixTime, 0).UTC().Format(format)
}
