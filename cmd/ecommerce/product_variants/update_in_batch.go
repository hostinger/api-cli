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

var UpdateInBatchCmd = &cobra.Command{
	Use:               "update-in-batch <store_id> <product_id>",
	Short:             "Update product variants in batch",
	Long:              "Update up to 100 existing variants in place by id — title, SKU, inventory, stock tracking\nand prices. Variants omitted from the request are left untouched. Prices replace the variant's\nexisting prices in full. Returns the updated variants.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --variants\n    [\n      {\n        \"inventory_quantity\": 42,\n        \"manage_inventory\": true,\n        \"prices\": [\n          {\n            \"amount\": 1999,\n            \"currency\": \"usd\",\n            \"sale_amount\": 1499\n          }\n        ],\n        \"sku\": \"TSHIRT-RED-L\",\n        \"title\": \"Red / L\",\n        \"variant_id\": \"variant_01J8Z5F8W9K8M4A7B3C2D1E0FG\"\n      }\n    ]",
	Example:           "  hostinger ecommerce product-variants update-in-batch store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG --variants @variants.json",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := updateInBatchBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceUpdateProductVariantsInBatchV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateInBatchCmd, utils.StringArg("store_id"), utils.StringArg("product_id"))
	UpdateInBatchCmd.Flags().StringP("variants", "", "", "Variants to update in place by id, up to 100. Variants omitted from the list are left untouched. (JSON, @file or -; shape in the help above)")
	UpdateInBatchCmd.MarkFlagRequired("variants")
	UpdateInBatchCmd.RegisterFlagCompletionFunc("variants", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func updateInBatchBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	variantsVal, _ := cmd.Flags().GetString("variants")
	variantsValJSON, err := utils.JSONValue(variantsVal, "variants")
	if err != nil {
		return nil, err
	}
	body["variants"] = variantsValJSON
	return body, nil
}
