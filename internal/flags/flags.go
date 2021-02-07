package flags

import (
	"errors"

	"github.com/rs/zerolog"
	"github.com/sethvargo/go-githubactions"
)

// Reporter reporter flags
type Reporter struct {
	CoverageReport   string
	Verbose          bool
	GithubToken      string `env:"GITHUB_TOKEN"`
	GithubEventName  string `env:"GITHUB_EVENT_NAME"`
	GithubEventPath  string `env:"GITHUB_EVENT_PATH"`
	GithubOwner      string `env:"GITHUB_REPOSITORY_OWNER"`
	GithubActor      string `env:"GITHUB_ACTOR"`
	GithubRepository string `env:"GITHUB_REPOSITORY"`
}

// LogLevel configure the zerolog log level
func (rep Reporter) LogLevel() zerolog.Level {
	if rep.Verbose {
		return zerolog.DebugLevel
	}

	return zerolog.InfoLevel
}

func (rep Reporter) ValidateToken() (string, error) {
	githubToken := githubactions.GetInput("GITHUB_TOKEN")
	if githubToken != "" {
		return githubToken, nil
	}

	if rep.GithubToken != "" {
		return rep.GithubToken, nil
	}

	return "", errors.New("missing required GITHUB_TOKEN")
}

// func (rep Reporter) GetRepo() (string, string, error) {
// 	var owner string

// 	if owner = rep.GitHubOwner; owner == "" {
// 		owner = rep.GitHubActor
// 	}

// 	if rep.GitHubRepository == "" {

// 	}

// 	repoParts := rep.GitHubRepository
// }
