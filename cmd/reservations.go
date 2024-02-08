package cmd

import (
	"encoding/json"
	"flyawayhub-cli/config"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Define structs to match the JSON structure
type Flight struct {
	ID         string     `json:"id"`
	StartTime  int64      `json:"start_time"`
	EndTime    int64      `json:"end_time"`
	FlightType FlightType `json:"flight_type"`
	Resources  []Resource `json:"resources"`
	CreatedAt  int64      `json:"created_at"`
}

type FlightType struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Resource struct {
	ResourceType string         `json:"resource_type"`
	Resource     ResourceDetail `json:"resource"`
}

type ResourceDetail struct {
	ID      string `json:"id"`
	Display string `json:"display"`
}

// HandleReservationsCommand fetches and displays reservations
func HandleReservationsCommand(output string) {
	sessionData, err := LoadSession()
	if err != nil {
		fmt.Println("Error loading session:", err)
		return
	}

	fetchReservations(*sessionData)
}

// fetchReservations makes an HTTP GET request to the reservations endpoint
func fetchReservations(sessionData SessionData) {
	currentTime := time.Now()
	currentTimestamp := currentTime.Unix()

	expand := "resources.resource.tracking_list"
	startTime := strconv.FormatInt(currentTimestamp, 10)
	page := "0"
	limit := "20"
	url := config.APIEndpoint("reservations/%s/organization?expand=" + expand + "&start_time=" + startTime + "&page=" + page + "&limit=" + limit)

	client := &http.Client{}
	reqURL := fmt.Sprintf(url, sessionData.Id)

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

	var flights []Flight
	if err := json.Unmarshal(body, &flights); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Now `flights` is populated with data from `body`, proceed to create the table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"#",
		"ID",
		"Date",
		"Flight Type",
		"Start",
		"End",
		"Aircraft",
		"Crew"})

	var i = 0
	for _, flight := range flights {
		dateStr, _ := FormatTime(flight.StartTime, "Mon, Jan 2, 2006")
		startTimeStr, _ := FormatTime(flight.StartTime, "1504")
		endTimeStr, _ := FormatTime(flight.EndTime, "1504")
		aircraft, crew := extractResourceDetails(flight.Resources)

		row := []string{
			strconv.Itoa(i),
			flight.ID,
			dateStr,
			flight.FlightType.Name,
			startTimeStr,
			endTimeStr,
			aircraft,
			crew,
		}
		table.Append(row)
		i++
	}

	table.Render()
}

func extractResourceDetails(resources []Resource) (aircraft, crew string) {
	aircraft, crew = "N/A", "N/A" // Default values
	for _, resource := range resources {
		if resource.ResourceType == "aircraft" {
			aircraft = resource.Resource.Display
		} else if resource.ResourceType == "instructor" {
			crew = resource.Resource.Display
		}
	}
	return
}
