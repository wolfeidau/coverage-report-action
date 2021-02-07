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
	GitHubToken      string `env:"GITHUB_TOKEN"`
	GitHubEventName  string `env:"GITHUB_EVENT_NAME"`
	GitHubEventPath  string `env:"GITHUB_EVENT_PATH"`
	GitHubOwner      string `env:"GITHUB_REPOSITORY_OWNER"`
	GitHubActor      string `env:"GITHUB_ACTOR"`
	GitHubRepository string `env:"GITHUB_REPOSITORY"`
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

	if rep.GitHubToken != "" {
		return rep.GitHubToken, nil
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
