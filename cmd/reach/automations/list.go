package automations

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
	Short:             "List automations",
	Long:              "Get a paginated list of the automations in a profile.\n\nEvery automation comes with the counts of contacts that entered it, are moving through it,\nfinished it or failed on the way. Those counts describe the contact journey and are not\nemail engagement metrics - for opens, clicks and unsubscribes use the campaign statistics\nendpoint instead.",
	Example:           "  hostinger reach automations list 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "status", []string{"active", "paused", "draft"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "sort-direction", []string{"asc", "desc"}); err != nil {
			return err
		}
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachListAutomationsV1WithResponse(ctx, args[0], params)
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
	ListCmd.Flags().StringP("status", "", "", "Filter automations by status. There is no `completed` status. An automation that has finished for every contact still reports `active`. (one of: active, paused, draft)")
	ListCmd.Flags().StringP("sort-direction", "", "", "Order automations by creation date. Newest first unless set to `asc`. (one of: asc, desc)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"active", "paused", "draft"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("sort-direction", cobra.FixedCompletions([]cobra.Completion{"asc", "desc"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
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
