package root

import (
	"github.com/spf13/cobra"

	"mmdev/internal/cmd/auth"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "mmdev <command>",
	}

	cmd.AddCommand(auth.NewCmd())

	return cmd
}
