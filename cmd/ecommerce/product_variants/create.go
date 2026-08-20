package product_variants

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
	Use:   "create <store_id> <product_id>",
	Short: "Create a product variant",
	Long:  "Add a variant to a product along one or more option dimensions (e.g. Size, Color). Options\nmissing from the product are created automatically; provide a value for every option the\nproduct already has. Prices are integers in the smallest currency unit and default to the\nstore currency. Returns the created variant.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceCreateAProductVariantV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().IntP("inventory-quantity", "", 0, "Units in stock. Defaults to 0.")
	CreateCmd.Flags().BoolP("manage-inventory", "", false, "Whether stock is tracked for this variant. Defaults to false.")
	CreateCmd.Flags().StringP("options", "", "", "Option name/value pairs that distinguish this variant, e.g. [{name: Size, value: M}]. Options missing from the product are created; provide a value for every option the product already has. (JSON)")
	CreateCmd.Flags().StringP("prices", "", "", "Prices per currency. Amounts are integers in the smallest currency unit. A free item is amount: 0. (JSON)")
	CreateCmd.Flags().StringP("sku", "", "", "The variant SKU.")
	CreateCmd.Flags().StringP("title", "", "", "The variant title. Defaults to the option values joined with ' / ' (e.g. 'Red / L').")
	CreateCmd.MarkFlagRequired("options")
}

func createBody(cmd *cobra.Command) map[string]any {
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
	body["options"] = utils.JSONValue(optionsVal, "options")
	if cmd.Flags().Changed("prices") {
		v, _ := cmd.Flags().GetString("prices")
		body["prices"] = utils.JSONValue(v, "prices")
	}
	if cmd.Flags().Changed("sku") {
		v, _ := cmd.Flags().GetString("sku")
		body["sku"] = v
	}
	if cmd.Flags().Changed("title") {
		v, _ := cmd.Flags().GetString("title")
		body["title"] = v
	}
	return body
}
