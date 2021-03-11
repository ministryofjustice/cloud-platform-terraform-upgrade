package main

import (
	"fmt"
	"log"
	"os"

	"example.com/hello/pkg/get"
)

func main() {
	// pass GH oauth2 token to variable
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorised: No token present, please map your env vars.")
	}

	// get list of repositories
	repos, err := get.Repositories("ministryofjustice", token, "cloud-platform-terraform")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(repos)

}
