package jira

import (
	"fmt"

	"github.com/spf13/cobra"

	"mmdev/internal/config"
)

type MeRunner struct {
	config *config.Config
}

func NewMeCmd(cfg *config.Config) *cobra.Command {
	r := &MeRunner{cfg}
	cmd := &cobra.Command{
		Use:  "me",
		RunE: r.Run,
	}

	return cmd
}

func (r *MeRunner) Run(cmd *cobra.Command, args []string) error {
	c, err := NewClient(r.config)
	if err != nil {
		return err
	}

	u, _, err := c.User.GetSelf()
	if err != nil {
		return err
	}

	fmt.Println("Hey, " + u.DisplayName + "!")

	return nil
}
