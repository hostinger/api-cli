package campaigns

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list <profile-uuid>",
	Short:             "List campaigns",
	Long:              "Get a paginated list of the campaigns in a profile.\n\nEach campaign carries its headline engagement rates. Filter by status to find drafts,\nscheduled, sending or sent campaigns, keeping in mind that a fully sent campaign has the\nstatus `publish`. By default only regular campaigns are returned - pass `type` to get the\nemails sent by automations or the double opt-in confirmations instead.",
	Example:           "  hostinger reach campaigns list 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "status", []string{"draft", "scheduled", "sending", "publish", "failed"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "type", []string{"campaign", "automation", "double_opt_in"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "sort-direction", []string{"asc", "desc"}); err != nil {
			return err
		}
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachListCampaignsV1WithResponse(ctx, args[0], params)
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
	utils.RegisterArgs(ListCmd, utils.StringArg("profile-uuid"))
	ListCmd.Flags().StringP("status", "", "", "Filter campaigns by status. A fully sent campaign has the status `publish`. There is no `sent` status, and campaigns can be neither paused nor archived. (one of: draft, scheduled, sending, publish, failed)")
	ListCmd.Flags().StringP("type", "", "campaign", "Filter campaigns by type. Defaults to `campaign`, which leaves out the emails sent by automations and the double opt-in confirmations. (one of: campaign, automation, double_opt_in)")
	ListCmd.Flags().StringP("sort-direction", "", "", "Order campaigns by creation date. Newest first unless set to `asc`. (one of: asc, desc)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"draft", "scheduled", "sending", "publish", "failed"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("type", cobra.FixedCompletions([]cobra.Completion{"campaign", "automation", "double_opt_in"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("sort-direction", cobra.FixedCompletions([]cobra.Completion{"asc", "desc"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
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
