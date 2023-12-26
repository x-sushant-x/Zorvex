package cmd

import (
	"encoding/json"
	"errors"
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

func tableGenerator(tableHeadings ...string) table.Table {
	fmt.Println()
	interfaceHeadings := make([]interface{}, len(tableHeadings))
	for i, v := range tableHeadings {
		interfaceHeadings[i] = v
	}

	tbl := table.New(interfaceHeadings...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return tbl
}

func GetAllServices(cmd *cobra.Command, args []string) {
	services, err := services.GetAllServices()
	if err != nil {
		utils.Error(err.Error())
	}

	tbl := tableGenerator("ID", "Name", "HTTP Method", "IP", "Port", "Balancing Method", "Last Sync", "Endpoint", "Status")

	for _, service := range services {
		tbl.AddRow(service.ID, service.Name, service.HTTPMethod, service.IPAddress, service.Port, service.LoadBalancingMethod, service.LastSyncTime, service.Endpoint, service.Status)
	}

	tbl.Print()
	fmt.Println()

}

func GetAllDownServices(cmd *cobra.Command, args []string) {
	services, err := services.GetAllDownServices()
	if err != nil {
		utils.Error(err.Error())
	}

	tbl := tableGenerator("ID", "Name", "HTTP Method", "IP", "Port", "Balancing Method", "Last Sync", "Endpoint", "Status")

	for _, service := range services {
		tbl.AddRow(service.ID, service.Name, service.HTTPMethod, service.IPAddress, service.Port, service.LoadBalancingMethod, service.LastSyncTime, service.Endpoint, service.Status)
	}

	tbl.Print()
	fmt.Println()
}

func GetService(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		utils.Error(errors.New("invalid command").Error())
	}

	service, err := services.GetSingleService(args[1])
	if err != nil {
		utils.Error(err.Error())
	}

	jsonData, err := json.MarshalIndent(service, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	log.Info().Msgf("SUCCESS")
	fmt.Println(string(jsonData))
}
