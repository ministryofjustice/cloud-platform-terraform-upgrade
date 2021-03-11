package get

import (
	"context"
	"strings"

	"github.com/google/go-github/v33/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"golang.org/x/oauth2"                    // oauth token for GH auth
)

// Repositories gets certain repositories in a given GitHub organisation matching an argument s.
// It take three arguments:
// - the name of an organisation, i.e. "google"
// - an oauth2 token, created by a GitHub user https://github.com/google/go-github#authentication
// - and a pattern to match the repository name, i.e. "go"
func Repositories(n, t, s string) (l []string, err error) {
	// Create authenticated client to avoid API rate limiting.
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: t},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{
		Type:        "public",
		Sort:        "full_name",
		ListOptions: github.ListOptions{PerPage: 10},
	}

	// Becuase of the potential number of org repositories pagination is added.
	// Warning: this can take a while if the org contains a number of repositories.
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, n, opt)
		if err != nil {
			return nil, err
		}

		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	// Loop over all repositories and grab only matching repositories.
	for _, repo := range allRepos {
		c := string(*repo.FullName)
		if strings.Contains(c, s) {
			l = append(l, c)
		}
	}

	return
}
