package cmd

import (
	"bufio"
	"context"
	appConfig "flyawayhub-cli/config"
	"flyawayhub-cli/logging"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	logger   logging.Logger
	cfgFile  string
	username string
)

// LoginCommand Define a struct to hold dependencies for the login command.
type LoginCommand struct {
	logger logging.Logger
	// Add other dependencies here
}

// NewLoginCommand creates a new login cobra.Command with dependencies injected.
func NewLoginCommand(logger logging.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login will sign you in to " + appConfig.AppName + " (will generate an authorization bearer).",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Starting login process")
			// Your login logic here...
			if err := loginProcess(); err != nil {
				logger.Error("Login process failed", "error", err)
				return
			}
			logger.Info("Login process completed successfully")
		},
	}
	return cmd
}

func sendCredentials(username, password string) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(appConfig.AWSCognitoRegion))
	if err != nil {
		logger.Error("Unable to load SDK config", "error", err)
		return
	}

	// Create a Cognito Identity Provider client
	svc := cognitoidentityprovider.NewFromConfig(cfg)

	// Replace these values with your Cognito app's details
	clientId := appConfig.AWSCognitoClientID

	// Perform the authentication request
	authResp, err := svc.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth, // Use the enum for the auth flow type
		AuthParameters: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},
		ClientId: aws.String(clientId),
	})
	if err != nil {
		logger.Error("authentication failed: %v\n", err)
		return
	}

	// Fetch and update session with organization info
	if err := FetchOrganizationInfo(*authResp.AuthenticationResult.AccessToken, "json"); err != nil {
		logger.Error("Failed to fetch organization info: %v\n", err)
		// Decide how you want to handle this error. For now, just printing the error.
	}
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login will sign you in to " + appConfig.AppName + " (will generate an authorization bearer).",
	Run:   login,
}

func login(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println("\nError reading password")
		return
	}
	password := string(bytePassword)

	// Send credentials using the AWS SDK
	sendCredentials(username, password)
}

func InitCommands(root *cobra.Command) {
	root.AddCommand(loginCmd)
	// Add other commands as needed
}
