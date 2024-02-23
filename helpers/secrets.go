package helpers

import (
	"flyawayhub-cli/logging"
	"fmt"
	"github.com/zalando/go-keyring"
)

type SecretManager interface {
	Set(service, account, secret string) error
	Get(service, account string) (string, error)
	Delete(service, account string) error
}

type KeyringSecretManager struct {
	logger logging.Logger
}

func NewKeyringSecretManager(logger logging.Logger) *KeyringSecretManager {
	return &KeyringSecretManager{logger: logger}
}

// Set stores a secret in the system's keyring for the specified service and account.
func (ksm *KeyringSecretManager) Set(service, account, secret string) error {
	err := keyring.Set(service, account, secret)
	if err != nil {
		ksm.logger.Error("Failed to set secret", "service", service, "account", account, "error", err)
		return fmt.Errorf("failed to set secret for %s/%s: %w", service, account, err)
	}
	ksm.logger.Info("Secret set successfully", "service", service, "account", account)

	return nil
}

// Get retrieves a secret from the system's keyring for the specified service and account.
func (ksm *KeyringSecretManager) Get(service, account string) (string, error) {
	retrievedSecret, err := keyring.Get(service, account)
	if err != nil {
		ksm.logger.Error("Failed to get secret", "service", service, "account", account, "error", err)
		return "", fmt.Errorf("failed to get secret for %s/%s: %w", service, account, err)
	}
	ksm.logger.Info("Secret retrieved successfully", "service", service, "account", account)
	return retrievedSecret, nil
}

// Delete removes a secret from the system's keyring for the specified service and account.
func (ksm *KeyringSecretManager) Delete(service, account string) error {
	err := keyring.Delete(service, account)
	if err != nil {
		ksm.logger.Error("Failed to delete secret", "service", service, "account", account, "error", err)
		return fmt.Errorf("failed to delete secret for %s/%s: %w", service, account, err)
	}
	ksm.logger.Info("Secret deleted successfully", "service", service, "account", account)
	return nil
}
