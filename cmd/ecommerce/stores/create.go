package stores

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:               "create",
	Short:             "Create store",
	Long:              "Create a new store for your account.\n\nA primary sales channel is created alongside the store.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --sales-channel\n    {\n      \"external_id\": \"string\",\n      \"type\": \"custom\"\n    }",
	Example:           "  hostinger ecommerce stores create",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceCreateStoreV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(CreateCmd)
	CreateCmd.Flags().StringP("company-email", "", "", "")
	CreateCmd.Flags().StringP("company-name", "", "", "")
	CreateCmd.Flags().StringP("country-code", "", "", "ISO 3166-1 alpha-2 country code.")
	CreateCmd.Flags().StringP("language", "", "", "ISO 639-1 language code.")
	CreateCmd.Flags().StringP("name", "", "", "")
	CreateCmd.Flags().StringP("sales-channel", "", "", "(JSON, @file or -; shape in the help above)")
	CreateCmd.RegisterFlagCompletionFunc("company-email", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("company-name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("country-code", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("language", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("sales-channel", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
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
		salesChannelVal, _ := cmd.Flags().GetString("sales-channel")
		salesChannelValJSON, err := utils.JSONValue(salesChannelVal, "sales-channel")
		if err != nil {
			return nil, err
		}
		body["sales_channel"] = salesChannelValJSON
	}
	return body, nil
}
