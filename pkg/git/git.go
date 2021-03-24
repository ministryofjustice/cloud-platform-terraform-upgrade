package git

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Clone peforms a git clone into a directory called ./tmp/
func Clone(repo, d string) error {
	url := "https://github.com/" + repo
	dir := d + "/" + repo

	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	branch := "refs/heads/amendments"
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

func Commit(dir string) error {
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

	commit, err := w.Commit("Executed command", &git.CommitOptions{
		All: true,
		Author: (&object.Signature{
			// Name:  "jasonbirchall",
			// Email: "jason.birchall@digital.justice.gov.uk",
			When: time.Now(),
		}),
	})
	if err != nil {
		fmt.Println("commit execute failed: ", err)
	}

	_, err = r.CommitObject(commit)
	if err != nil {
		return err
	}

	return nil
}
