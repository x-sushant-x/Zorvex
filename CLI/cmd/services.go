package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/sushant102004/Zorvex/CLI/services"
	"github.com/sushant102004/Zorvex/CLI/utils"
)

var (
	headerFmt = color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt = color.New(color.FgYellow).SprintfFunc()
)

func GetAllServices(cmd *cobra.Command, args []string) {
	fmt.Println()

	tbl := table.New("ID", "Name", "HTTP Method", "IP", "Port", "Balancing Method", "Last Sync", "Endpoint", "Status")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	services, err := services.GetAllServices()
	if err != nil {
		utils.Error(err.Error())
	}

	for _, service := range services {
		tbl.AddRow(service.ID, service.Name, service.HTTPMethod, service.IPAddress, service.Port, service.LoadBalancingMethod, service.LastSyncTime, service.Endpoint, service.Status)
	}

	tbl.Print()
	fmt.Println()

}

func GetAllDownServices(cmd *cobra.Command, args []string) {
	tbl := table.New("ID", "Name", "HTTP Method", "IP", "Port", "Balancing Method", "Last Sync", "Endpoint", "Status")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	services, err := services.GetAllDownServices()
	if err != nil {
		utils.Error(err.Error())
	}

	if len(services) == 0 {
		log.Info().Msg("All services are up! 😊")
		return
	}

	for _, service := range services {
		tbl.AddRow(service.ID, service.Name, service.HTTPMethod, service.IPAddress, service.Port, service.LoadBalancingMethod, service.LastSyncTime, service.Endpoint, service.Status)
	}

	tbl.Print()
	fmt.Println()
}
