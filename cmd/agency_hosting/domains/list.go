package domains

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
	Short: "List domains",
	Long:  "Returns a paginated list of domains associated with Agency Plan websites accessible to the authenticated client.\n\nUse the website_uuids filter to narrow results to specific websites.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListDomainsV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCmd.Flags().StringSliceP("website-uuids", "", nil, "Filter by website UIDs")
}

func listParams(cmd *cobra.Command) *client.AgencyHostingListDomainsV1Params {
	params := &client.AgencyHostingListDomainsV1Params{}
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
