package git

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListInstallationsCmd = &cobra.Command{
	Use:               "list-installations",
	Short:             "List Git installations",
	Long:              "Lists the Git provider accounts the customer has connected. Only installations with status\n`active` are returned unless the `status` filter says otherwise.\n\nAn empty list means the customer has no active installation. Check `status=suspended` and\n`status=pending` as well. If there is none at all, a Git provider (GitHub or GitLab) has to be\nconnected once in hPanel (Websites, Manage, Advanced, Git; or Add Website, Node.js Web App,\nImport Git Repository); this endpoint then lists the new installation.\n\nUse `uuid` as the path parameter of `List Git installation repositories`, and as\n`installation_uuid` in `Start Node.js build` with `source_type` `git` and in\n`Update Git auto-deployment settings`.",
	Example:           "  hostinger hosting git list-installations",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "provider", []string{"github", "gitlab", "bitbucket"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "status", []string{"pending", "active", "suspended"}); err != nil {
			return err
		}
		params := listInstallationsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListGitInstallationsV1WithResponse(ctx, params)
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
	utils.RegisterArgs(ListInstallationsCmd)
	ListInstallationsCmd.Flags().StringP("provider", "", "", "Filter by Git provider (one of: github, gitlab, bitbucket)")
	ListInstallationsCmd.Flags().StringP("status", "", "active", "Filter by installation status (one of: pending, active, suspended)")
	ListInstallationsCmd.RegisterFlagCompletionFunc("provider", cobra.FixedCompletions([]cobra.Completion{"github", "gitlab", "bitbucket"}, cobra.ShellCompDirectiveNoFileComp))
	ListInstallationsCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"pending", "active", "suspended"}, cobra.ShellCompDirectiveNoFileComp))
}

func listInstallationsParams(cmd *cobra.Command) *client.HostingListGitInstallationsV1Params {
	params := &client.HostingListGitInstallationsV1Params{}
	if cmd.Flags().Changed("provider") {
		v, _ := cmd.Flags().GetString("provider")
		e := client.HostingListGitInstallationsV1ParamsProvider(v)
		params.Provider = &e
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		e := client.HostingListGitInstallationsV1ParamsStatus(v)
		params.Status = &e
	}
	return params
}
