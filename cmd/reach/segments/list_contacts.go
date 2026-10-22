package segments

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListContactsCmd = &cobra.Command{
	Use:     "list-contacts <segment-uuid>",
	Short:   "List segment contacts",
	Long:    "Retrieve contacts associated with a specific segment.\n\nThis endpoint allows you to fetch and filter contacts that belong to a particular segment,\nidentified by its UUID.\n\nDeprecated. This endpoint cannot target a profile, so it always falls back to\nthe client's default profile and cannot read segments of any other profile. Use\n`GET /api/reach/v1/profiles/{profileUuid}/segmentation/segments/{segmentUuid}/contacts` instead.",
	Example: "  hostinger reach segments list-contacts 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: utils.CompleteArgs(
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
		params := listContactsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachListSegmentContactsV1WithResponse(ctx, args[0], params)
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
	utils.RegisterArgs(ListContactsCmd, utils.StringArg("segment-uuid"))
	ListContactsCmd.Flags().IntP("page", "", 0, "Page number")
	ListContactsCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListContactsCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListContactsCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListContactsCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListContactsCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listContactsParams(cmd *cobra.Command) *client.ReachListSegmentContactsV1Params {
	params := &client.ReachListSegmentContactsV1Params{}
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
