package api

import (
	appConfig "flyawayhub-cli/config"
	"flyawayhub-cli/logging"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
)

const (
	maxRedirects = 5
)

var logger logging.Logger

// LoginCommand Define a struct to hold dependencies for the login command.
type LoginCommand struct {
	logger logging.Logger
	// Add other dependencies here
}

// setHeaders sets common headers for the HTTP request.
func setHeaders(req *http.Request, accessToken string) {
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}

// handleHttpResponseStatus evaluates HTTP response status codes and logs or returns appropriate errors.
// It is designed to complement custom redirection handling and provide specific logging or error handling based on the response status.
func handleHttpResponseStatus(statusCode int, status string) error {
	switch {
	case statusCode >= 100 && statusCode <= 199:
		// Informational responses
		logger.Info("Informational response", "status", status)
		return nil // Typically no action needed for informational responses

	case statusCode >= 200 && statusCode <= 299:
		// Successful responses
		logger.Info("Successful response", "status", status)
		return nil

	case statusCode >= 300 && statusCode <= 399:
		// Redirection messages are handled by the http.Client's CheckRedirect function
		return nil // No additional action needed here if redirection is already handled

	case statusCode == 401 || statusCode == 403:
		// Client error responses: Unauthorized or Forbidden
		logger.Warn("Client error response", "status", status)
		if err := authCallback(); err != nil {
			return fmt.Errorf("authentication error: %s, %w", status, err)
		}
		return fmt.Errorf("authentication error: %s", status)

	case statusCode == 418:
		// 418 I'm a teapot 🫖
		logger.Info("I'm a teapot 🫖", "status", status)
		return fmt.Errorf("teapot 🫖")

	case statusCode >= 400 && statusCode <= 499:
		// Client error responses
		logger.Error("Client error response", "status", status)
		return fmt.Errorf("client error: %s", status)

	case statusCode >= 500 && statusCode <= 599:
		// Server error responses
		logger.Error("Server error response", "status", status)
		return fmt.Errorf("server error: %s", status)

	default:
		// Unexpected status code
		logger.Error("Unexpected response status", "status", status)
		return fmt.Errorf("unexpected response: %s", status)
	}
}

// Get performs a GET request and returns the raw response body.
func Get(endpoint, accessToken string) ([]byte, error) {
	reqURL := appConfig.APIEndpoint(endpoint)

	client := &http.Client{
		// CheckRedirect is called before following an HTTP redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Log the redirect
			logger.Info("Redirect detected",
				zap.String("from", via[len(via)-1].URL.String()),
				zap.String("to", req.URL.String()),
			)

			// If too many redirects occurred, stop following them
			if len(via) >= maxRedirects {
				logger.Info("Too many redirects",
					zap.String("from", via[len(via)-1].URL.String()),
					zap.String("to", req.URL.String()),
					zap.String("count", strconv.Itoa(maxRedirects)),
				)
				return fmt.Errorf("stopped after %d redirects", maxRedirects)
			}
			return nil // Continue following redirects
		},
	}
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		logger.Error("Error creating request", "error", err)
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	setHeaders(req, accessToken)

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error making request", "error", err)
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body", "error", err)
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	err = handleHttpResponseStatus(resp.StatusCode, resp.Status)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Post performs a POST request and returns the raw response body.
func Post(logger logging.Logger, endpoint, accessToken string) ([]byte, error) {
	reqURL := appConfig.APIEndpoint(endpoint)

	client := &http.Client{}
	req, err := http.NewRequest("POST", reqURL, nil)
	if err != nil {
		logger.Error("Error creating request", "error", err)
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	setHeaders(req, accessToken)

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error making request", "error", err)
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		if cerr := resp.Body.Close(); cerr != nil {
			logger.Error("Failed to close response body", "error", cerr)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body", "error", err)
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return body, nil
}
