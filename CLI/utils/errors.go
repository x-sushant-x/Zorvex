package utils

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func TypoError(command cobra.Command, err error) {
	fmt.Println()
	log.Error().Msgf("ERROR")
	fmt.Println(command.Short)
	fmt.Println(command.Long)
	fmt.Println("Error: " + err.Error())
	fmt.Println()
	os.Exit(-1)

}

func Error(err string) {
	fmt.Println()
	log.Error().Msgf("ERROR")
	fmt.Println(err)
	fmt.Println()
	os.Exit(-1)
}
