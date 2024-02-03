package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// Assuming the ApiConfig struct and loadConfig function are defined in a shared place or replicated here

var listReservationsCmd = &cobra.Command{
	Use:   "list-reservations",
	Short: "List your flying reservations",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := loadConfig() // Load API configuration
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading configuration: %v\n", err)
			return
		}
		listReservations(config)
	},
}

func listReservations(config *ApiConfig) {
	// Replace "config.ApiPath" with the specific path for listing reservations if different
	url := fmt.Sprintf("%s://%s%s", config.ApiScheme, config.ApiHost, config.ApiPath)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add necessary headers, including Authorization
	req.Header.Add("Authorization", "Bearer "+config.BearerToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to API endpoint:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response:", string(body))
}

/*func init() {
	// Register the listReservationsCmd with the Cobra root command
	rootCmd.AddCommand(listReservationsCmd)
}*/
