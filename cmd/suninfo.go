package cmd

import (
	"encoding/json"
	"flyawayhub-cli/config"
	"flyawayhub-cli/helpers"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type SunInfo struct {
	Sunrise                   int64 `json:"sunrise"`
	Sunset                    int64 `json:"sunset"`
	Transit                   int64 `json:"transit"`
	CivilTwilightBegin        int64 `json:"civil_twilight_begin"`
	CivilTwilightEnd          int64 `json:"civil_twilight_end"`
	NauticalTwilightBegin     int64 `json:"nautical_twilight_begin"`
	NauticalTwilightEnd       int64 `json:"nautical_twilight_end"`
	AstronomicalTwilightBegin int64 `json:"astronomical_twilight_begin"`
	AstronomicalTwilightEnd   int64 `json:"astronomical_twilight_end"`
}

func HandleSunInfoCommand(output string) {
	sessionData, err := LoadSession()
	if err != nil {
		fmt.Println("Error loading session:", err)
		return
	}

	fetchSunInfo(*sessionData, output)
}

func fetchSunInfo(sessionData SessionData, output string) {
	url := config.APIEndpoint("locations/%s/suninfo")

	client := &http.Client{}
	reqURL := fmt.Sprintf(url, sessionData.Locations[0].ID)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+sessionData.AccessToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	//fmt.Println(string(body))

	var sunInfo SunInfo // Adjusted to expect a single instance instead of a slice
	if err := json.Unmarshal(body, &sunInfo); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Create and populate the table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Event Type",
		"Timestamp",
		"Local Time",
		"UTC"})

	// Define a slice of event types and corresponding times
	events := []struct {
		EventType string
		timestamp string
		LocalTime string
		UTCTime   string
	}{
		{
			"Sunrise",
			strconv.FormatInt(sunInfo.Sunrise, 10),
			helpers.FormatLocalDateTime(sunInfo.Sunrise),
			helpers.FormatUTCDateTime(sunInfo.Sunrise),
		},
		{
			"Sunset",
			strconv.FormatInt(sunInfo.Sunset, 10),
			helpers.FormatLocalDateTime(sunInfo.Sunset),
			helpers.FormatUTCDateTime(sunInfo.Sunset),
		},
		{
			"Solar Noon",
			strconv.FormatInt(sunInfo.Transit, 10),
			helpers.FormatLocalDateTime(sunInfo.Transit),
			helpers.FormatUTCDateTime(sunInfo.Transit),
		},
		{
			"Civil Twilight Begin",
			strconv.FormatInt(sunInfo.CivilTwilightBegin, 10),
			helpers.FormatLocalDateTime(sunInfo.CivilTwilightBegin),
			helpers.FormatUTCDateTime(sunInfo.CivilTwilightBegin),
		},
		{
			"Civil Twilight End",
			strconv.FormatInt(sunInfo.CivilTwilightEnd, 10),
			helpers.FormatLocalDateTime(sunInfo.CivilTwilightEnd),
			helpers.FormatUTCDateTime(sunInfo.CivilTwilightEnd),
		},
		{
			"Nautical Twilight Begin",
			strconv.FormatInt(sunInfo.NauticalTwilightBegin, 10),
			helpers.FormatLocalDateTime(sunInfo.NauticalTwilightBegin),
			helpers.FormatUTCDateTime(sunInfo.NauticalTwilightBegin),
		},
		{
			"Nautical Twilight End",
			strconv.FormatInt(sunInfo.NauticalTwilightEnd, 10),
			helpers.FormatLocalDateTime(sunInfo.NauticalTwilightEnd),
			helpers.FormatUTCDateTime(sunInfo.NauticalTwilightEnd),
		},
		{
			"Astronomical Twilight Begin",
			strconv.FormatInt(sunInfo.AstronomicalTwilightBegin, 10),
			helpers.FormatLocalDateTime(sunInfo.AstronomicalTwilightBegin),
			helpers.FormatUTCDateTime(sunInfo.AstronomicalTwilightBegin),
		},
		{
			"Astronomical Twilight End",
			strconv.FormatInt(sunInfo.AstronomicalTwilightEnd, 10),
			helpers.FormatLocalDateTime(sunInfo.AstronomicalTwilightEnd),
			helpers.FormatUTCDateTime(sunInfo.AstronomicalTwilightEnd),
		},
	}

	// Iterate over events to populate the table
	for _, event := range events {
		table.Append([]string{
			event.EventType,
			event.timestamp,
			event.LocalTime,
			event.UTCTime})
	}

	// Render the table to stdout
	table.Render()
}
