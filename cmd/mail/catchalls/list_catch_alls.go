package catchalls

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCatchAllsCmd = &cobra.Command{
	Use:     "list-catch-alls [order-id]",
	Short:   "List catch-alls",
	Long:    "Retrieve a paginated list of catch-alls across all mailboxes of a\nmail order.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).",
	Example: "  hostinger mail catchalls list-catch-alls OR1a2b3c4d5e6f7g",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "id",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().MailListOrdersV1WithResponse(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		orderID, args, err := site.ResolveOrder(cmd, args, 0, "mail")
		if err != nil {
			return api.CommandError(err)
		}
		params := listCatchAllsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailListCatchAllsV1WithResponse(ctx, orderID, params)
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
	utils.RegisterScopedArgs(ListCatchAllsCmd, 1, utils.StringArg("order-id"))
	ListCatchAllsCmd.Flags().IntP("page", "", 0, "Page number")
	ListCatchAllsCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCatchAllsCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCatchAllsCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListCatchAllsCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCatchAllsCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listCatchAllsParams(cmd *cobra.Command) *client.MailListCatchAllsV1Params {
	params := &client.MailListCatchAllsV1Params{}
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
