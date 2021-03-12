package main

import (
	"log"
	"os"

	"github.com/ministryofjustice/cloud-platform-terraform-upgrade/pkg/get"
	"github.com/ministryofjustice/cloud-platform-terraform-upgrade/pkg/git"
)

func main() {
	// pass GH oauth2 token to variable
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatalln("Unauthorised: No token present, please map your env vars.")
	}

	// get list of repositories
	repos, err := get.Repositories("ministryofjustice", token, "cloud-platform-terraform")
	if err != nil {
		log.Fatalln(err)
	}

	// Clone repository locally
	for _, repo := range repos {
		err = git.Clone(repo)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
