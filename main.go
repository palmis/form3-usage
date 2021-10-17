package main

import (
	"fmt"
	"log"

	fm3a "github.com/palmis/go-form3-accounts"
)

func main() {
	client := fm3a.NewClient("http://localhost:8080")
	country := "GB"

	// Create account
	account := fm3a.AccountData{
		Type:           "accounts",
		ID:             "d8ba96f3-5f6a-4ae1-a4e3-54a04e02f918",
		OrganisationID: "7263da44-4509-4357-9dc6-b8f8eca1a680",
		Attributes: &fm3a.AccountAttributes{
			Country: &country,
			Name:    []string{"test", "template"},
		},
	}
	if err := client.Create(account); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account created. ID: %s\n", account.ID)

	// Get account
	created, err := client.Fetch(account.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account fetched. ID: %s Version: %d\n", created.ID, *created.Version)

	// Delete versioned account
	if err := client.Delete(created.ID, int32(*created.Version)); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account deleted. ID: %s Version: %d\n", created.ID, *created.Version)
}
