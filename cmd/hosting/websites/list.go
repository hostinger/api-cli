package websites

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list",
	Short:             "List websites",
	Long:              "Retrieve a paginated list of websites (CloudLinux, Builder, and Horizons) accessible to the\nauthenticated client.\n\nThis endpoint returns websites from your hosting accounts as well as\nwebsites from other client hosting accounts that have shared access\nwith you.\n\nEach website includes a `website_type` field describing the type of\nwebsite detected on the underlying platform (`wordpress`, `builder`,\n`horizons`, `nodejs`, or `other`). Some fields, such as\n`vhost_type`, `username`, and `root_directory`, only apply to\nCloudLinux websites and are null for other platforms.\n\nUse `website_types` to list only websites of a given detected type, e.g. only\nWordPress websites (`website_types=wordpress`) or only Node.js websites\n(`website_types=nodejs`). Combine with the other available query parameters to\nfilter by username, order ID, enabled status, or domain name for more targeted\nresults.\n\nA website appears in this list before its server-side setup has finished, and\n`is_enabled` reflects suspension, not readiness. To know when a newly created website\nis ready for file, deploy or database operations, poll the list website setups\nendpoint instead.",
	Example:           "  hostinger hosting websites list",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListWebsitesV1WithResponse(ctx, params)
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
	utils.RegisterArgs(ListCmd)
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCmd.Flags().StringP("username", "", "", "Filter by specific username")
	ListCmd.Flags().IntP("order-id", "", 0, "Order ID")
	ListCmd.Flags().BoolP("is-enabled", "", false, "Filter by enabled status")
	ListCmd.Flags().StringP("domain", "", "", "Filter by domain name (case-insensitive substring match)")
	ListCmd.Flags().StringSliceP("website-types", "", nil, "Filter by detected website type, e.g. wordpress,nodejs. Accepts a comma-separated list. (one of: wordpress, builder, horizons, nodejs, other)")
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("username", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("order-id", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("is-enabled", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("website-types", cobra.FixedCompletions([]cobra.Completion{"wordpress", "builder", "horizons", "nodejs", "other"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listParams(cmd *cobra.Command) *client.HostingListWebsitesV1Params {
	params := &client.HostingListWebsitesV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	if cmd.Flags().Changed("username") {
		v, _ := cmd.Flags().GetString("username")
		params.Username = &v
	}
	if cmd.Flags().Changed("order-id") {
		v, _ := cmd.Flags().GetInt("order-id")
		params.OrderId = &v
	}
	if cmd.Flags().Changed("is-enabled") {
		v, _ := cmd.Flags().GetBool("is-enabled")
		params.IsEnabled = &v
	}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	if cmd.Flags().Changed("website-types") {
		v, _ := cmd.Flags().GetStringSlice("website-types")
		params.WebsiteTypes = &v
	}
	return params
}
