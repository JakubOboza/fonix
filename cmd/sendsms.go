/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sendsmsCmd represents the sendsms command
var sendsmsCmd = &cobra.Command{
	Use:   "sendsms",
	Short: "send sms via fonix from command line",
	Long: `Use this subcommand to send bulk/free sms to a number 
	using your API keys via fonix gateway`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sendsms called")
	},
}

func init() {
	rootCmd.AddCommand(sendsmsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendsmsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendsmsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
