package cmd

import (
	"encoding/json"
	"flyawayhub-cli/config"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Aircraft struct {
	ID            string   `json:"id"`
	Display       string   `json:"display"`
	Logo          string   `json:"logo"`
	UserID        *string  `json:"user_id"` // Using pointer to handle null
	Simulator     bool     `json:"simulator"`
	SortOrder     int      `json:"sort_order"`
	Seats         *int     `json:"seats"`      // Using pointer to handle null
	Horsepower    *int     `json:"horsepower"` // Using pointer to handle null
	Hangar        string   `json:"hangar"`
	NumberOfSnags int      `json:"number_of_snags"`
	Location      Location `json:"location"`
	CurrentTTAF   *float64 `json:"currentTTAF"` // Using pointer to handle null
	Fuel          *string  `json:"fuel"`        // Using pointer to handle null
}

/*type Location struct {
	ID          string `json:"id"`
	AirportCode string `json:"airport_code"`
	MainBase    int    `json:"main_base"`
	Display     string `json:"display"`
}*/

func HandleSchedulesCommand() {
	sessionData, err := LoadSession()
	if err != nil {
		fmt.Println("Error loading session:", err)
		return
	}

	fetchSchedules(*sessionData)
}

func fetchAircraftsOrganization(sessionData SessionData) {

	expand := "fuel,location,hangar,seats,horsepower,upcoming_maintenance_by_ttaf,upcoming_maintenance_by_date,number_of_snags,currentTTAF"
	url := config.APIEndpoint("aircrafts/%s/organization?expand=" + expand)

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

	// Assuming you have a variable `body` which is a []byte containing the JSON
	var aircrafts []Aircraft
	err = json.Unmarshal(body, &aircrafts)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

}

func fetchSchedules(sessionData SessionData) {
	fetchAircraftsOrganization(sessionData)
}
