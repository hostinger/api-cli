package contacts

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListProfileCmd = &cobra.Command{
	Use:               "list-profile <profile-uuid>",
	Short:             "List profile contacts",
	Long:              "Get a paginated list of contacts belonging to a profile.\n\nContacts can be filtered by subscription status, by tag, and by an email search term.\nThe `meta.total` field of the response is the number of contacts matching the filters,\nso calling this endpoint without filters gives the profile's total contact count.",
	Example:           "  hostinger reach contacts list-profile 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "subscription-status", []string{"subscribed", "unsubscribed", "confirmed", "pending"}); err != nil {
			return err
		}
		if err := utils.FormatCheck(cmd, "tag-uuid", "uuid"); err != nil {
			return err
		}
		params := listProfileParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachListProfileContactsV1WithResponse(ctx, args[0], params)
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
	utils.RegisterArgs(ListProfileCmd, utils.StringArg("profile-uuid"))
	ListProfileCmd.Flags().StringP("subscription-status", "", "", "Filter contacts by subscription status (one of: subscribed, unsubscribed, confirmed, pending)")
	ListProfileCmd.Flags().StringP("tag-uuid", "", "", "Filter contacts by tag UUID")
	ListProfileCmd.Flags().StringP("search", "", "", "Search contacts by email")
	ListProfileCmd.Flags().IntP("page", "", 0, "Page number")
	ListProfileCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListProfileCmd.RegisterFlagCompletionFunc("subscription-status", cobra.FixedCompletions([]cobra.Completion{"subscribed", "unsubscribed", "confirmed", "pending"}, cobra.ShellCompDirectiveNoFileComp))
	ListProfileCmd.RegisterFlagCompletionFunc("tag-uuid", cobra.NoFileCompletions)
	ListProfileCmd.RegisterFlagCompletionFunc("search", cobra.NoFileCompletions)
	ListProfileCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListProfileCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListProfileCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListProfileCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listProfileParams(cmd *cobra.Command) *client.ReachListProfileContactsV1Params {
	params := &client.ReachListProfileContactsV1Params{}
	if cmd.Flags().Changed("subscription-status") {
		v, _ := cmd.Flags().GetString("subscription-status")
		e := client.ReachListProfileContactsV1ParamsSubscriptionStatus(v)
		params.SubscriptionStatus = &e
	}
	if cmd.Flags().Changed("tag-uuid") {
		v, _ := cmd.Flags().GetString("tag-uuid")
		t := utils.StringToUUID(v)
		params.TagUuid = &t
	}
	if cmd.Flags().Changed("search") {
		v, _ := cmd.Flags().GetString("search")
		params.Search = &v
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
