package product_variants

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
	Use:               "create <store_id> <product_id>",
	Short:             "Create a product variant",
	Long:              "Add a variant to a product along one or more option dimensions (e.g. Size, Color). Options\nmissing from the product are created automatically; provide a value for every option the\nproduct already has. Prices are integers in the smallest currency unit and default to the\nstore currency. Returns the created variant.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --options\n    [\n      {\n        \"name\": \"Size\",\n        \"value\": \"M\"\n      }\n    ]\n\n  --prices\n    [\n      {\n        \"amount\": 1999,\n        \"currency\": \"usd\",\n        \"sale_amount\": 1499\n      }\n    ]",
	Example:           "  hostinger ecommerce product-variants create store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG --options '[{\"name\":\"Size\",\"value\":\"M\"}]'",
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
			r, err := api.Request().EcommerceCreateAProductVariantV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd, utils.StringArg("store_id"), utils.StringArg("product_id"))
	CreateCmd.Flags().IntP("inventory-quantity", "", 0, "Units in stock. Defaults to 0.")
	CreateCmd.Flags().BoolP("manage-inventory", "", false, "Whether stock is tracked for this variant. Defaults to false.")
	CreateCmd.Flags().StringP("options", "", "", "Option name/value pairs that distinguish this variant, e.g. [{name: Size, value: M}]. Options missing from the product are created; provide a value for every option the product already has. (JSON, @file or -; shape in the help above)")
	CreateCmd.Flags().StringP("prices", "", "", "Prices per currency. Amounts are integers in the smallest currency unit. A free item is amount: 0. (JSON, @file or -; shape in the help above)")
	CreateCmd.Flags().StringP("sku", "", "", "The variant SKU.")
	CreateCmd.Flags().StringP("title", "", "", "The variant title. Defaults to the option values joined with ' / ' (e.g. 'Red / L').")
	CreateCmd.MarkFlagRequired("options")
	CreateCmd.RegisterFlagCompletionFunc("inventory-quantity", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("manage-inventory", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("options", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateCmd.RegisterFlagCompletionFunc("prices", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateCmd.RegisterFlagCompletionFunc("sku", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("title", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("inventory-quantity") {
		v, _ := cmd.Flags().GetInt("inventory-quantity")
		body["inventory_quantity"] = v
	}
	if cmd.Flags().Changed("manage-inventory") {
		v, _ := cmd.Flags().GetBool("manage-inventory")
		body["manage_inventory"] = v
	}
	optionsVal, _ := cmd.Flags().GetString("options")
	optionsValJSON, err := utils.JSONValue(optionsVal, "options")
	if err != nil {
		return nil, err
	}
	body["options"] = optionsValJSON
	if cmd.Flags().Changed("prices") {
		pricesVal, _ := cmd.Flags().GetString("prices")
		pricesValJSON, err := utils.JSONValue(pricesVal, "prices")
		if err != nil {
			return nil, err
		}
		body["prices"] = pricesValJSON
	}
	if cmd.Flags().Changed("sku") {
		v, _ := cmd.Flags().GetString("sku")
		body["sku"] = v
	}
	if cmd.Flags().Changed("title") {
		v, _ := cmd.Flags().GetString("title")
		body["title"] = v
	}
	return body, nil
}
