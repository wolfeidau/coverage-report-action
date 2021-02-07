package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/google/go-github/github"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/wolfeidau/coverage-report-action/internal/cobertura"
	"github.com/wolfeidau/coverage-report-action/internal/flags"
	"github.com/wolfeidau/coverage-report-action/internal/ptr"
	"golang.org/x/oauth2"
)

var (
	cfg     = new(flags.Reporter)
	version = "unknown"
)

func main() {
	kong.Parse(cfg,
		kong.Vars{"version": version}, // bind a var for version
	)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.Kitchen}).Level(cfg.LogLevel())

	token, err := cfg.ValidateToken()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to validate config")
	}

	reportData, err := ioutil.ReadFile(cfg.CoverageReport)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read coverage report")
	}

	report, err := cobertura.ParseCoverageReport(bytes.NewBuffer(reportData))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read coverage report")
	}

	buf := new(bytes.Buffer)

	err = cobertura.RunTemplate(buf, report)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run coverage template")
	}

	log.Debug().Msg("processing event")

	event, err := ioutil.ReadFile(cfg.GithubEventPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load event")
	}

	var pr *github.PullRequestEvent

	switch cfg.GithubEventName {
	case "pull_request", "pull_request_target":

		err = json.Unmarshal(event, &pr)
	default:
		log.Warn().Msgf("unable to process event :%s", cfg.GithubEventName)
		return
	}

	log.Info().Int("pr", ptr.ToInt(pr.Number)).Msg("processing pull request")

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	comment, _, err := client.PullRequests.CreateComment(ctx, ptr.ToString(pr.Repo.Owner.Login), ptr.ToString(pr.Repo.Name), ptr.ToInt(pr.Number), &github.PullRequestComment{
		Body: ptr.String(buf.String()),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create comment")
	}

	log.Info().Int64("id", ptr.ToInt64(comment.ID)).Msg("comment created")
}
