package orders

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List orders",
	Long:  "Retrieve a paginated list of orders accessible to the authenticated client.\n\nThis endpoint returns orders of your hosting accounts as well as orders\nof other client hosting accounts that have shared access with you.\n\nUse the available query parameters to filter results by order statuses\nor specific order IDs for more targeted results.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListOrdersV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCmd.Flags().StringSliceP("statuses", "", nil, "Filter by order statuses (one of: active, deleting, deleted, suspended)")
	ListCmd.Flags().IntSliceP("order-ids", "", nil, "Filter by specific order IDs")
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
