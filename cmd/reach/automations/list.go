package automations

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
	Short: "List automations",
	Long:  "Get a paginated list of the automations in a profile.\n\nEvery automation comes with the counts of contacts that entered it, are moving through it,\nfinished it or failed on the way. Those counts describe the contact journey and are not\nemail engagement metrics - for opens, clicks and unsubscribes use the campaign statistics\nendpoint instead.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"active", "paused", "draft"})
		utils.EnumCheck(cmd, "sort-direction", []string{"asc", "desc"})
		r, err := api.Request().ReachListAutomationsV1WithResponse(context.TODO(), args[0], listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("status", "", "", "Filter automations by status.\n\nThere is no `completed` status. An automation that has finished for every contact still\nreports `active`. (one of: active, paused, draft)")
	ListCmd.Flags().StringP("sort-direction", "", "", "Order automations by creation date. Newest first unless set to `asc`. (one of: asc, desc)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listParams(cmd *cobra.Command) *client.ReachListAutomationsV1Params {
	params := &client.ReachListAutomationsV1Params{}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		e := client.ReachListAutomationsV1ParamsStatus(v)
		params.Status = &e
	}
	if cmd.Flags().Changed("sort-direction") {
		v, _ := cmd.Flags().GetString("sort-direction")
		e := client.ReachListAutomationsV1ParamsSortDirection(v)
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
