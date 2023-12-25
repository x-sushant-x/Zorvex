package cmd

import (
	"github.com/spf13/cobra"
)

var CommandsList = map[string]cobra.Command{
	"get_all_services": {
		Short: "Usage: zorvex get_all_services",
		Long:  "Prints are the registered services in tabular format.",
	},

	"get_agent_url": {
		Short: "Usage: zorvex get_agent_url",
		Long:  "Returns the URL for the agent.",
	},

	"get_all_down_services": {
		Short: "Usage: zorvex get_all_down_services",
		Long:  "Returns all the services that are down.",
	},
}
