package domains

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
	Short: "List Agency Plan domains",
	Long:  "Returns a paginated list of domains associated with Agency Plan websites accessible to the authenticated client.\n\nUse the website_uuids filter to narrow results to specific websites.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListAgencyPlanDomainsV1WithResponse(context.TODO(), listPlanParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListPlanCmd.Flags().IntP("page", "", 0, "Page number")
	ListPlanCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListPlanCmd.Flags().StringSliceP("website-uuids", "", nil, "Filter by website UIDs")
}

func listPlanParams(cmd *cobra.Command) *client.AgencyHostingListAgencyPlanDomainsV1Params {
	params := &client.AgencyHostingListAgencyPlanDomainsV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	if cmd.Flags().Changed("website-uuids") {
		v, _ := cmd.Flags().GetStringSlice("website-uuids")
		params.WebsiteUuids = &v
	}
	return params
}
