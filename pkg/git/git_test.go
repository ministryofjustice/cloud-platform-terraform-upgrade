package git

import (
	"fmt"
	"os"
	"testing"
)

// TestClone clones a local GitHub repository down and checks to see if
// it exists.
func TestClone(t *testing.T) {
	repository := "TheAlgorithms/Go"
	dir := "./"

	err := Clone(repository, dir)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Unable to clone the repository: %s\n", err)
	}

	if _, err := os.Stat(dir + repository); os.IsNotExist(err) {
		t.Errorf("Repository doesn't exist: %s\n", err)
	}

	err = os.RemoveAll(dir + repository)
	if err != nil {
		t.Errorf("Unable to remove repository: %s\n", err)
	}

}
