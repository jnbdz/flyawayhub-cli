package helpers

import (
	"fmt"
	"github.com/zalando/go-keyring"
)

type SecretManager interface {
	SetSecret(service, account, secret string) error
	GetSecret(service, account string) (string, error)
	DeleteSecret(service, account string) error
}

func SetSecret(service, account, secret string) {
	err := keyring.Set(service, account, secret)
	if err != nil {
		fmt.Println("Failed to set secret:", err)
		return
	}
	return
}

func GetSecret(service, account string) string {
	retrievedSecret, err := keyring.Get(service, account)
	if err != nil {
		fmt.Println("Failed to get secret:", err)
		return ""
	}

	return retrievedSecret
}

func DeleteSecret(service, user string) {
	err := keyring.Delete(service, user)
	if err != nil {
		fmt.Println("Failed to get secret:", err)
		return
	}
	return
}
