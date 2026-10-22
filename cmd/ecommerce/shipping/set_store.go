package shipping

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SetStoreCmd = &cobra.Command{
	Use:               "set-store <store_id>",
	Short:             "Set store shipping",
	Long:              "Set the flat-rate shipping price for a store, creating the shipping zone if it does not exist yet.",
	Example:           "  hostinger ecommerce shipping set-store store_01J8Z5F8W9K8M4A7B3C2D1E0FG --price 500",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := setStoreBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceSetStoreShippingV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(SetStoreCmd, utils.StringArg("store_id"))
	SetStoreCmd.Flags().IntP("price", "", 0, "Flat shipping rate in the smallest currency unit (e.g. cents). Use 0 for free shipping.")
	SetStoreCmd.MarkFlagRequired("price")
	SetStoreCmd.RegisterFlagCompletionFunc("price", cobra.NoFileCompletions)
}

func setStoreBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	priceVal, _ := cmd.Flags().GetInt("price")
	body["price"] = priceVal
	return body, nil
}
