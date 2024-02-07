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

type ApiResponse struct {
	Name    string              `json:"name"`
	Message string              `json:"message"`
	Code    int                 `json:"code"`
	Status  string              `json:"status"`
	Type    string              `json:"type"`
	Models  NotificationsModels `json:"models"`
}

type NotificationsModels struct {
	Notifications []Notification `json:"notifications"`
}

type Notification struct {
	ID               string           `json:"id"`
	Message          string           `json:"message"`
	ShortMessage     string           `json:"short_message"`
	IsRead           int              `json:"isread"`
	NotificationType NotificationType `json:"notificationtype"`
	CreatedAt        int64            `json:"created_at"`
}

type NotificationType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func HandleNotificationsCommand() {
	sessionData, err := LoadSession()
	if err != nil {
		fmt.Println("Error loading session:", err)
		return
	}

	fetchNotifications(*sessionData)
}

func fetchNotifications(sessionData SessionData) {
	// https://api.prod.flyawayhub.com/v1/notifications/all?page=0&limit=20
	page := "0"
	limit := "20"
	url := config.APIEndpoint("notifications/all?=page=" + page + "&limit=" + limit)

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

	var response ApiResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "Message", "Short Message", "Notification Type Name", "Created At"})

	var i = 0
	for _, notification := range response.Models.Notifications {
		row := []string{
			strconv.Itoa(i),
			notification.ID,
			notification.Message,
			notification.ShortMessage,
			notification.NotificationType.Name,
			strconv.FormatInt(notification.CreatedAt, 10),
		}
		table.Append(row)
		i++
	}

	table.Render()

}
