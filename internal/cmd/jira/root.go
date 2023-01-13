package jira

import (
	"github.com/spf13/cobra"

	"mmdev/internal/config"
)

func NewJiraRootCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "jira <command>",
	}

	cmd.AddCommand(NewBranchCmd(cfg))
	cmd.AddCommand(NewMeCmd(cfg))

	return cmd
}
