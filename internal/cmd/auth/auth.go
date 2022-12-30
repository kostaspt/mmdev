package auth

import (
	"github.com/spf13/cobra"

	"mmdev/internal/cmd/auth/jira"
	"mmdev/internal/config"
)

func NewCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "auth <command>",
	}

	cmd.AddCommand(jira.NewCmd(cfg))

	return cmd
}
