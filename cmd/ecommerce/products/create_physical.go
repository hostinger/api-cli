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

var CreatePhysicalCmd = &cobra.Command{
	Use:               "create-physical <store_id>",
	Short:             "Create physical product",
	Long:              "Create a published physical product with a single variant priced in the store currency.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --seo-settings\n    {\n      \"description\": \"Soft combed cotton, unisex fit.\",\n      \"slug\": \"blue-t-shirt\",\n      \"title\": \"Blue T-Shirt | My Shop\"\n    }",
	Example:           "  hostinger ecommerce products create-physical store_01J8Z5F8W9K8M4A7B3C2D1E0FG --name 'My Product' --price 1000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createPhysicalBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceCreatePhysicalProductV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreatePhysicalCmd, utils.StringArg("store_id"))
	CreatePhysicalCmd.Flags().StringP("currency", "", "", "ISO 4217 currency code. Defaults to the store's default currency when omitted.")
	CreatePhysicalCmd.Flags().StringP("description", "", "", "The product description.")
	CreatePhysicalCmd.Flags().StringP("name", "", "", "The product name.")
	CreatePhysicalCmd.Flags().IntP("price", "", 0, "Price in the smallest currency unit (e.g. cents). Must be positive.")
	CreatePhysicalCmd.Flags().StringP("seo-settings", "", "", "Search-engine settings. Omitted fields are generated from the name and description; a slug already in use gets a numeric suffix. Send at least one field. (JSON, @file or -; shape in the help above)")
	CreatePhysicalCmd.MarkFlagRequired("name")
	CreatePhysicalCmd.MarkFlagRequired("price")
	CreatePhysicalCmd.RegisterFlagCompletionFunc("currency", cobra.NoFileCompletions)
	CreatePhysicalCmd.RegisterFlagCompletionFunc("description", cobra.NoFileCompletions)
	CreatePhysicalCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreatePhysicalCmd.RegisterFlagCompletionFunc("price", cobra.NoFileCompletions)
	CreatePhysicalCmd.RegisterFlagCompletionFunc("seo-settings", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func createPhysicalBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("currency") {
		v, _ := cmd.Flags().GetString("currency")
		body["currency"] = v
	}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	priceVal, _ := cmd.Flags().GetInt("price")
	body["price"] = priceVal
	if cmd.Flags().Changed("seo-settings") {
		seoSettingsVal, _ := cmd.Flags().GetString("seo-settings")
		seoSettingsValJSON, err := utils.JSONValue(seoSettingsVal, "seo-settings")
		if err != nil {
			return nil, err
		}
		body["seo_settings"] = seoSettingsValJSON
	}
	return body, nil
}
