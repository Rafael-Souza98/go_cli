/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/go_cli/utils"
	"github.com/spf13/cobra"
)

// whoisCmd represents the whois command
var whoisCmd = &cobra.Command{
	Use:   "whois",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host:= cmd.Flag("host").Value.String()
		if host == ""{
			log.Fatal("Inform the host")
			return
		}
		fmt.Println(utils.ConsultDomain(host))
	},
}

func init() {
	rootCmd.AddCommand(whoisCmd)

	
	whoisCmd.PersistentFlags().String("host", "", "Consult a domain")

	
}
