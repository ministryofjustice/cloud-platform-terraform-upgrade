package git

import (
	"fmt"

	"github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
)

func Clone(r string) error {
	url := "https://github.com/" + r
	fmt.Println(url)
	_, err := git.PlainClone("tmp/"+r, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return err
	}

	return nil
}
