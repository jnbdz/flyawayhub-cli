package cmd

import (
	"bufio"
	"context"
	"flyawayhub-cli/auth"
	appConfig "flyawayhub-cli/config"
	"flyawayhub-cli/logging"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// PromptForLogin prompts the user for login credentials and attempts to authenticate.
func PromptForLogin(authService auth.Service, logger logging.Logger) error {
	ctx := context.Background() // Or context.TODO() if unsure

	fmt.Print("Enter Username: ")
	username, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		logger.Error("Error reading username", map[string]interface{}{"error": err})
		return err
	}
	username = strings.TrimSpace(username)

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(0)
	fmt.Println() // Move to a new line after password input.
	if err != nil {
		logger.Error("Error reading password", map[string]interface{}{"error": err})
		return err
	}
	password := string(bytePassword)

	// Use authService to perform login
	accessToken, err := authService.Login(ctx, username, password)
	if err != nil {
		logger.Error("Login failed", map[string]interface{}{"error": err})
		return err
	}

	if err := FetchOrganizationInfo(accessToken, "json"); err != nil {
		logger.Error("Failed to fetch organization info: %v\n", err)
		// Decide how you want to handle this error. For now, just printing the error.
	}

	logger.Info("Login successful")
	return nil
}

// NewLoginCommand updated to use loginProcess within the Run function.
func NewLoginCommand(authService auth.Service, logger logging.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login will sign you in to " + appConfig.AppName,
		RunE: func(cmd *cobra.Command, args []string) error {
			return PromptForLogin(authService, logger)
		},
	}
	return cmd
}

func InitCommands(root *cobra.Command, authService auth.Service, logger logging.Logger) {
	// Instantiate the login command with required dependencies
	loginCmd := NewLoginCommand(authService, logger)

	// Add the login command to the root command
	root.AddCommand(loginCmd)
}
