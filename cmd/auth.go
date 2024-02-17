package cmd

import (
	"bufio"
	"context"
	"flyawayhub-cli/auth"
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
	logger logging.Logger
)

type authService struct{}

// Login
func (s *authService) Login(ctx context.Context) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter Password: ")
	bytePassword, _ := terminal.ReadPassword(0)
	password := string(bytePassword)

	return sendCredentials(ctx, username, password)
}

func loginProcess(ctx context.Context, username, password string) error {
	err := sendCredentials(ctx, username, password)
	if err != nil {
		return err
	}
	return nil
}

// NewLoginCommand updated to use loginProcess within the Run function.
func NewLoginCommand(authService auth.Service, logger logging.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login will sign you in to " + appConfig.AppName,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background() // Or context.TODO() if unsure

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter Username: ")
			username, err := reader.ReadString('\n')
			if err != nil {
				logger.Error("Error reading username", "error", err)
				return err
			}
			username = strings.TrimSpace(username)

			fmt.Print("Enter Password: ")
			bytePassword, err := terminal.ReadPassword(0)
			if err != nil {
				logger.Error("Error reading password", "error", err)
				return err
			}
			password := string(bytePassword)

			// Use authService to perform login
			err = authService.Login(ctx, username, password)
			if err != nil {
				logger.Error("Login process failed", "error", err)
				return err
			}

			logger.Info("Login successful")
			return nil
		},
	}
	return cmd
}

// sendCredentials
func sendCredentials(ctx context.Context, username, password string) error {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(appConfig.AWSCognitoRegion))
	if err != nil {
		logger.Error("Unable to load SDK config", "error", err)
		return err
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
		return err
	}

	// Fetch and update session with organization info
	if err := FetchOrganizationInfo(*authResp.AuthenticationResult.AccessToken, "json"); err != nil {
		logger.Error("Failed to fetch organization info: %v\n", err)
		// Decide how you want to handle this error. For now, just printing the error.
	}

	logger.Info("Login successful")
	return nil
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login will sign you in to " + appConfig.AppName + " (will generate an authorization bearer).",
	Run:   login,
}

func login(cmd *cobra.Command, args []string) {
	ctx := context.Background() // Or context.TODO() if unsure
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err != nil {
		logger.Error("\nError reading password")
		return
	}
	password := string(bytePassword)

	// Send credentials using the AWS SDK
	err = sendCredentials(ctx, username, password)
	if err != nil {
		return
	}
}

func InitCommands(root *cobra.Command) {
	root.AddCommand(loginCmd)
	// Add other commands as needed
}
