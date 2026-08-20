package products

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

var UpdateCmd = &cobra.Command{
	Use:   "update <store_id> <product_id>",
	Short: "Update a product",
	Long:  "Update a product's name, description or status. Set status to published to make it buyable,\ndraft to hide it, or archived to retire it. Variants, prices and inventory are managed\nthrough the variant endpoints, not here. Returns the updated product summary.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"draft", "published", "archived"})
		payload, err := json.Marshal(updateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().EcommerceUpdateAProductV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateCmd.Flags().StringP("description", "", "", "The product description.")
	UpdateCmd.Flags().StringP("name", "", "", "The product name.")
	UpdateCmd.Flags().StringP("status", "", "", "Set \"published\" to make the product buyable, \"draft\" to hide it, or \"archived\" to retire it. (one of: draft, published, archived)")
}

func updateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		body["status"] = v
	}
	return body
}
