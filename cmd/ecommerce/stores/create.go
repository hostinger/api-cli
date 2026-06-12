package stores

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create store",
	Long:  "Create a new store for your account.\n\nA primary sales channel is created alongside the store.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceCreateStoreV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("company-email", "", "", "")
	CreateCmd.Flags().StringP("company-name", "", "", "")
	CreateCmd.Flags().StringP("country-code", "", "", "ISO 3166-1 alpha-2 country code.")
	CreateCmd.Flags().StringP("language", "", "", "ISO 639-1 language code.")
	CreateCmd.Flags().StringP("name", "", "", "")
	CreateCmd.Flags().StringP("sales-channel", "", "", " (JSON)")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("company-email") {
		v, _ := cmd.Flags().GetString("company-email")
		body["company_email"] = v
	}
	if cmd.Flags().Changed("company-name") {
		v, _ := cmd.Flags().GetString("company-name")
		body["company_name"] = v
	}
	if cmd.Flags().Changed("country-code") {
		v, _ := cmd.Flags().GetString("country-code")
		body["country_code"] = v
	}
	if cmd.Flags().Changed("language") {
		v, _ := cmd.Flags().GetString("language")
		body["language"] = v
	}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("sales-channel") {
		v, _ := cmd.Flags().GetString("sales-channel")
		body["sales_channel"] = utils.JSONValue(v, "sales-channel")
	}
	return body
}
