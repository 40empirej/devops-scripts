package devops_scripts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type GitHubResponse struct {
	HTMLURL string `json:"html_url"`
}

func GetGitHubRepoInfo(repo string) (*GitHubResponse, error) {
	githubAPIEndpoint := fmt.Sprintf("https://api.github.com/repos/%s", repo)
	req, err := http.NewRequest("GET", githubAPIEndpoint, nil)
	if err!= nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err!= nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
		return nil, err
	}

	var githubResponse GitHubResponse
	err = json.Unmarshal(body, &githubResponse)
	if err!= nil {
		return nil, err
	}

	return &githubResponse, nil
}

func GetGitHubToken() string {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN environment variable is not set")
	}
	return token
}

func GetGitHubUsername() string {
	username := os.Getenv("GITHUB_USERNAME")
	if username == "" {
		log.Fatal("GITHUB_USERNAME environment variable is not set")
	}
	return username
}

func GetGitHubRepoName(repo string) string {
	repoParts := strings.Split(repo, "/")
	if len(repoParts)!= 2 {
		log.Fatal("Invalid repository format")
	}
	return repoParts[1]
}