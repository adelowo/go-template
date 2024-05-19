package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	if err := Execute(); err != nil {
		log.Fatal(err)
	}
}

func Execute() error {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "App is a golang template",
	}

	// rootCmd.AddCommand(startAPICommand())

	return rootCmd.Execute()
}
