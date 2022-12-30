package auth

import "github.com/spf13/cobra"

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "auth <command>",
	}

	return cmd
}
