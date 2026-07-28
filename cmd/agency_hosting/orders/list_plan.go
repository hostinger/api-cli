package orders

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
	Short: "List Agency Plan orders",
	Long:  "Returns a paginated list of Agency Plan orders accessible to the authenticated client.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListAgencyPlanOrdersV1WithResponse(context.TODO(), listPlanParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListPlanCmd.Flags().IntP("page", "", 0, "Page number")
	ListPlanCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listPlanParams(cmd *cobra.Command) *client.AgencyHostingListAgencyPlanOrdersV1Params {
	params := &client.AgencyHostingListAgencyPlanOrdersV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	return params
}
