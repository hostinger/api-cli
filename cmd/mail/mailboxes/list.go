package mailboxes

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:     "list <order-id>",
	Short:   "List mailboxes",
	Long:    "Retrieve a paginated list of mailboxes belonging to a mail order.\n\nUse this endpoint to monitor mailboxes of your mail service, including\ntheir status, enabled protocols, attached resource counts, and\nperiodically synced usage numbers (usage may lag behind live values).",
	Example: "  hostinger mail mailboxes list OR1a2b3c4d5e6f7g",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().MailListOrdersV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "sort", []string{"address", "-address"}); err != nil {
			return err
		}
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailListMailboxesV1WithResponse(ctx, args[0], params)
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
	utils.RegisterArgs(ListCmd, utils.StringArg("order-id"))
	ListCmd.Flags().StringP("search", "", "", "Filter mailboxes whose email address contains the given string")
	ListCmd.Flags().StringP("sort", "", "address", "Sort mailboxes by field. Prefix with `-` for descending order. (one of: address, -address)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCmd.RegisterFlagCompletionFunc("search", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("sort", cobra.FixedCompletions([]cobra.Completion{"address", "-address"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listParams(cmd *cobra.Command) *client.MailListMailboxesV1Params {
	params := &client.MailListMailboxesV1Params{}
	if cmd.Flags().Changed("search") {
		v, _ := cmd.Flags().GetString("search")
		params.Search = &v
	}
	if cmd.Flags().Changed("sort") {
		v, _ := cmd.Flags().GetString("sort")
		e := client.MailListMailboxesV1ParamsSort(v)
		params.Sort = &e
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
