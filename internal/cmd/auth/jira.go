package auth

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"mmdev/internal/config"
)

type JiraRunner struct {
	config *config.Config
}

func NewJiraCmd(cfg *config.Config) *cobra.Command {
	r := &JiraRunner{cfg}
	cmd := &cobra.Command{
		Use:  "jira",
		RunE: r.Run,
	}

	return cmd
}

func (r *JiraRunner) Run(cmd *cobra.Command, args []string) error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your username: ")
	scanner.Scan()
	r.config.Auth.Jira.Username = scanner.Text()

	fmt.Print("Enter your API token: ")
	scanner.Scan()
	r.config.Auth.Jira.ApiToken = scanner.Text()

	return nil
}
