package orders

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list",
	Short:             "List orders",
	Long:              "Retrieve a paginated list of orders accessible to the authenticated client.\n\nOnly Web and Cloud hosting orders are listed. Agency Plan orders are listed by\n`GET /api/agency-hosting/v1/orders`.\n\nThis endpoint returns orders of your hosting accounts as well as orders\nof other client hosting accounts that have shared access with you.\n\nUse the available query parameters to filter results by order statuses\nor specific order IDs for more targeted results.",
	Example:           "  hostinger hosting orders list",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListOrdersV1WithResponse(ctx, params)
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
	utils.RegisterArgs(ListCmd)
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCmd.Flags().StringSliceP("statuses", "", nil, "Filter by order statuses (one of: active, deleting, deleted, suspended)")
	ListCmd.Flags().IntSliceP("order-ids", "", nil, "Filter by specific order IDs")
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("statuses", cobra.FixedCompletions([]cobra.Completion{"active", "deleting", "deleted", "suspended"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("order-ids", cobra.NoFileCompletions)
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listParams(cmd *cobra.Command) *client.HostingListOrdersV1Params {
	params := &client.HostingListOrdersV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	if cmd.Flags().Changed("statuses") {
		v, _ := cmd.Flags().GetStringSlice("statuses")
		params.Statuses = &v
	}
	if cmd.Flags().Changed("order-ids") {
		v, _ := cmd.Flags().GetIntSlice("order-ids")
		params.OrderIds = &v
	}
	return params
}
