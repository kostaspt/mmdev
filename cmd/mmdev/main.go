package main

import (
	"fmt"
	"os"

	"mmdev/internal/cmd/root"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	rootCmd := root.NewCmd()
	return rootCmd.Execute()
}
