package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sushant102004/Zorvex/CLI/cmd"
	"github.com/sushant102004/Zorvex/CLI/services"
	"github.com/sushant102004/Zorvex/CLI/utils"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		utils.Error(err.Error())
	}

	agentURL := os.Getenv("AGENT_URL")
	if agentURL == "" {
		utils.Error(errors.New("unable to get AGENT_URL from enviroment variables").Error())
	}

	services.SetAgentURL(agentURL)
	cmd.Execute()
}
