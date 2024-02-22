// Package auth provides interfaces and implementations for handling
// authentication within the application. It includes functionality
// for initiating re-authentication flows, managing authentication
// tokens, and validating user credentials.
package auth

import (
	"context"
	appConfig "flyawayhub-cli/config"
	"flyawayhub-cli/logging"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

// Service interface defines the methods required for authentication.
type Service interface {
	Login(
		ctx context.Context,
		username,
		password string,
		triggerFetchOrganizationInfo func(accessToken string) error) error
}

// authService implements the Service interface.
type authService struct {
	logger logging.Logger
}

// NewAuthService returns a new instance of an authService.
func NewAuthService(logger logging.Logger) Service {
	return &authService{logger: logger}
}

// Login handles the user login process.
func (s *authService) Login(
	ctx context.Context,
	username,
	password string,
	triggerFetchOrganizationInfo func(accessToken string) error) error {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(appConfig.AWSCognitoRegion))
	if err != nil {
		s.logger.Error("Unable to load SDK config", "error", err)
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
		s.logger.Error("authentication failed: %v\n", err)
		return err
	}

	// Fetch and update session with organization info
	// FetchOrganizationInfo(*authResp.AuthenticationResult.AccessToken, "json")
	if err := triggerFetchOrganizationInfo(*authResp.AuthenticationResult.AccessToken); err != nil {
		s.logger.Error("Failed to fetch organization info: %v\n", err)
		// Decide how you want to handle this error. For now, just printing the error.
	}

	s.logger.Info("Login successful")
	return nil
}
