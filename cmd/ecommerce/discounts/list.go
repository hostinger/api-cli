package discounts

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list <store_id>",
	Short:             "List discounts",
	Long:              "List a store's discounts. Filter by free text over code and name, or by disabled state.\nAmounts for fixed discounts are integers in the smallest currency unit; percentage\ndiscounts carry a whole-number value between 1 and 100.",
	Example:           "  hostinger ecommerce discounts list store_01J8Z5F8W9K8M4A7B3C2D1E0FG",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "is-disabled", []string{"true", "false"}); err != nil {
			return err
		}
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().EcommerceListDiscountsV1WithResponse(ctx, args[0], params)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		if all, _ := cmd.Flags().GetBool("all"); all {
			call = utils.AllPages(call, func(page int) { params.Page = &page })
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(ListCmd, utils.StringArg("store_id"))
	ListCmd.Flags().StringP("q", "", "", "Free-text search over discount code and name.")
	ListCmd.Flags().StringP("is-disabled", "", "", "Filter by disabled state. (one of: true, false)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.RegisterFlagCompletionFunc("q", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("is-disabled", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listParams(cmd *cobra.Command) *client.EcommerceListDiscountsV1Params {
	params := &client.EcommerceListDiscountsV1Params{}
	if cmd.Flags().Changed("q") {
		v, _ := cmd.Flags().GetString("q")
		params.Q = &v
	}
	if cmd.Flags().Changed("is-disabled") {
		v, _ := cmd.Flags().GetString("is-disabled")
		e := client.EcommerceListDiscountsV1ParamsIsDisabled(v)
		params.IsDisabled = &e
	}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	return params
}
