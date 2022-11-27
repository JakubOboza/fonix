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

// chargemobileCmd represents the chargemobile command
var chargemobileCmd = &cobra.Command{
	Use:   "chargemobile",
	Short: "charge mobile number",
	Long: `charge mobile number from command line. 
	This method call charges the mobile phone account of one or many phone numbers.
	The charge is only ever attempted once. If unsuccessful we do not re-attempt to charge in the same request unless you specify smsfallback=yes, in which case the charge is attempted again through premium SMS.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("API_KEY")

		flagApiKey, err := cmd.Flags().GetString("url")
		if err == nil && flagApiKey != "" {
			apiKey = flagApiKey
		}

		// params
		amount, err := cmd.Flags().GetInt("currency")

		if err != nil {
			fmt.Println(err)
			return
		}

		ttl, err := cmd.Flags().GetInt("ttl")

		if err != nil {
			fmt.Println(err)
			return
		}

		currency, err := cmd.Flags().GetString("currency")

		if err != nil {
			fmt.Println(err)
			return
		}

		numbers, err := cmd.Flags().GetString("numbers")

		if err != nil {
			fmt.Println(err)
			return
		}

		body, err := cmd.Flags().GetString("body")

		if err != nil {
			fmt.Println(err)
			return
		}

		chargeDescription, err := cmd.Flags().GetString("description")

		if err != nil {
			fmt.Println(err)
			return
		}

		chargeSilent, err := cmd.Flags().GetString("chargesilent")

		if err != nil {
			fmt.Println(err)
			return
		}

		smsfallback, err := cmd.Flags().GetString("fallback")

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

		params := &client.ChargeMobileParams{
			Numbers:           numbers,
			Amount:            amount,
			Currency:          currency,
			RequestID:         rid,
			ChargeDescription: chargeDescription,
			TimeToLive:        ttl,
			ChargeSilent:      chargeSilent,
			Body:              body,
			Dummy:             dummy,
			SmsFallback:       smsfallback,
		}

		fonixClient := client.New(apiKey)

		newBaseUrl, _ := cmd.Flags().GetString("baseurl")
		if newBaseUrl != "" {
			fonixClient.SetBaseRefundURL(newBaseUrl)
		}

		result, err := fonixClient.ChargeMobile(context.Background(), params)

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
	rootCmd.AddCommand(chargemobileCmd)

	chargemobileCmd.Flags().StringP("apikey", "k", "", "apikey for the service eg: --apikey=live:myKey123456XYZ")
	chargemobileCmd.Flags().StringP("numbers", "n", "", "numbers to send the sms to eg: --numbers=4474123456778")
	chargemobileCmd.Flags().StringP("dummy", "d", "", "dummy flag yes or no, you can skip it. dummy=yes will make fonix not send sms but mock respond eg: --dummy=yes")
	chargemobileCmd.Flags().StringP("requestid", "r", "", "setup request id for the request, this will be resend in DR. Max 80 chars eg: --requestid=RAA12233222")
	chargemobileCmd.Flags().StringP("currency", "c", "", "setup currency for the charge eg: --currency=GBP")
	chargemobileCmd.Flags().IntP("amount", "a", 0, "setup amountto charge eg: --amount=25")
	chargemobileCmd.Flags().IntP("ttl", "t", 10, "setup time to live for the charge eg: --ttl=25")
	chargemobileCmd.Flags().StringP("fallback", "j", "no", "setup force payment fallback in case of fail eg: --fallback=no")
	chargemobileCmd.Flags().StringP("chargesilent", "z", "no", "setup charging silent for the request eg: --chargesilent=no")
	chargemobileCmd.Flags().StringP("body", "b", "", "body for the charge request eg: --body=Thanks")
	chargemobileCmd.Flags().StringP("description", "p", "", "charge description eg: --description=Thanks")

	chargemobileCmd.MarkFlagRequired("numbers")

	chargemobileCmd.Flags().StringP("baseurl", "u", "", "base url of fonix/refund that isnt default. --baseurl=https://refund.fonix.io")
	chargemobileCmd.Flags().StringP("format", "f", "", "output format eg: --format=json")
}
