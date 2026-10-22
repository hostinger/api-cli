package products

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateDigitalCmd = &cobra.Command{
	Use:               "create-digital <store_id>",
	Short:             "Create digital product",
	Long:              "Create a published digital product with a single variant and an optional external download link.",
	Example:           "  hostinger ecommerce products create-digital store_01J8Z5F8W9K8M4A7B3C2D1E0FG --name 'My eBook' --price 999",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createDigitalBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceCreateDigitalProductV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateDigitalCmd, utils.StringArg("store_id"))
	CreateDigitalCmd.Flags().StringP("currency", "", "", "ISO 4217 currency code. Defaults to the store's default currency when omitted.")
	CreateDigitalCmd.Flags().StringP("description", "", "", "The product description.")
	CreateDigitalCmd.Flags().StringP("download-url", "", "", "Optional external download link delivered to the customer after purchase.")
	CreateDigitalCmd.Flags().StringP("name", "", "", "The product name.")
	CreateDigitalCmd.Flags().IntP("price", "", 0, "Price in the smallest currency unit (e.g. cents). Must be positive.")
	CreateDigitalCmd.MarkFlagRequired("name")
	CreateDigitalCmd.MarkFlagRequired("price")
	CreateDigitalCmd.RegisterFlagCompletionFunc("currency", cobra.NoFileCompletions)
	CreateDigitalCmd.RegisterFlagCompletionFunc("description", cobra.NoFileCompletions)
	CreateDigitalCmd.RegisterFlagCompletionFunc("download-url", cobra.NoFileCompletions)
	CreateDigitalCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateDigitalCmd.RegisterFlagCompletionFunc("price", cobra.NoFileCompletions)
}

func createDigitalBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("currency") {
		v, _ := cmd.Flags().GetString("currency")
		body["currency"] = v
	}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	if cmd.Flags().Changed("download-url") {
		v, _ := cmd.Flags().GetString("download-url")
		body["download_url"] = v
	}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	priceVal, _ := cmd.Flags().GetInt("price")
	body["price"] = priceVal
	return body, nil
}
