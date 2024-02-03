package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type ApiConfigurations struct {
	Api map[string]EndpointConfig `json:"api"`
}

type EndpointConfig struct {
	Scheme      string `json:"scheme"`
	Host        string `json:"host"`
	Path        string `json:"path"`
	BearerToken string `json:"bearerToken"`
}

func loadConfig() (*ApiConfigurations, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("opening config file: %w", err)
	}
	defer file.Close()

	var configs ApiConfigurations
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configs); err != nil {
		return nil, fmt.Errorf("decoding config file: %w", err)
	}

	return &configs, nil
}
