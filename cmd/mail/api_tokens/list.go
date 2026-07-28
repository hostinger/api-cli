package api_tokens

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List API tokens",
	Long:  "Retrieve a paginated list of\n[Hostinger Email API](https://api.mail.hostinger.com/) tokens across\nall your mail orders, optionally filtered by order. Plaintext tokens\nare never included; they are returned only when a token is created.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailListAPITokensV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("order-id", "", "", "Filter tokens by order resource ID. Single value or comma-separated list.")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listParams(cmd *cobra.Command) *client.MailListAPITokensV1Params {
	params := &client.MailListAPITokensV1Params{}
	if cmd.Flags().Changed("order-id") {
		v, _ := cmd.Flags().GetString("order-id")
		params.OrderId = &v
	}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	return params
}
