package root

import (
	"github.com/spf13/cobra"

	"mmdev/internal/cmd/auth"
	"mmdev/internal/config"
)

func NewCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "mmdev <command>",
	}

	cmd.AddCommand(auth.NewCmd(cfg))

	return cmd
}
