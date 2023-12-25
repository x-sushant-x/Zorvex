package cmd

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/sushant102004/Zorvex/CLI/utils"
)

func SetAgent(cmd *cobra.Command, args []string) {
	if len(args) < 3 {
		utils.TypoError(utils.CommandsList["set_url"])
	}

	if err := os.Setenv("AGENT_URL", args[2]); err != nil {
		utils.Error(err.Error())
	}

	log.Info().Msg("DONE")
}
