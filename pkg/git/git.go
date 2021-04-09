package git

import (
	"context"
	"fmt"
	"time"

	"github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var branch = "refs/heads/amendments"

// Clone peforms a git clone into a directory called ./tmp/
func Clone(repo, d, token, user string) error {
	url := "https://github.com/" + repo
	dir := d + "/" + repo

	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		// Progress: os.Stdout,
		URL: url,
		Auth: &http.BasicAuth{
			Username: user,
			Password: token,
		},
	})
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	b := plumbing.ReferenceName(branch)

	err = w.Checkout(&git.CheckoutOptions{
		Create: true,
		Force:  false,
		Branch: b,
	})
	if err != nil {
		return err
	}

	return nil
}

func Add(dir string) error {
	r, err := git.PlainOpen(dir)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	err = w.AddWithOptions(&git.AddOptions{
		All:  true,
		Glob: ".",
	})
	if err != nil {
		return err
	}

	commit, err := w.Commit("message", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "jasonBirchall",
			Email: "jason.birchall@digital.Justice.gov.uk",
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	_, err = r.CommitObject(commit)
	if err != nil {
		return err
	}

	err = r.Push(&git.PushOptions{
		RemoteName: "amendments",
	})

	return nil
}

func PullRequest(token, command, org, repo string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	newPR := &github.NewPullRequest{
		Title:               github.String("Execute over each directory"),
		Head:                github.String("amendments"),
		Base:                github.String("main"),
		Body:                github.String("Walked each path and executed: " + command),
		MaintainerCanModify: github.Bool(true),
	}

	// fmt.Println(token, command, org, repo, newPR)
	pr, _, err := client.PullRequests.Create(ctx, org, repo, newPR)
	if err != nil {
		fmt.Println("shiiit:", err)
	}

	fmt.Printf("PR created: %s\n", pr.GetHTMLURL())
	return nil
}
