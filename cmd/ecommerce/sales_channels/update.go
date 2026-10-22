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

var UpdateCmd = &cobra.Command{
	Use:               "update <store_id> <sales_channel_id>",
	Short:             "Update sales channel",
	Long:              "Update a custom sales channel. The merchant-facing `name` and the public `url`\n(returned as the channel `domain`) can be changed. Pass `null` to clear a value.",
	Example:           "  hostinger ecommerce sales-channels update store_01J8Z5F8W9K8M4A7B3C2D1E0FG scha_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceUpdateSalesChannelV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("store_id"), utils.StringArg("sales_channel_id"))
	UpdateCmd.Flags().StringP("name", "", "", "Merchant-facing custom name shown in the sales channels list. Pass null to clear it.")
	UpdateCmd.Flags().StringP("url", "", "", "Public address where the custom sales channel lives. Pass null to clear it.")
	UpdateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("url", cobra.NoFileCompletions)
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("url") {
		v, _ := cmd.Flags().GetString("url")
		body["url"] = v
	}
	return body, nil
}
