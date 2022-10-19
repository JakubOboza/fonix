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

// sendwappushCmd represents the sendwappush command
var sendwappushCmd = &cobra.Command{
	Use:   "sendwappush",
	Short: "sendwappush sends wap api push via fonix api",
	Long:  `command line tool to send wap push via fonix gateway v2 api`,
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

		pushTitle, err := cmd.Flags().GetString("title")

		if err != nil {
			fmt.Println(err)
			return
		}

		pushLink, err := cmd.Flags().GetString("link")

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

		params := &client.SmsWapParams{
			Originator: originator,
			PushTitle:  pushTitle,
			PushLink:   pushLink,
			Numbers:    numbers,
			Dummy:      dummy,
		}

		fonixClient := client.New(apiKey)

		newBaseUrl, _ := cmd.Flags().GetString("baseurl")
		if newBaseUrl != "" {
			fonixClient.SetBaseURL(newBaseUrl)
		}

		result, err := fonixClient.SendWapPush(context.Background(), params)

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
	rootCmd.AddCommand(sendwappushCmd)

	sendwappushCmd.Flags().StringP("apikey", "k", "", "apikey for the service eg: --apikey=live:myKey123456XYZ")
	sendwappushCmd.Flags().StringP("title", "t", "", "push title eg: --title=hello")
	sendwappushCmd.Flags().StringP("link", "l", "", "push link eg: --link=http://google.com/link/to/push")
	sendwappushCmd.Flags().StringP("originator", "o", "", "originator of the message eg: --originator=889988")
	sendwappushCmd.Flags().StringP("numbers", "n", "", "numbers to send the sms to eg: --numbers=4474123456778")
	sendwappushCmd.Flags().StringP("dummy", "d", "", "dummy flag yes or no, you can skip it. dummy=yes will make fonix not send sms but mock respond eg: --dummy=yes")

	sendwappushCmd.MarkFlagRequired("title")
	sendwappushCmd.MarkFlagRequired("link")
	sendwappushCmd.MarkFlagRequired("originator")
	sendwappushCmd.MarkFlagRequired("numbers")

	sendwappushCmd.Flags().StringP("baseurl", "u", "", "base url of fonix/sonar that isnt default. --baseurl=https://sonar.fonix.io")
	sendwappushCmd.Flags().StringP("format", "f", "", "output format eg: --format=json")
}
