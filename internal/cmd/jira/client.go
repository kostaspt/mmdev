package jira

import (
	"errors"

	"github.com/andygrunwald/go-jira"

	"mmdev/internal/config"
)

func NewClient(cfg *config.Config) (*jira.Client, error) {
	if cfg.Auth.Jira.Email == "" || cfg.Auth.Jira.APIToken == "" {
		return nil, errors.New("jira credentials not set")
	}

	tp := jira.BasicAuthTransport{
		Username: cfg.Auth.Jira.Email,
		Password: cfg.Auth.Jira.APIToken,
	}

	return jira.NewClient(tp.Client(), "https://mattermost.atlassian.net")
}
