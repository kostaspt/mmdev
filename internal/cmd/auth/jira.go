package auth

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"mmdev/internal/config"
)

type JiraCmdOpts struct {
	config   *config.Config
	Username string
	ApiToken string
}

func NewJiraCmd(cfg *config.Config) *cobra.Command {
	opts := &JiraCmdOpts{
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

			return runJira(opts)
		},
	}

	return cmd
}

func runJira(opts *JiraCmdOpts) error {
	opts.config.Auth.Jira.Username = opts.Username
	opts.config.Auth.Jira.ApiToken = opts.ApiToken
	return nil
}
