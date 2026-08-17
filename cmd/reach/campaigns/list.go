package campaigns

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <profile-uuid>",
	Short: "List campaigns",
	Long:  "Get a paginated list of the campaigns in a profile.\n\nEach campaign carries its headline engagement rates. Filter by status to find drafts,\nscheduled, sending or sent campaigns, keeping in mind that a fully sent campaign has the\nstatus `publish`. By default only regular campaigns are returned - pass `type` to get the\nemails sent by automations or the double opt-in confirmations instead.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"draft", "scheduled", "sending", "publish", "failed"})
		utils.EnumCheck(cmd, "type", []string{"campaign", "automation", "double_opt_in"})
		utils.EnumCheck(cmd, "sort-direction", []string{"asc", "desc"})
		r, err := api.Request().ReachListCampaignsV1WithResponse(context.TODO(), args[0], listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("status", "", "", "Filter campaigns by status.\n\nA fully sent campaign has the status `publish`. There is no `sent` status, and campaigns can\nbe neither paused nor archived. (one of: draft, scheduled, sending, publish, failed)")
	ListCmd.Flags().StringP("type", "", "campaign", "Filter campaigns by type.\n\nDefaults to `campaign`, which leaves out the emails sent by automations and the double\nopt-in confirmations. (one of: campaign, automation, double_opt_in)")
	ListCmd.Flags().StringP("sort-direction", "", "", "Order campaigns by creation date. Newest first unless set to `asc`. (one of: asc, desc)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listParams(cmd *cobra.Command) *client.ReachListCampaignsV1Params {
	params := &client.ReachListCampaignsV1Params{}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		e := client.ReachListCampaignsV1ParamsStatus(v)
		params.Status = &e
	}
	if cmd.Flags().Changed("type") {
		v, _ := cmd.Flags().GetString("type")
		e := client.ReachListCampaignsV1ParamsType(v)
		params.Type = &e
	}
	if cmd.Flags().Changed("sort-direction") {
		v, _ := cmd.Flags().GetString("sort-direction")
		e := client.ReachListCampaignsV1ParamsSortDirection(v)
		params.SortDirection = &e
	}
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
