package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// FetchOrganizationInfo fetches organization information and updates the session
func FetchOrganizationInfo(accessToken, output string) error {
	sessionData, err := LoadSession()
	if err != nil {
		return fmt.Errorf("loading session: %w", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.prod.flyawayhub.com/v1/organizations", nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	// Set the authorization header with the access token
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	// Assuming the API response matches the SessionData structure or is adaptable
	var orgData []SessionData
	if err := json.Unmarshal(body, &orgData); err != nil {
		return fmt.Errorf("unmarshaling response: %w", err)
	}

	if len(orgData) > 0 {
		orgData := orgData[0] // Extract the first organization's data

		// Merge the extracted organization data with the existing session data
		MergeStructs(sessionData, &orgData)

		// Restore the AccessToken and any other fields you wish to preserve
		sessionData.AccessToken = accessToken

		// Save the updated session data
		return saveSession(*sessionData)
	} else {
		return fmt.Errorf("no organization data found in response")
	}
}
