package git

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListInstallationRepositoriesCmd = &cobra.Command{
	Use:               "list-installation-repositories <uuid>",
	Short:             "List Git installation repositories",
	Long:              "Lists the repositories the Git installation can access, read live from the provider. Works\nfor github and gitlab installations. Use an active installation: a suspended or pending one\nis still queried and the call fails with whatever the provider answers. The list is cut at\nthe first 500 repositories in the order the provider returns them; when the account has\nmore, name the repository directly instead of searching this list.\n\n`owner`, `name` and a branch (`default_branch` or another one) go into `source_options` of\n`Start Node.js build` or into `Update Git auto-deployment settings`. Returns 404 when the\ninstallation does not belong to the customer. Limited to 10 calls per minute per API client\n(429 above that).",
	Example:           "  hostinger hosting git list-installation-repositories 018f5e2a-1234-7890-abcd-1234567890ab",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListGitInstallationRepositoriesV1WithResponse(ctx, uuid.MustParse(args[0]))
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
	utils.RegisterArgs(ListInstallationRepositoriesCmd, utils.UUIDArg("uuid"))
}
