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

type ScheduleMember struct {
	ID           string                   `json:"id"`
	Display      string                   `json:"display"`
	UserID       string                   `json:"user_id"`
	ProfileImage *string                  `json:"profile_image"` // Using pointer to handle null values
	SortOrder    int                      `json:"sort_order"`
	Locations    []ScheduleMemberLocation `json:"location"`
}

type ScheduleMemberLocation struct {
	ID          string  `json:"id"`
	MemberID    string  `json:"member_id"`
	LocationID  string  `json:"location_id"`
	Display     string  `json:"display"`
	PhoneNumber *string `json:"phone_number"` // Using pointer to handle null values
	Status      int     `json:"status"`
}

type Response struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Type    string `json:"type"`
	Models  Models `json:"models"`
}

type Models struct {
	Schedule []Schedule `json:"schedule"`
}

type Schedule struct {
	ID           string  `json:"id"`
	ResourceID   string  `json:"resource_id"`
	ResourceType string  `json:"resource_type"`
	Day          string  `json:"day"`
	Status       int     `json:"status"`
	Shifts       []Shift `json:"shifts"`
}

type Shift struct {
	ID             string  `json:"id"`
	StartTime      string  `json:"start_time"`
	EndTime        string  `json:"end_time"`
	WorkscheduleID string  `json:"workschedule_id"`
	Oncall         int     `json:"oncall"`
	CallTo         *string `json:"call_to"` // Using pointer to handle potential null values
	ResourcePhone  *string `json:"resource_phone"`
	LocationID     *string `json:"location_id"`
}

type FlightTypes struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type MembersPeople struct {
	ID           string                  `json:"id"`
	Display      string                  `json:"display"`
	UserId       string                  `json:"user_id"`
	ProfileImage string                  `json:"profile_image"`
	Location     []MembersPeopleLocation `json:"location"` // Changed "locations" to "location"
}

type MembersPeopleLocation struct {
	ID          string `json:"id"`
	MemberId    string `json:"member_id"`
	LocationId  string `json:"location_id"`
	Display     string `json:"display"`
	PhoneNumber string `json:"phone_number"`
	Status      int64  `json:"status"`
}

type ResourcesScheduleApiResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Type    string `json:"type"`
	Models  Models `json:"models"`
}

func HandleSchedulesCommand(output string) {
	sessionData, err := LoadSession()
	if err != nil {
		fmt.Println("Error loading session:", err)
		return
	}

	fetchSchedules(*sessionData, output)
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "Name", "Hangar", "Location ID", "Location airport code"})

	var i = 0
	for _, aircraft := range aircrafts {
		row := []string{
			strconv.Itoa(i),
			aircraft.ID,
			aircraft.Display,
			aircraft.Hangar,
			aircraft.Location.ID,
			aircraft.Location.AirportCode,
		}
		table.Append(row)
		i++
	}

	table.Render()

}

func fetchInstructors(sessionData SessionData) {
	expand := "location,contacts"
	url := config.APIEndpoint("organizations/%s/instructors?expand=" + expand)

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
	var instructors []ScheduleMember
	err = json.Unmarshal(body, &instructors)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"#",
		"ID",
		"User ID",
		"Name",
		"Location ID",
		"Location airport code"})

	var i = 0
	for _, instructor := range instructors {
		row := []string{
			strconv.Itoa(i),
			instructor.ID,
			instructor.UserID,
			instructor.Display,
			instructor.Locations[0].LocationID,
			instructor.Locations[0].Display,
		}
		table.Append(row)
		i++
	}

	table.Render()
}

func fetchFlightTypes(sessionData SessionData) {
	url := config.APIEndpoint("flighttypes/%s")

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
	var flightTypes []FlightTypes
	err = json.Unmarshal(body, &flightTypes)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"#",
		"ID",
		"Name",
		"Color"})

	var i = 0
	for _, flightType := range flightTypes {
		row := []string{
			strconv.Itoa(i),
			strconv.FormatInt(flightType.ID, 10),
			flightType.Name,
			flightType.Color,
		}
		table.Append(row)
		i++
	}

	table.Render()
}

func fetchMembersPeople(sessionData SessionData) {
	expand := "location"
	url := config.APIEndpoint("members/%s/people?expand=" + expand)

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
	var membersPeople []MembersPeople
	err = json.Unmarshal(body, &membersPeople)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"#",
		"ID",
		"User ID",
		"Name",
		"Location Main ID",
		"Location ID",
		"Location Member ID",
		"Location Name",
		"Location Phone Number",
		"Location Status"})

	var i = 0
	for _, memberPeople := range membersPeople {
		var locationMainId = ""
		var locationId = ""
		var locationMemberId = ""
		var locationDisplay = ""
		var locationPhoneNumber = ""
		var locationStatus = ""

		if len(memberPeople.Location) > 0 {
			locationMainId = memberPeople.Location[0].ID
			locationId = memberPeople.Location[0].LocationId
			locationMemberId = memberPeople.Location[0].MemberId
			locationDisplay = memberPeople.Location[0].Display
			locationPhoneNumber = memberPeople.Location[0].PhoneNumber
			locationStatus = strconv.FormatInt(memberPeople.Location[0].Status, 10)
		}

		row := []string{
			strconv.Itoa(i),
			memberPeople.ID,
			memberPeople.UserId,
			memberPeople.Display,
			locationMainId,
			locationId,
			locationMemberId,
			locationDisplay,
			locationPhoneNumber,
			locationStatus,
		}
		table.Append(row)
		i++
	}

	table.Render()
}

func fetchResourcesSchedule(sessionData SessionData) {
	startTime := 1704794972
	endTime := 1738922972
	resources := "{\"aircraft\":[\"76ace3fb-6e7c-11ee-8d38-025d5774c5eb\",\"fdd28866-a7e2-11eb-bd86-acde48001122\",\"fdd288fc-a7e2-11eb-bd86-acde48001122\",\"89ed2ac9-9eb4-11ec-9414-025d5774c5eb\",\"25779fa2-a1f3-11ec-9414-025d5774c5eb\"],\"instructor\":[\"5a505e0b-bc24-11ec-9414-025d5774c5eb\",\"ddb5b2b2-03b6-11ee-8d38-025d5774c5eb\",\"b4c61062-cf2b-11ed-b69f-025d5774c5eb\",\"469f652a-a7e1-11eb-bd86-acde48001122\",\"bbfb41b0-132e-11ed-9414-025d5774c5eb\",\"a7c52203-4cd7-11ee-8d38-025d5774c5eb\",\"469f6732-a7e1-11eb-bd86-acde48001122\"]}"
	url := config.APIEndpoint("resources/%s/schedule?start_time=%d&end_time=%d&resources=%s")

	client := &http.Client{}
	reqURL := fmt.Sprintf(url, sessionData.Id, startTime, endTime, resources)

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
	var apiResponse ResourcesScheduleApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"#",
		"ID",
		"Resource ID",
		"Resource Type",
		"Day",
		"Status",
	})

	var i = 0
	for _, schedule := range apiResponse.Models.Schedule {
		row := []string{
			strconv.Itoa(i),
			schedule.ID,
			schedule.ResourceID,
			schedule.ResourceType,
			schedule.Day,
			strconv.Itoa(schedule.Status),
		}
		table.Append(row)
		i++
	}

	table.Render()
}

func fetchSchedules(sessionData SessionData, output string) {
	//fetchAircraftsOrganization(sessionData)
	//fetchInstructors(sessionData)
	//fetchFlightTypes(sessionData)
	//fetchMembersPeople(sessionData)
	fetchResourcesSchedule(sessionData)

	url := config.APIEndpoint("resources/organization/%s/schedule")

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
	var instructors []ScheduleMember
	err = json.Unmarshal(body, &instructors)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"#",
		"ID",
		"Start Time",
		"End Time"})

	var i = 0
	for _, instructor := range instructors {
		row := []string{
			strconv.Itoa(i),
			instructor.ID,
			instructor.UserID,
			instructor.Display,
			instructor.Locations[0].LocationID,
			instructor.Locations[0].Display,
		}
		table.Append(row)
		i++
	}

	table.Render()
}

// DONE! flight type https://api.prod.flyawayhub.com/v1/flighttypes/78114770-a7e3-11eb-bd86-acde48001122
// DONE! people https://api.prod.flyawayhub.com/v1/members/78114770-a7e3-11eb-bd86-acde48001122/people?expand=location
// resources schedule https://api.prod.flyawayhub.com/v1/resources/78114770-a7e3-11eb-bd86-acde48001122/schedule?start_time=1704728168&end_time=1738856168&resources={"aircraft":["76ace3fb-6e7c-11ee-8d38-025d5774c5eb","fdd28866-a7e2-11eb-bd86-acde48001122","fdd288fc-a7e2-11eb-bd86-acde48001122","89ed2ac9-9eb4-11ec-9414-025d5774c5eb","25779fa2-a1f3-11ec-9414-025d5774c5eb"],"instructor":["5a505e0b-bc24-11ec-9414-025d5774c5eb","ddb5b2b2-03b6-11ee-8d38-025d5774c5eb","b4c61062-cf2b-11ed-b69f-025d5774c5eb","469f652a-a7e1-11eb-bd86-acde48001122","bbfb41b0-132e-11ed-9414-025d5774c5eb","a7c52203-4cd7-11ee-8d38-025d5774c5eb","469f6732-a7e1-11eb-bd86-acde48001122"]}
// Reservations public https://api.prod.flyawayhub.com/v1/reservations/78114770-a7e3-11eb-bd86-acde48001122/public?expand=resources.resource.phone_number&resources={"dispatcher":[],"room":[],"aircraft":["76ace3fb-6e7c-11ee-8d38-025d5774c5eb","fdd28866-a7e2-11eb-bd86-acde48001122","fdd288fc-a7e2-11eb-bd86-acde48001122","89ed2ac9-9eb4-11ec-9414-025d5774c5eb","25779fa2-a1f3-11ec-9414-025d5774c5eb"],"instructor":["5a505e0b-bc24-11ec-9414-025d5774c5eb","ddb5b2b2-03b6-11ee-8d38-025d5774c5eb","b4c61062-cf2b-11ed-b69f-025d5774c5eb","469f652a-a7e1-11eb-bd86-acde48001122","bbfb41b0-132e-11ed-9414-025d5774c5eb","a7c52203-4cd7-11ee-8d38-025d5774c5eb","469f6732-a7e1-11eb-bd86-acde48001122"]}&start_time=1707282001&end_time=1707368399&personal=false
