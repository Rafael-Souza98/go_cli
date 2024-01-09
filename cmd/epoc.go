/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/go_cli/utils"
	"github.com/spf13/cobra"
)

// epocCmd represents the epoc command
var epocCmd = &cobra.Command{
	Use:   "epoc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		epoch:= cmd.Flag("c").Value.String()
		fmt.Println(utils.ConvertEpocToHuman(epoch))
	},
}

func init() {
	rootCmd.AddCommand(epocCmd)

	epocCmd.PersistentFlags().String("c", "", "Convert epoch to datetime")

	
}
