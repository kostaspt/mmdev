package jira

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/gosimple/slug"
	"github.com/spf13/cobra"

	"mmdev/internal/config"
)

type BranchRunner struct {
	config *config.Config
}

func NewBranchCmd(cfg *config.Config) *cobra.Command {
	r := &BranchRunner{cfg}
	cmd := &cobra.Command{
		Use:  "branch <ticket url or key>",
		RunE: r.Run,
	}

	return cmd
}

func (r *BranchRunner) Run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("ticket url or key is required")
	}

	c, err := NewClient(r.config)
	if err != nil {
		return err
	}

	key, err := r.parseID(args[0])
	if err != nil {
		return err
	}

	i, _, err := c.Issue.Get(key, nil)
	if err != nil {
		return err
	}

	bn := fmt.Sprintf("%s-%s", i.Key, r.sanitizeTitle(i.Fields.Summary))

	osCmd := exec.Command("git", "switch", "-c", bn)
	if err = osCmd.Run(); err != nil {
		return err
	}

	return nil
}

func (r *BranchRunner) parseID(ident string) (string, error) {
	id := regexp.MustCompile(`(MM-[0-9]+)`).FindString(ident)
	if id == "" {
		return "", errors.New("invalid issue identifier")
	}
	return id, nil
}

func (r *BranchRunner) sanitizeTitle(title string) string {
	t := slug.Make(
		stopwords.CleanString(
			regexp.MustCompile(`(?i)(?:\A|\s)p[0-9](?::|\s|\z)`).ReplaceAllString(title, ""),
			"en", true,
		),
	)

	words := strings.Split(t, "-")

	finalWords := make([]string, 0)
	for _, word := range words {
		if word == "" {
			continue
		}

		if regexp.MustCompile("^p[0-9]$").MatchString(word) {
			continue
		}

		if len(finalWords) == 5 {
			break
		}

		finalWords = append(finalWords, word)
	}

	res := strings.Join(finalWords, "-")

	// Limit max size of the branch to 255 characters as it's enforced by git.
	if len(res) >= 255 {
		return res[:255]
	}

	return res
}
