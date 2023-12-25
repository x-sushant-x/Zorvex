package cmd

import (
	"github.com/spf13/cobra"
)

var CommandsList = map[string]cobra.Command{
	"set_agent_url": {
		Short: "Usage: zorvex set_agent_url = https://agenturl.com",
		Long:  "This command is used to set the URL for agent.",
	},

	"get_agent_url": {
		Short: "Usage: zorvex get_agent_url",
		Long:  "Returns the URL for the agent",
	},
}
