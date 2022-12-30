package main

import (
	"fmt"
	"os"

	"mmdev/internal/cmd/root"
	"mmdev/internal/config"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	defer func(cfg *config.Config) {
		if dErr := cfg.Save(); dErr != nil {
			fmt.Fprintln(os.Stderr, dErr)
		}
	}(cfg)

	rootCmd := root.NewCmd(cfg)

	return rootCmd.Execute()
}
