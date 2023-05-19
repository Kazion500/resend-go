package main

import (
	"fmt"
	"os"

	"github.com/drish/resend-go"
)

func main() {

	apiKey := os.Getenv("RESEND_API_KEY")

	client := resend.NewClient(apiKey)

	// Create API Key
	params := &resend.CreateDomainRequest{
		Name: "exampledomain.com",
	}

	domain, err := client.Domains.Create(params)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created Domain entry id: " + domain.Id)
	fmt.Println("Status: " + domain.Status)

	for _, record := range domain.Records {
		fmt.Printf("%v\n", record)
	}

	// List
	domains, err := client.Domains.List()
	if err != nil {
		panic(err)
	}
	fmt.Printf("You have %d domains in your project\n", len(domains.Data))

	// Verify
	verified, err := client.Domains.Verify(domain.Id)
	if err != nil {
		panic(err)
	}
	if verified {
		println("verified domain id: " + domain.Id)
	} else {
		println("could not verify domain id: " + domain.Id)
	}

	// Remove
	removed, err := client.Domains.Remove(domain.Id)
	if err != nil {
		panic(err)
	}
	if removed {
		println("removed domain id: " + domain.Id)
	} else {
		println("could not remove domain id: " + domain.Id)
	}
}