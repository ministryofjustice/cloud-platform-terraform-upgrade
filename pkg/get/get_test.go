package get

import (
	"os"
	"testing"
)

// TestRepositories tests the Repositories function in the get package returns the correct number of
// repositories when called.
func TestRepositories(t *testing.T) {
	// Pass GH oauth2 token to variable
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		t.Errorf("No token present, test failed.")
	}

	// Get the repositories for the i3 org and pattern match "Go".
	// "Go" was just the simplist name I could find in an org that only had double figure repositories.
	repos, err := Repositories("TheAlgorithms", "Go", token)
	if err != nil {
		t.Errorf("Failed to get repositories")
	}
	// There should only be one repositories in this org called "Go".
	if len(repos) < 1 {
		t.Errorf("Failed to get repositories")
	}
}
