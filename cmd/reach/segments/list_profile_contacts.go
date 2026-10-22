package segments

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListProfileContactsCmd = &cobra.Command{
	Use:     "list-profile-contacts <profile-uuid> <segment-uuid>",
	Short:   "List profile segment contacts",
	Long:    "Retrieve contacts associated with a specific segment for a given profile.\n\nThis endpoint allows you to fetch and filter contacts that belong to a particular segment,\nidentified by its UUID, scoped to a specific profile.",
	Example: "  hostinger reach segments list-profile-contacts 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "uuid",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().ReachListSegmentsV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listProfileContactsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachListProfileSegmentContactsV1WithResponse(ctx, args[0], args[1], params)
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
	utils.RegisterArgs(ListProfileContactsCmd, utils.StringArg("profile-uuid"), utils.StringArg("segment-uuid"))
	ListProfileContactsCmd.Flags().IntP("page", "", 0, "Page number")
	ListProfileContactsCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListProfileContactsCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListProfileContactsCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListProfileContactsCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListProfileContactsCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listProfileContactsParams(cmd *cobra.Command) *client.ReachListProfileSegmentContactsV1Params {
	params := &client.ReachListProfileSegmentContactsV1Params{}
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
