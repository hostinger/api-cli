package discounts

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
	Short:             "Create a discount",
	Long:              "Create a discount for a store. Fixed discounts take an amount in the smallest currency\nunit (e.g. $10 is 1000); percentage discounts take a whole-number value between 1 and 100.\nFree-shipping discounts ignore value. Returns the created discount.",
	Example:           "  hostinger ecommerce discounts create store_01J8Z5F8W9K8M4A7B3C2D1E0FG --code BLACKFRIDAY --type percentage --value 20",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "allocation", []string{"total", "item"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "type", []string{"percentage", "fixed", "free_shipping"}); err != nil {
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
			r, err := api.Request().EcommerceCreateADiscountV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	CreateCmd.Flags().StringP("allocation", "", "", "Whether the discount applies to the cart total or to each eligible item. (one of: total, item)")
	CreateCmd.Flags().StringP("code", "", "", "The discount code customers enter at checkout.")
	CreateCmd.Flags().StringP("ends-at", "", "", "When the discount expires. A bare date runs to the end of that day in time_zone. Never expires when omitted.")
	CreateCmd.Flags().IntP("min-cart-value", "", 0, "Minimum cart value in the smallest currency unit required for the discount to apply.")
	CreateCmd.Flags().StringP("name", "", "", "A human-friendly discount name.")
	CreateCmd.Flags().StringP("starts-at", "", "", "When the discount becomes active. A bare date (2026-11-27) anchors to time_zone. Defaults to now when omitted.")
	CreateCmd.Flags().StringP("time-zone", "", "", "IANA time zone used to interpret starts_at and ends_at.")
	CreateCmd.Flags().StringP("type", "", "", "The discount type. (one of: percentage, fixed, free_shipping)")
	CreateCmd.Flags().IntP("usage-limit", "", 0, "Maximum number of times the discount can be redeemed.")
	CreateCmd.Flags().IntP("value", "", 0, "For percentage discounts a whole number 1-100; for fixed discounts an amount in the smallest currency unit (e.g. $10 is 1000). Ignored for free_shipping.")
	CreateCmd.MarkFlagRequired("code")
	CreateCmd.MarkFlagRequired("type")
	CreateCmd.MarkFlagRequired("value")
	CreateCmd.RegisterFlagCompletionFunc("allocation", cobra.FixedCompletions([]cobra.Completion{"total", "item"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("code", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("ends-at", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("min-cart-value", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("starts-at", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("time-zone", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("type", cobra.FixedCompletions([]cobra.Completion{"percentage", "fixed", "free_shipping"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("usage-limit", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("value", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("allocation") {
		v, _ := cmd.Flags().GetString("allocation")
		body["allocation"] = v
	}
	codeVal, _ := cmd.Flags().GetString("code")
	body["code"] = codeVal
	if cmd.Flags().Changed("ends-at") {
		v, _ := cmd.Flags().GetString("ends-at")
		body["ends_at"] = v
	}
	if cmd.Flags().Changed("min-cart-value") {
		v, _ := cmd.Flags().GetInt("min-cart-value")
		body["min_cart_value"] = v
	}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("starts-at") {
		v, _ := cmd.Flags().GetString("starts-at")
		body["starts_at"] = v
	}
	if cmd.Flags().Changed("time-zone") {
		v, _ := cmd.Flags().GetString("time-zone")
		body["time_zone"] = v
	}
	typeVal, _ := cmd.Flags().GetString("type")
	body["type"] = typeVal
	if cmd.Flags().Changed("usage-limit") {
		v, _ := cmd.Flags().GetInt("usage-limit")
		body["usage_limit"] = v
	}
	valueVal, _ := cmd.Flags().GetInt("value")
	body["value"] = valueVal
	return body, nil
}
