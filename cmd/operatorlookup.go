/*
Copyright © 2022 Jakub Oboza <jakub.oboza@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/JakubOboza/fonix/client"
	"github.com/spf13/cobra"
)

// operatorlookupCmd represents the operatorlookup command
var operatorlookupCmd = &cobra.Command{
	Use:   "operatorlookup",
	Short: "lookup operator for a given number using fonix api",
	Long:  `lookup operator for a given number using fonix api`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("API_KEY")

		flagApiKey, err := cmd.Flags().GetString("url")
		if err == nil && flagApiKey != "" {
			apiKey = flagApiKey
		}

		// params

		number, err := cmd.Flags().GetString("number")

		if err != nil {
			fmt.Println(err)
			return
		}

		dummy, err := cmd.Flags().GetString("dummy")

		if err != nil {
			fmt.Println(err)
			return
		}

		params := &client.OperatorLookupParams{
			Number: number,
			Dummy:  dummy,
		}

		fonixClient := client.New(apiKey)

		newBaseUrl, _ := cmd.Flags().GetString("baseurl")
		if newBaseUrl != "" {
			fonixClient.SetBaseURL(newBaseUrl)
		}

		result, err := fonixClient.OperatorLookup(context.Background(), params)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("======Operator Lookup Result======")
		fmt.Println("Operator: ", result.Operator)
		fmt.Println("Mnc: ", result.Mnc)
		fmt.Println("Mcc: ", result.Mcc)
	},
}

func init() {
	rootCmd.AddCommand(operatorlookupCmd)

	operatorlookupCmd.Flags().StringP("apikey", "k", "", "apikey for the service eg: --apikey=live:myKey123456XYZ")
	operatorlookupCmd.Flags().StringP("number", "n", "", "number to have operator lookup made for eg: --number=4474123456778")
	operatorlookupCmd.Flags().StringP("dummy", "d", "", "dummy flag yes or no, you can skip it. dummy=yes will make fonix not send sms but mock respond eg: --dummy=yes")

	operatorlookupCmd.MarkFlagRequired("number")

	operatorlookupCmd.Flags().StringP("baseurl", "u", "", "base url of fonix/sonar that isnt default. --baseurl=https://sonar.fonix.io")

}
