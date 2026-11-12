package pocketbase

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteBackendCmd = &cobra.Command{
	Use:     "delete-backend [domain]",
	Short:   "Delete PocketBase backend",
	Long:    "Permanently deletes the PocketBase backend of a website with all its data (collections,\nrecords, users and uploaded files) and gives the `/api` and `/_` paths back to the website;\na `headless_pocketbase` website becomes `headless_empty`. This cannot be undone and every\napp using this backend stops working, so confirm with the user before calling it and\nexport the data you need from the dashboard (`admin_url`) first.\n\nReturns success also when the website has no PocketBase backend; then it only clears a\n`headless_pocketbase` mark. To regain access after losing the superuser password, use the\n`Set PocketBase superuser` endpoint instead of deleting the backend.\n\nPocketBase is rolling out gradually: an account without access yet gets 422; an unknown\nwebsite gets 404.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting pocketbase delete-backend mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "domain",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// A destructive command acts without a question only on a target that was typed.
		typed := len(args) > 0
		username, domain, args, err := site.Resolve(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.ConfirmScoped(cmd, "Delete PocketBase backend", append([]string{domain}, args...), typed); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeletePocketBaseBackendV1WithResponse(ctx, username, domain)
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
	utils.RegisterScopedArgs(DeleteBackendCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	DeleteBackendCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
