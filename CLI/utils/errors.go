package utils

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func TypoError(command cobra.Command) {
	fmt.Println()
	log.Error().Msgf("ERROR")
	fmt.Printf("Usage: %s\n", command.Short)
	fmt.Printf("Description: %s\n", command.Long)
	fmt.Println()
	os.Exit(-1)

}

func Error(err string) {
	fmt.Println()
	log.Error().Msgf("ERROR")
	fmt.Println(err)
}
