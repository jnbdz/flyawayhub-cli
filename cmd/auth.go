package cmd

import (
	"bufio"
	"context"
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
	appName = "Flyawayhub"
	cfgFile string

	username string
)

func sendCredentials(username, password string) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to load SDK config, %v\n", err)
		return
	}

	// Create a Cognito Identity Provider client
	svc := cognitoidentityprovider.NewFromConfig(cfg)

	// Replace these values with your Cognito app's details
	clientId := ""

	// Initiate the authentication request
	/*authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth, // Corrected enum value
		AuthParameters: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},
		ClientId: aws.String(clientId),
	}*/

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
		fmt.Fprintf(os.Stderr, "authentication failed: %v\n", err)
		return
	}

	// Handle the successful authentication response
	// For demonstration purposes, printing the access token
	fmt.Printf("Authentication successful. Access Token: %s\n", *authResp.AuthenticationResult.AccessToken)
}

func initConfig() {
	// Existing configuration loading code
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login will sign you in to " + appName + " (will generate an authorization bearer).",
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

func init() {
	cobra.OnInitialize(initConfig)
}

func InitCommands(root *cobra.Command) {
	root.AddCommand(loginCmd)
	// Add other commands as needed
}
