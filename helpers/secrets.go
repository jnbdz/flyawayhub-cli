package helpers

import (
	"flyawayhub-cli/logging"
	"fmt"
	"github.com/zalando/go-keyring"
)

// Assuming logger is an initialized global variable of type logging.Logger
var logger logging.Logger

func init() {
	// Initialize your logger here. This is a simplified example.
	// In a real-world scenario, you might want to inject the logger as a dependency.
	logger = logging.NewZapLogger()
}

type SecretManager interface {
	SetSecret(service, account, secret string) error
	GetSecret(service, account string) (string, error)
	DeleteSecret(service, account string) error
}

// SetSecret stores a secret in the system's keyring for the specified service and account.
func SetSecret(service, account, secret string) error {
	err := keyring.Set(service, account, secret)
	if err != nil {
		logger.Error("Failed to set secret", "service", service, "account", account, "error", err)
		return fmt.Errorf("failed to set secret for %s/%s: %w", service, account, err)
	}
	logger.Info("Secret set successfully", "service", service, "account", account)
	return nil
}

// GetSecret retrieves a secret from the system's keyring for the specified service and account.
func GetSecret(service, account string) (string, error) {
	retrievedSecret, err := keyring.Get(service, account)
	if err != nil {
		logger.Error("Failed to get secret", "service", service, "account", account, "error", err)
		return "", fmt.Errorf("failed to get secret for %s/%s: %w", service, account, err)
	}
	logger.Info("Secret retrieved successfully", "service", service, "account", account)
	return retrievedSecret, nil
}

// DeleteSecret removes a secret from the system's keyring for the specified service and account.
func DeleteSecret(service, account string) error {
	err := keyring.Delete(service, account)
	if err != nil {
		logger.Error("Failed to delete secret", "service", service, "account", account, "error", err)
		return fmt.Errorf("failed to delete secret for %s/%s: %w", service, account, err)
	}
	logger.Info("Secret deleted successfully", "service", service, "account", account)
	return nil
}
