package payments

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var EnableManualMethodCmd = &cobra.Command{
	Use:               "enable-manual-method <store_id>",
	Short:             "Enable manual payment method",
	Long:              "Enable a manual payment method so the store can accept orders without an online payment provider.",
	Example:           "  hostinger ecommerce payments enable-manual-method store_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := enableManualMethodBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceEnableManualPaymentMethodV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(EnableManualMethodCmd, utils.StringArg("store_id"))
	EnableManualMethodCmd.Flags().StringP("title", "", "", "Optional display name shown to customers at checkout.")
	EnableManualMethodCmd.RegisterFlagCompletionFunc("title", cobra.NoFileCompletions)
}

func enableManualMethodBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("title") {
		v, _ := cmd.Flags().GetString("title")
		body["title"] = v
	}
	return body, nil
}
