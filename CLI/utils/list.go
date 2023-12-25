package utils

import (
	"github.com/spf13/cobra"
)

var CommandsList = map[string]cobra.Command{
	"set_url": {
		Short: "Usage: zorvex set_url = https://agenturl.com",
		Long:  "This command is used to set the URL for agent.",
	},
}
