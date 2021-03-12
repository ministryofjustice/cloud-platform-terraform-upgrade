package git

import (
	"github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
)

// Clone peforms a git clone into a directory called ./tmp/
func Clone(r, d string) error {
	url := "https://github.com/" + r
	dir := d + "/" + r

	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return err
	}

	return nil
}
