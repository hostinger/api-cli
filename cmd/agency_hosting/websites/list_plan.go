package websites

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListPlanCmd = &cobra.Command{
	Use:   "list-plan",
	Short: "List Agency Plan websites",
	Long:  "Retrieve a paginated list of Agency Plan websites (H5G, Builder, and Horizons) accessible to\nthe authenticated client.\n\nThis endpoint returns websites from your hosting accounts as well as\nwebsites from other client hosting accounts that have shared access\nwith you.\n\nThe response shape differs per platform — see the `platform` field on each item.\n\nUse `website_types` to list only websites of a given detected type, e.g. only\nWordPress websites (`website_types=wordpress`) or only Node.js websites\n(`website_types=nodejs`). Combine with `order_ids`, `states`, or `domain` for more\ntargeted results.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListAgencyPlanWebsitesV1WithResponse(context.TODO(), listPlanParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListPlanCmd.Flags().IntP("page", "", 0, "Page number")
	ListPlanCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListPlanCmd.Flags().IntSliceP("order-ids", "", nil, "Filter by order IDs. Accepts a comma-separated list.")
	ListPlanCmd.Flags().StringSliceP("states", "", nil, "Filter by website state. Accepts a comma-separated list. (one of: active, locked, suspended, deleting, deleted)")
	ListPlanCmd.Flags().StringSliceP("website-types", "", nil, "Filter by detected website type, e.g. wordpress,nodejs. Accepts a comma-separated list. (one of: wordpress, builder, horizons, nodejs, other)")
	ListPlanCmd.Flags().StringP("domain", "", "", "Filter by domain name (case-insensitive substring match)")
}

func listPlanParams(cmd *cobra.Command) *client.AgencyHostingListAgencyPlanWebsitesV1Params {
	params := &client.AgencyHostingListAgencyPlanWebsitesV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	if cmd.Flags().Changed("order-ids") {
		v, _ := cmd.Flags().GetIntSlice("order-ids")
		params.OrderIds = &v
	}
	if cmd.Flags().Changed("states") {
		v, _ := cmd.Flags().GetStringSlice("states")
		params.States = &v
	}
	if cmd.Flags().Changed("website-types") {
		v, _ := cmd.Flags().GetStringSlice("website-types")
		params.WebsiteTypes = &v
	}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	return params
}
