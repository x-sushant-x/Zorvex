package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "zorvex",
		Short: "Zorvex CLI. Manage your services using this CLI application easily.",
	}

	getAgentURL = &cobra.Command{
		Use:   "get_agent_url",
		Short: CommandsList["get_agent_url"].Short,
		Long:  CommandsList["get_agent_url"].Long,
		Run:   GetAgentURL,
	}
)

func Execute() error {
	rootCmd.AddCommand(getAgentURL)

	return rootCmd.Execute()
}
