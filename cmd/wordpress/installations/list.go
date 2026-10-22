package installations

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
	Short:             "List WordPress installations",
	Long:              "List WordPress installations on shared and cloud web hosting accounts (the\n`/api/hosting/v1` family, identified by `username`) accessible to the\nauthenticated client.\n\nAgency Plan websites are not included. To work with WordPress on an Agency Plan\nwebsite, list them with GET /api/agency-hosting/v1/websites?website_types=wordpress\nand read their details with\nGET /api/agency-hosting/v1/websites/{website_uid}/wordpress/settings. An account\nthat only has Agency Plan websites receives an empty list from this endpoint.\n\nUse this endpoint to discover existing WordPress installations and to poll\nfor installation status after calling the install endpoint. When a newly\nrequested installation appears in this list, WordPress is ready. Filter by\nusername and domain to narrow results to a specific website.\n\nEach installation includes a `valid` flag and, when invalid, a\n`validationError` describing why.",
	Example:           "  hostinger wordpress installations list",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "ownership", []string{"owned", "managed", "all"}); err != nil {
			return err
		}
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListWordPressInstallationsV1WithResponse(ctx, params)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
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
	ListCmd.Flags().StringP("username", "", "", "Filter by specific username")
	ListCmd.Flags().StringP("domain", "", "", "Filter by domain name (case-insensitive substring match)")
	ListCmd.Flags().StringP("ownership", "", "", "Filter by ownership type. Defaults to \"owned\". Use \"all\" to include both owned and managed installations. (one of: owned, managed, all)")
	ListCmd.RegisterFlagCompletionFunc("username", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("ownership", cobra.FixedCompletions([]cobra.Completion{"owned", "managed", "all"}, cobra.ShellCompDirectiveNoFileComp))
}

func listParams(cmd *cobra.Command) *client.HostingListWordPressInstallationsV1Params {
	params := &client.HostingListWordPressInstallationsV1Params{}
	if cmd.Flags().Changed("username") {
		v, _ := cmd.Flags().GetString("username")
		params.Username = &v
	}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	if cmd.Flags().Changed("ownership") {
		v, _ := cmd.Flags().GetString("ownership")
		e := client.HostingListWordPressInstallationsV1ParamsOwnership(v)
		params.Ownership = &e
	}
	return params
}
