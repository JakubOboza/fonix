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

		rid, err := cmd.Flags().GetString("requestid")

		if err != nil {
			fmt.Println(err)
			return
		}

		params := &client.SmsParams{
			Originator: originator,
			Body:       body,
			Numbers:    numbers,
			Dummy:      dummy,
			RequestID:  rid,
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

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		out, err := client.Output(result, format)

		fmt.Println(out)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(sendsmsCmd)

	sendsmsCmd.Flags().StringP("apikey", "k", "", "apikey for the service eg: --apikey=live:myKey123456XYZ")
	sendsmsCmd.Flags().StringP("body", "b", "", "body of the text message eg: --body=hello")
	sendsmsCmd.Flags().StringP("originator", "o", "", "originator of the message eg: --originator=889988")
	sendsmsCmd.Flags().StringP("numbers", "n", "", "numbers to send the sms to eg: --numbers=4474123456778")
	sendsmsCmd.Flags().StringP("dummy", "d", "", "dummy flag yes or no, you can skip it. dummy=yes will make fonix not send sms but mock respond eg: --dummy=yes")
	sendsmsCmd.Flags().StringP("requestid", "r", "", "setup request id for the request, this will be resend in DR. Max 80 chars eg: --requestid=RAA12233222")

	sendsmsCmd.MarkFlagRequired("body")
	sendsmsCmd.MarkFlagRequired("originator")
	sendsmsCmd.MarkFlagRequired("numbers")

	sendsmsCmd.Flags().StringP("baseurl", "u", "", "base url of fonix/sonar that isnt default. --baseurl=https://sonar.fonix.io")

	sendsmsCmd.Flags().StringP("format", "f", "", "output format eg: --format=json")
}
