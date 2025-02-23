package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple todo application for the command line.",
	Long:  "A simple CLI tool written to help organise my tasks and mainain a track of the important steps that I need to make. Initially this will be a command line only tool. But Jira integration would be kinda cool... In the future though. Right now I'm not even using Jira, I am however, using Monday.com... Hmmmm",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
