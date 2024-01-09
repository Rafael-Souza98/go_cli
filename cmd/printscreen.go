/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (

	"github.com/go_cli/utils"
	"github.com/spf13/cobra"
)

// printscreenCmd represents the printscreen command
var printscreenCmd = &cobra.Command{
	Use:   "printscreen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url:= cmd.Flag("site").Value.String()
		utils.ChromeScreenShot(url, 100)
	},
}

func init() {
	rootCmd.AddCommand(printscreenCmd)
	printscreenCmd.PersistentFlags().String("site", "www.google.com", "Execute a screenshot of site")

}
