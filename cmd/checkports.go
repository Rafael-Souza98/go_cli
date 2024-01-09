/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/go_cli/utils"
	"github.com/spf13/cobra"
)

// checkportsCmd represents the checkports command
var checkportsCmd = &cobra.Command{
	Use:   "checkports",
	Short: "",
	Long: ``,
	
	Run: func(cmd *cobra.Command, args []string) {
		host, _:= cmd.Flags().GetString("h")
		if host == "" {
			cmd.Println("Inform the host")
			return
		}
		
		ports, _:= cmd.Flags().GetString("p")
		if ports == "" {
			cmd.Println("Inform the ports list")
			return
		}
		portList:= strings.Split(ports, ",")

		utils.CheckPorts(host, portList)
	},
}

func init() {
	rootCmd.AddCommand(checkportsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	checkportsCmd.PersistentFlags().String("h", "", "Host has been validated")
	checkportsCmd.PersistentFlags().String("p", "", "Ports list")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkportsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
