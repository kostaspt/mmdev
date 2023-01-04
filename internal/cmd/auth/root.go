package auth

import (
	"github.com/spf13/cobra"

	"mmdev/internal/config"
)

func NewAuthRootCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "auth <command>",
	}

	cmd.AddCommand(NewJiraCmd(cfg))

	return cmd
}
