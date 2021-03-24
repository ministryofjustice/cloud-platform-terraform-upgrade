package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ministryofjustice/cloud-platform-terraform-upgrade/pkg/get"
	"github.com/ministryofjustice/cloud-platform-terraform-upgrade/pkg/git"
	"github.com/ministryofjustice/cloud-platform-terraform-upgrade/pkg/utils"
)

const temp = "tmp/"

var (
	org, repo, command string
	commit             bool
)

func init() {
	// Initialise flags and parse.
	flag.StringVar(&org, "o", "ministryofjustice", "Name of the GitHub organisation")
	flag.StringVar(&repo, "r", "cloud-platform-terraform", "Pattern of the repository to match")
	flag.StringVar(&command, "c", "ls -latr", "command to execute")
	flag.BoolVar(&commit, "commit", true, "whether you want to commit changes")

	flag.Parse()
}

func main() {
	// Authenticate with an oauth2 token from GitHub.
	a := utils.Config{
		Token: os.Getenv("GITHUB_AUTH_TOKEN"),
	}
	if a.Token == "" {
		log.Fatalln("Unauthorised: No user or token present")
	}

	// Get list of repositories
	repos, err := get.Repositories(org, repo, a.Token)
	if err != nil {
		log.Fatalln(err)
	}

	// Clone repository locally
	var dirs []string
	for _, repo := range repos {
		err = git.Clone(repo, temp)
		if err != nil {
			log.Fatalln(err)
		} else {
			dirs = append(dirs, repo)
		}
	}

	// loop over each repository
	for _, dir := range dirs {
		path := temp + dir
		err := utils.Execute(path, command)
		if err != nil {
			log.Fatalln(err)
		}

		// if commit is set to true, then commit and pr
		if commit {
			err := git.Commit(path)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			fmt.Println("commit isn't selected")
		}
	}
}
