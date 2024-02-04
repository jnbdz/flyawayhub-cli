package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Extended session data to include additional organization details
type SessionData struct {
	Token             string                `json:"token"`
	OrganizationID    string                `json:"organizationId"`
	Address           *string               `json:"address"`
	City              *string               `json:"city"`
	State             *string               `json:"state"`
	Country           *string               `json:"country"`
	Zipcode           string                `json:"zipcode"`
	RegistrationLink  string                `json:"registration_link"`
	Locations         []Location            `json:"locations"`
	MemberID          string                `json:"member_id"`
	MemberDisplay     string                `json:"member_display"`
	UserID            string                `json:"user_id"`
	DisplayName       string                `json:"display"`
	OrganizationType  OrgType               `json:"organizationtype"`
	AviationAuthority Authority             `json:"aviationauthority"`
	LogoID            *string               `json:"logo_id"`
	Logo              string                `json:"logo"`
	Name              string                `json:"name"`
	Email             string                `json:"email"`
	CountryCode       *string               `json:"country_code"`
	PhoneNumber       *string               `json:"phonenumber"`
	AirportCode       *string               `json:"airport_code"`
	URL               string                `json:"url"`
	TimeZone          TimeZone              `json:"timeZone"`
	MemberRoles       []MemberRole          `json:"memberroles"`
	Permissions       map[string]Permission `json:"permissions"`
}

type Location struct {
	ID          string `json:"id"`
	AirportCode string `json:"airport_code"`
	MainBase    int    `json:"main_base"`
	Display     string `json:"display"`
}

type OrgType struct {
	ID      int    `json:"id"`
	Display string `json:"display"`
}

type Authority struct {
	ID      int    `json:"id"`
	Display string `json:"display"`
}

type TimeZone struct {
	ID      int    `json:"id"`
	Display string `json:"display"`
}

type MemberRole struct {
	ID         string `json:"id"`
	MemberID   string `json:"member_id"`
	UserRoleID int    `json:"userrole_id"`
	Role       string `json:"role"`
}

type Permission struct {
	ID       *string `json:"id"`
	IsEnable bool    `json:"isenable"`
}

// saveSession saves session data to a file
func saveSession(data SessionData) error {
	path, err := getSessionFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("creating session file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("encoding session data: %w", err)
	}

	return nil
}

// loadSession loads session data from a file
func loadSession() (*SessionData, error) {
	path, err := getSessionFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening session file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var data SessionData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding session data: %w", err)
	}

	return &data, nil
}

// getSessionFilePath determines the correct path for the session file based on the OS
func getSessionFilePath() (string, error) {
	var dir string
	switch operatingSystem := runtime.GOOS; operatingSystem {
	case "linux", "darwin":
		dir = os.Getenv("HOME")
		if operatingSystem == "linux" {
			dir = filepath.Join(dir, ".cache")
		} else { // darwin (macOS)
			dir = filepath.Join(dir, "Library", "Caches")
		}
	case "windows":
		dir = os.Getenv("AppData")
		if dir == "" {
			return "", fmt.Errorf("AppData not set")
		}
	default:
		return "", fmt.Errorf("unsupported platform")
	}

	dir = filepath.Join(dir, "flyawayhub")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", fmt.Errorf("creating session directory: %w", err)
	}

	return filepath.Join(dir, "session.json"), nil
}
