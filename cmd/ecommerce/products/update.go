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
	Long:  "Update a product's name, description, status, SEO settings or collections. Set status to\npublished to make it buyable, draft to hide it, or archived to retire it. SEO settings (title,\ndescription, slug) left out keep their values. collection_changes adds or removes collections\nwithout touching the others. Variants, prices and inventory are managed through the variant\nendpoints, not here. Returns the updated product summary.",
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
	UpdateCmd.Flags().StringP("collection-changes", "", "", "Adds the product to or removes it from collections without touching its other collections. Send at least one collection in add or remove. (JSON)")
	UpdateCmd.Flags().StringP("description", "", "", "The product description.")
	UpdateCmd.Flags().StringP("name", "", "", "The product name.")
	UpdateCmd.Flags().StringP("seo-settings", "", "", "Search-engine settings to change. Send at least one field; omitted fields keep their current values. (JSON)")
	UpdateCmd.Flags().StringP("status", "", "", "Set \"published\" to make the product buyable, \"draft\" to hide it, or \"archived\" to retire it. (one of: draft, published, archived)")
}

func updateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("collection-changes") {
		v, _ := cmd.Flags().GetString("collection-changes")
		body["collection_changes"] = utils.JSONValue(v, "collection-changes")
	}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("seo-settings") {
		v, _ := cmd.Flags().GetString("seo-settings")
		body["seo_settings"] = utils.JSONValue(v, "seo-settings")
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		body["status"] = v
	}
	return body
}
