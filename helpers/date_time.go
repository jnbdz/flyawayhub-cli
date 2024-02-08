package helpers

import "time"

// Function to format Unix timestamp to readable string
func FormatLocalDateTime(unixTime int64) string {
	// Check if unixTime needs division by 1000 to convert milliseconds to seconds
	return time.Unix(unixTime, 0).Local().Format("15:04 MST")
}

func FormatUTCDateTime(unixTime int64) string {
	// Check if unixTime needs division by 1000 to convert milliseconds to seconds
	return time.Unix(unixTime, 0).UTC().Format("15:04 MST")
}
