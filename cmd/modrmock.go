/*
Copyright © 2022 Jakub Oboza <jakub.oboza@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/JakubOboza/fonix/client"
	"github.com/spf13/cobra"
)

// modrmockCmd represents the modrmock command
var modrmockCmd = &cobra.Command{
	Use:   "modrmock",
	Short: "mock handler for mos and drs testing",
	Long:  `MO & DR mock handler for integration testing`,
	Run: func(cmd *cobra.Command, args []string) {

		port, err := cmd.Flags().GetInt("port")

		if err != nil {
			fmt.Println(err)
			return
		}

		client.StartMockDrMoHandler(port)
	},
}

func init() {
	rootCmd.AddCommand(modrmockCmd)

	modrmockCmd.Flags().IntP("port", "p", 8090, "mock port eg: --port=6677")
}
