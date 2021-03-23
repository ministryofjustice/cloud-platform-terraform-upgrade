package git

import (
	"github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"github.com/go-git/go-git/v5/plumbing"
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
