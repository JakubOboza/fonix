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

// adultverifyCmd represents the adultverify command
var adultverifyCmd = &cobra.Command{
	Use:   "adultverify",
	Short: "adult verification api check",
	Long:  `sync/async adult vertification tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("API_KEY")

		flagApiKey, err := cmd.Flags().GetString("url")
		if err == nil && flagApiKey != "" {
			apiKey = flagApiKey
		}

		async, err := cmd.Flags().GetBool("async")

		if err != nil {
			fmt.Println(err)
			return
		}

		// params

		networkretry, err := cmd.Flags().GetString("networkretry")

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

		params := &client.AvParams{
			NetworkRetry: networkretry,
			Numbers:      numbers,
			Dummy:        dummy,
		}

		fonixClient := client.New(apiKey)

		newBaseUrl, _ := cmd.Flags().GetString("baseurl")

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if async {

			if newBaseUrl != "" {
				fonixClient.SetBaseURL(newBaseUrl)
			}

			result, err := fonixClient.AdultVerify(context.Background(), params)

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

		} else {

			if newBaseUrl != "" {
				fonixClient.SetBaseAvSoloURL(newBaseUrl)
			}

			result, err := fonixClient.AvSolo(context.Background(), params)

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

		}
	},
}

func init() {
	rootCmd.AddCommand(adultverifyCmd)

	adultverifyCmd.Flags().StringP("apikey", "k", "", "apikey for the service eg: --apikey=live:myKey123456XYZ")
	adultverifyCmd.Flags().StringP("numbers", "n", "", "numbers to send the sms to eg: --numbers=4474123456778")
	adultverifyCmd.Flags().StringP("networkretry", "r", "", "network retry eg: --networkretry=no")
	adultverifyCmd.Flags().StringP("dummy", "d", "", "dummy flag yes or no, you can skip it. dummy=yes will make fonix not send sms but mock respond eg: --dummy=yes")

	adultverifyCmd.Flags().BoolP("async", "a", false, "trigger async call --async=true")

	sendsmsCmd.MarkFlagRequired("networkretry")
	sendsmsCmd.MarkFlagRequired("numbers")

	adultverifyCmd.Flags().StringP("baseurl", "u", "", "base url of fonix/sonar that isnt default. --baseurl=https://sonar.fonix.io")
	adultverifyCmd.Flags().StringP("format", "f", "", "output format eg: --format=json")
}
