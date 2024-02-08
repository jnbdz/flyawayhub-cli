package helpers

import (
	"fmt"
	"github.com/zalando/go-keyring"
)

func setSecret(service, account, secret string) {
	err := keyring.Set(service, account, secret)
	if err != nil {
		fmt.Println("Failed to set secret:", err)
		return
	}
	return
}

func getSecret(service, account string) string {
	retrievedSecret, err := keyring.Get(service, account)
	if err != nil {
		fmt.Println("Failed to get secret:", err)
		return ""
	}

	return retrievedSecret
}

func deleteSecret(service, user string) {
	err := keyring.Delete(service, user)
	if err != nil {
		fmt.Println("Failed to get secret:", err)
		return
	}
	return
}
