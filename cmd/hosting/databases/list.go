package databases

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:     "list <username>",
	Short:   "List account databases",
	Long:    "Returns a paginated list of databases for the specified account.\n\nUse the domain and is_assigned filters to find databases assigned to a specific domain.",
	Example: "  hostinger hosting databases list u123456789",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListAccountDatabasesV1WithResponse(ctx, args[0], params)
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
	utils.RegisterArgs(ListCmd, utils.StringArg("username"))
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCmd.Flags().StringP("domain", "", "", "Filter by domain name (case-insensitive substring match)")
	ListCmd.Flags().BoolP("is-assigned", "", false, "When used with domain, return only databases assigned to that domain.")
	ListCmd.Flags().StringP("search", "", "", "Search databases by name, user, or creation date.")
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("is-assigned", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("search", cobra.NoFileCompletions)
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listParams(cmd *cobra.Command) *client.HostingListAccountDatabasesV1Params {
	params := &client.HostingListAccountDatabasesV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	if cmd.Flags().Changed("is-assigned") {
		v, _ := cmd.Flags().GetBool("is-assigned")
		params.IsAssigned = &v
	}
	if cmd.Flags().Changed("search") {
		v, _ := cmd.Flags().GetString("search")
		params.Search = &v
	}
	return params
}
