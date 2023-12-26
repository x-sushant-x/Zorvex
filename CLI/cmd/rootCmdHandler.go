package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "zorvex",
		Short: "Zorvex CLI. Manage your services using this CLI application easily.",
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

	getService = &cobra.Command{
		Use:   "get_service",
		Short: CommandsList["get_service"].Short,
		Long:  CommandsList["get_service"].Long,
		Run:   GetService,
	}
)

func Execute() error {
	rootCmd.AddCommand(getAllServices)
	rootCmd.AddCommand(getAllDownServices)
	rootCmd.AddCommand(getService)

	return rootCmd.Execute()
}
