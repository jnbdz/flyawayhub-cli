package helpers

import (
	"strings"
	"testing"
)

func TestFormatLocalDateTimeWithStandardFormat(t *testing.T) {
	const testTimestamp = 0
	expectedContains := "1969-12-31 19:00:00 EST"
	formattedTime := FormatLocalDateTime(testTimestamp, DateTimeFormatLocal)

	if !strings.Contains(formattedTime, expectedContains) {
		t.Errorf("Formatted local datetime should contain %s, got %s", expectedContains, formattedTime)
	}
}

func TestFormatLocalDateTimeWithNameMonthFormat(t *testing.T) {
	const testTimestamp = 0
	expectedContains := "Dec 31, 1969 19:00 EST"
	formattedTime := FormatLocalDateTime(testTimestamp, NameMonthDateTimeFormat)

	if !strings.Contains(formattedTime, expectedContains) {
		t.Errorf("Formatted local datetime should contain %s, got %s", expectedContains, formattedTime)
	}
}

func TestFormatUTCDateTimeWithStandardFormat(t *testing.T) {
	const testTimestamp = 0
	expected := "1970-01-01 00:00:00 UTC"
	formattedTime := FormatUTCDateTime(testTimestamp, DateTimeFormatLocal)

	if formattedTime != expected {
		t.Errorf("Expected UTC datetime to be %s, got %s", expected, formattedTime)
	}
}

func TestFormatUTCDateTimeWithNameMonthFormat(t *testing.T) {
	const testTimestamp = 0
	expected := "Jan 1, 1970 00:00 UTC"
	formattedTime := FormatUTCDateTime(testTimestamp, NameMonthDateTimeFormat)

	if formattedTime != expected {
		t.Errorf("Expected UTC datetime to be %s, got %s", expected, formattedTime)
	}
}

func TestFormatLocalDateTimeWithTimeFormat(t *testing.T) {
	const testTimestamp = 0
	expected := "19:00 EST"
	formattedTime := FormatLocalDateTime(testTimestamp, Time)

	if formattedTime != expected {
		t.Errorf("Expected UTC datetime to be %s, got %s", expected, formattedTime)
	}
}

func TestFormatUTCDateTimeWithTimeFormat(t *testing.T) {
	const testTimestamp = 0
	expected := "00:00 UTC"
	formattedTime := FormatUTCDateTime(testTimestamp, Time)

	if formattedTime != expected {
		t.Errorf("Expected UTC datetime to be %s, got %s", expected, formattedTime)
	}
}
