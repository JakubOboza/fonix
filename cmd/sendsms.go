/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/JakubOboza/fonix/client"
	"github.com/spf13/cobra"
)

// sendsmsCmd represents the sendsms command
var sendsmsCmd = &cobra.Command{
	Use:   "sendsms",
	Short: "send sms via fonix from command line",
	Long: `Use this subcommand to send bulk/free sms to a number 
	using your API keys via fonix gateway`,
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

		body, err := cmd.Flags().GetString("body")

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

		params := &client.SmsParams{
			Originator: originator,
			Body:       body,
			Numbers:    numbers,
			Dummy:      dummy,
		}

		fonixClient := client.New(apiKey)

		newBaseUrl, _ := cmd.Flags().GetString("baseurl")
		if newBaseUrl != "" {
			fonixClient.SetBaseURL(newBaseUrl)
		}

		result, err := fonixClient.SendSms(context.Background(), params)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("======Success======")
		fmt.Println("Guid: ", result.SuccessData.TxGuid)
		fmt.Println("Numbers: ", result.SuccessData.Numbers)
		fmt.Println("Parts: ", result.SuccessData.SmsParts)
		fmt.Println("Encoding: ", result.SuccessData.Encoding)
	},
}

func init() {
	rootCmd.AddCommand(sendsmsCmd)

	sendsmsCmd.Flags().StringP("apikey", "k", "", "apikey for the service eg: --apikey=live:myKey123456XYZ")
	sendsmsCmd.Flags().StringP("body", "b", "", "body of the text message eg: --body=hello")
	sendsmsCmd.Flags().StringP("originator", "o", "", "originator of the message eg: --originator=889988")
	sendsmsCmd.Flags().StringP("numbers", "n", "", "numbers to send the sms to eg: --numbers=4474123456778")
	sendsmsCmd.Flags().StringP("dummy", "d", "", "dummy flag yes or no, you can skip it. dummy=yes will make fonix not send sms but mock respond eg: --dummy=yes")

	sendsmsCmd.MarkFlagRequired("body")
	sendsmsCmd.MarkFlagRequired("originator")
	sendsmsCmd.MarkFlagRequired("numbers")

	sendsmsCmd.Flags().StringP("baseurl", "u", "", "base url of fonix/sonar that isnt default. --baseurl=https://sonar.fonix.io")
}
