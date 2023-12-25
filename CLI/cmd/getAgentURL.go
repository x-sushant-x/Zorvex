package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sushant102004/Zorvex/CLI/utils"
)

func GetAgentURL(cmd *cobra.Command, args []string) {
	agentURL := os.Getenv("AGENT_URL")
	if agentURL == "" {
		utils.Error("AGENT_URL not set in environment variable. Please set using set_url=<URL> command")
	}
	fmt.Println("Agent URL: ", agentURL)
}
