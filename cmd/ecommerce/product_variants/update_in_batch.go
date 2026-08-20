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

var UpdateInBatchCmd = &cobra.Command{
	Use:   "update-in-batch <store_id> <product_id>",
	Short: "Update product variants in batch",
	Long:  "Update up to 100 existing variants in place by id — title, inventory, stock tracking and\nprices. Variants omitted from the request are left untouched. Prices replace the variant's\nexisting prices in full. Returns the updated variants.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateInBatchBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceUpdateProductVariantsInBatchV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateInBatchCmd.Flags().StringP("variants", "", "", "Variants to update in place by id, up to 100. Variants omitted from the list are left untouched. (JSON)")
	UpdateInBatchCmd.MarkFlagRequired("variants")
}

func updateInBatchBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	variantsVal, _ := cmd.Flags().GetString("variants")
	body["variants"] = utils.JSONValue(variantsVal, "variants")
	return body
}
