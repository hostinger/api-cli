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

var UpdateCmd = &cobra.Command{
	Use:               "update <store_id> <product_id>",
	Short:             "Update a product",
	Long:              "Update a product's name, description, status, SEO settings or collections. Set status to\npublished to make it buyable, draft to hide it, or archived to retire it. SEO settings (title,\ndescription, slug) left out keep their values. collection_changes adds or removes collections\nwithout touching the others. Variants, prices and inventory are managed through the variant\nendpoints, not here. Returns the updated product summary.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --collection-changes\n    {\n      \"add\": [\n        {\n          \"id\": \"pcol_01J8Z5F8W9K8M4A7B3C2D1E0FG\"\n        }\n      ],\n      \"remove\": [\n        {\n          \"id\": \"pcol_01J8Z5F8W9K8M4A7B3C2D1E0FG\"\n        }\n      ]\n    }\n\n  --seo-settings\n    {\n      \"description\": \"Soft combed cotton, unisex fit.\",\n      \"slug\": \"blue-t-shirt\",\n      \"title\": \"Blue T-Shirt | My Shop\"\n    }",
	Example:           "  hostinger ecommerce products update store_01J8Z5F8W9K8M4A7B3C2D1E0FG prod_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "status", []string{"draft", "published", "archived"}); err != nil {
			return err
		}
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceUpdateAProductV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("store_id"), utils.StringArg("product_id"))
	UpdateCmd.Flags().StringP("collection-changes", "", "", "Adds the product to or removes it from collections without touching its other collections. Send at least one collection in add or remove. (JSON, @file or -; shape in the help above)")
	UpdateCmd.Flags().StringP("description", "", "", "The product description.")
	UpdateCmd.Flags().StringP("name", "", "", "The product name.")
	UpdateCmd.Flags().StringP("seo-settings", "", "", "Search-engine settings to change. Send at least one field; omitted fields keep their current values. (JSON, @file or -; shape in the help above)")
	UpdateCmd.Flags().StringP("status", "", "", "Set \"published\" to make the product buyable, \"draft\" to hide it, or \"archived\" to retire it. (one of: draft, published, archived)")
	UpdateCmd.RegisterFlagCompletionFunc("collection-changes", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	UpdateCmd.RegisterFlagCompletionFunc("description", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("seo-settings", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	UpdateCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"draft", "published", "archived"}, cobra.ShellCompDirectiveNoFileComp))
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("collection-changes") {
		collectionChangesVal, _ := cmd.Flags().GetString("collection-changes")
		collectionChangesValJSON, err := utils.JSONValue(collectionChangesVal, "collection-changes")
		if err != nil {
			return nil, err
		}
		body["collection_changes"] = collectionChangesValJSON
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
		seoSettingsVal, _ := cmd.Flags().GetString("seo-settings")
		seoSettingsValJSON, err := utils.JSONValue(seoSettingsVal, "seo-settings")
		if err != nil {
			return nil, err
		}
		body["seo_settings"] = seoSettingsValJSON
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		body["status"] = v
	}
	return body, nil
}
