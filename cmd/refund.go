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

// refundCmd represents the refund command
var refundCmd = &cobra.Command{
	Use:   "refund",
	Short: "refund charges from command line",
	Long:  `refund charges from command line. This tool can help you make manual refunds`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("API_KEY")

		flagApiKey, err := cmd.Flags().GetString("url")
		if err == nil && flagApiKey != "" {
			apiKey = flagApiKey
		}

		// params
		chargeGUID, err := cmd.Flags().GetString("chargeguid")

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

		params := &client.RefundParams{
			Numbers:    numbers,
			Dummy:      dummy,
			ChargeGuid: chargeGUID,
			RequestID:  rid,
		}

		fonixClient := client.New(apiKey)

		newBaseUrl, _ := cmd.Flags().GetString("baseurl")
		if newBaseUrl != "" {
			fonixClient.SetBaseRefundURL(newBaseUrl)
		}

		result, err := fonixClient.Refund(context.Background(), params)

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
	rootCmd.AddCommand(refundCmd)

	refundCmd.Flags().StringP("apikey", "k", "", "apikey for the service eg: --apikey=live:myKey123456XYZ")
	refundCmd.Flags().StringP("numbers", "n", "", "numbers to send the sms to eg: --numbers=4474123456778")
	refundCmd.Flags().StringP("dummy", "d", "", "dummy flag yes or no, you can skip it. dummy=yes will make fonix not send sms but mock respond eg: --dummy=yes")
	refundCmd.Flags().StringP("requestid", "r", "", "setup request id for the request, this will be resend in DR. Max 80 chars eg: --requestid=RAA12233222")
	refundCmd.Flags().StringP("chargeguid", "g", "", "setup request id for the request, this will be resend in DR. Max 80 chars eg: --chargeguid=RAA12233222")

	refundCmd.MarkFlagRequired("numbers")

	refundCmd.Flags().StringP("baseurl", "u", "", "base url of fonix/refund that isnt default. --baseurl=https://refund.fonix.io")
	refundCmd.Flags().StringP("format", "f", "", "output format eg: --format=json")
}
