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

	getAllServices = &cobra.Command{
		Use:   "get_all_services",
		Short: CommandsList["get_all_services"].Short,
		Long:  CommandsList["get_all_services"].Long,
		Run:   GetAllServices,
	}

	getAllDownServices = &cobra.Command{
		Use:   "get_all_down_services",
		Short: CommandsList["get_all_down_services"].Short,
		Long:  CommandsList["get_all_down_services"].Long,
		Run:   GetAllDownServices,
	}
)

func Execute() error {
	rootCmd.AddCommand(getAgentURL)
	rootCmd.AddCommand(getAllServices)
	rootCmd.AddCommand(getAllDownServices)

	return rootCmd.Execute()
}
