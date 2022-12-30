package jira

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"mmdev/internal/config"
)

type Opts struct {
	config   *config.Config
	Username string
	ApiToken string
}

func NewCmd(cfg *config.Config) *cobra.Command {
	opts := &Opts{
		config: cfg,
	}

	cmd := &cobra.Command{
		Use: "jira",
		RunE: func(cmd *cobra.Command, args []string) error {
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Print("Enter your username: ")
			scanner.Scan()
			opts.Username = scanner.Text()

			fmt.Print("Enter your API token: ")
			scanner.Scan()
			opts.ApiToken = scanner.Text()

			return run(opts)
		},
	}

	return cmd
}

func run(opts *Opts) error {
	opts.config.Auth.Jira.Username = opts.Username
	opts.config.Auth.Jira.ApiToken = opts.ApiToken
	return nil
}
