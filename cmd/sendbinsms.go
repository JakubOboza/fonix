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

// sendbinsmsCmd represents the sendbinsms command
var sendbinsmsCmd = &cobra.Command{
	Use:   "sendbinsms",
	Short: "sendbinsms sends binary configured sms via fonix api",
	Long:  `endbinsms sends binary configured sms via fonix api`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("API_KEY")

		flagApiKey, err := cmd.Flags().GetString("url")
		if err == nil && flagApiKey != "" {
			apiKey = flagApiKey
		}

		// params

		originator, err := cmd.Flags().GetString("originator")

		if err != nil {
			fmt.Println(err)
			return
		}

		binBody, err := cmd.Flags().GetString("binbody")

		if err != nil {
			fmt.Println(err)
			return
		}

		numbers, err := cmd.Flags().GetString("numbers")

		if err != nil {
			fmt.Println(err)
			return
		}

		dummy, err := cmd.Flags().GetString("dummy")

		if err != nil {
			fmt.Println(err)
			return
		}

		params := &client.SmsBinParams{
			Originator: originator,
			BinBody:    binBody,
			Numbers:    numbers,
			Dummy:      dummy,
		}

		fonixClient := client.New(apiKey)

		newBaseUrl, _ := cmd.Flags().GetString("baseurl")
		if newBaseUrl != "" {
			fonixClient.SetBaseURL(newBaseUrl)
		}

		result, err := fonixClient.SendBinSms(context.Background(), params)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("======Success======")
		fmt.Println("Guid: ", result.TxGuid)
		fmt.Println("Numbers: ", result.Numbers)
		fmt.Println("Parts: ", result.SmsParts)
		fmt.Println("Encoding: ", result.Encoding)
	},
}

func init() {
	rootCmd.AddCommand(sendbinsmsCmd)

	sendbinsmsCmd.Flags().StringP("apikey", "k", "", "apikey for the service eg: --apikey=live:myKey123456XYZ")
	sendbinsmsCmd.Flags().StringP("binbody", "b", "", "body of the text message eg: --binbody=C024A3A5905E195B081180800991C6106DA620420620A22C4986166184289B526204")
	sendbinsmsCmd.Flags().StringP("originator", "o", "", "originator of the message eg: --originator=889988")
	sendbinsmsCmd.Flags().StringP("numbers", "n", "", "numbers to send the sms to eg: --numbers=4474123456778")
	sendbinsmsCmd.Flags().StringP("dummy", "d", "", "dummy flag yes or no, you can skip it. dummy=yes will make fonix not send sms but mock respond eg: --dummy=yes")

	sendbinsmsCmd.MarkFlagRequired("binbody")
	sendbinsmsCmd.MarkFlagRequired("originator")
	sendbinsmsCmd.MarkFlagRequired("numbers")

	sendbinsmsCmd.Flags().StringP("baseurl", "u", "", "base url of fonix/sonar that isnt default. --baseurl=https://sonar.fonix.io")
}
