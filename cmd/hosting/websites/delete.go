package websites

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:     "delete <domain>",
	Short:   "Delete website",
	Long:    "This endpoint permanently removes a website and all of its data. This action\ncannot be undone. Before calling it, make sure the user understands the\nconsequences and explicitly confirms that they want to proceed.\n\nAll website files, databases and related configuration will be removed.\nThe hosting plan itself is kept, so a new website can be created on it afterwards.\n\nSupported websites: main and addon domain websites on web hosting plans, and\nWebsite Builder websites. Parked domains and subdomains cannot be deleted with\nthis endpoint. The domain must be the exact website domain, not a preview\ndomain or an alias.\n\nReturns 404 when the domain does not exist or does not belong to the\nauthenticated client.\n\nWebsite removal is processed asynchronously and can take a few minutes to\ncomplete. The response returns before the removal finishes.",
	Example: "  hostinger hosting websites delete mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "domain",
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
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete website", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeleteWebsiteV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(DeleteCmd, utils.StringArg("domain"))
	DeleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
