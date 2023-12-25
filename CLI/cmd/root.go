package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sushant102004/Zorvex/CLI/utils"
)

var (
	rootCmd = &cobra.Command{
		Use:   "set_url",
		Short: utils.CommandsList["set_url"].Short,
		Long:  utils.CommandsList["set_url"].Long,
		Run:   SetAgent,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
