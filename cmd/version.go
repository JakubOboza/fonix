/*
Copyright © 2022 Jakub Oboza <jakub.oboza@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/JakubOboza/fonix/client"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "client version",
	Long:  `client version. This shows the CLI/Lib client version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(client.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
