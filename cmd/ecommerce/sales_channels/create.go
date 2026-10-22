package sales_channels

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
	Use:               "create <store_id>",
	Short:             "Create a sales channel",
	Long:              "Create a sales channel for a store. A \"custom\" channel is headless: build your own frontend and keep\nyour catalog, orders, shipping and payments in sync through the Ecommerce API. A \"quick-link\" channel\nis a hosted one-page store whose handle is auto-generated.",
	Example:           "  hostinger ecommerce sales-channels create store_01J8Z5F8W9K8M4A7B3C2D1E0FG --type custom",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "type", []string{"custom", "quick-link"}); err != nil {
			return err
		}
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceCreateASalesChannelV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd, utils.StringArg("store_id"))
	CreateCmd.Flags().StringP("name", "", "", "Merchant-facing custom name. Required for custom channels; not supported for quick-link.")
	CreateCmd.Flags().StringP("type", "", "", "Sales channel type. \"custom\" is a headless channel: it requires a name and takes an optional public url. \"quick-link\" is a one-page store whose handle is auto-generated; it supports neither name nor url. (one of: custom, quick-link)")
	CreateCmd.Flags().StringP("url", "", "", "Optional public url for the channel. Custom channels only; not supported for quick-link.")
	CreateCmd.MarkFlagRequired("type")
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("type", cobra.FixedCompletions([]cobra.Completion{"custom", "quick-link"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("url", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	typeVal, _ := cmd.Flags().GetString("type")
	body["type"] = typeVal
	if cmd.Flags().Changed("url") {
		v, _ := cmd.Flags().GetString("url")
		body["url"] = v
	}
	return body, nil
}
