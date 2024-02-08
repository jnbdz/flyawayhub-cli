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

type Member struct {
	ID               string           `json:"id"`
	Display          string           `json:"display"`
	UserId           string           `json:"user_id"`
	IsVerify         bool             `json:"isverify"`
	Status           int64            `json:"status"`
	SortOrder        int64            `json:"sort_order"`
	OrganizationId   string           `json:"organization_id"`
	Permissions      MemberPermission `json:"permissions"`
	MemberRoles      []MemberRole     `json:"memberroles"`
	MembershipNumber string           `json:"membership_number"`
	IsVaccinated     int              `json:"is_vaccinated"`
	Email            string           `json:"email"`
	PhoneNumber      string           `json:"phone_number"`
}

type MemberPermission struct {
	SoloFlight    Permission `json:"solo_flight"`
	OnlineBooking Permission `json:"online_booking"`
}

func HandleMembersCommand(output string) {
	sessionData, err := LoadSession()
	if err != nil {
		fmt.Println("Error loading session:", err)
		return
	}

	fetchMembers(*sessionData)
}

func fetchMembers(sessionData SessionData) {
	currentTime := time.Now()
	currentTimestamp := currentTime.Unix()

	expand := "membership_number,isverify,email,phone_number,memberroles,profile_image,permissions,organization_id,status,sort_order,is_vaccinated"
	startTime := strconv.FormatInt(currentTimestamp, 10)
	page := "0"
	limit := "20"
	url := config.APIEndpoint("members/%s/people?expand=" + expand + "&start_time=" + startTime + "&page=" + page + "&limit=" + limit)

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

	var members []Member
	if err := json.Unmarshal(body, &members); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Now `flights` is populated with data from `body`, proceed to create the table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"#",
		"User ID",
		"Name",
		"Email",
		"Phone Number",
		"Is Vaccinated"})

	var i = 0
	for _, member := range members {
		row := []string{
			strconv.Itoa(i),
			member.UserId,
			member.Display,
			member.Email,
			member.PhoneNumber,
			ReadableIsVaccinated(member.IsVaccinated),
		}
		table.Append(row)
		i++
	}

	table.Render()
}
