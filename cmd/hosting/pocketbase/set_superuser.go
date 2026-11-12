package pocketbase

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SetSuperuserCmd = &cobra.Command{
	Use:     "set-superuser [domain]",
	Short:   "Set PocketBase superuser",
	Long:    "Sets the password of a superuser of the website's PocketBase backend, and creates that\nsuperuser when there is none with the given email. The old password of that superuser stops\nworking at once, so apps and scripts still using it lose superuser access; confirm with the\nuser before calling it.\n\nUse it when the superuser password is lost (Hostinger never stores it), when it leaked and\nmust be rotated, or to add another superuser with a different email. Returns 404 when the\nwebsite has no PocketBase backend; create one first. App user accounts are managed through\nthe PocketBase REST API, not here.\n\n`superuser_password` is optional: leave it out to get a strong generated password, or send\nyour own (10-128 characters). The response returns it in `superuser_password` once and\nHostinger never stores it, so save it in your app secrets right away, then authenticate\nagain with `pb.collection('_superusers').authWithPassword`. `superuser_email` defaults to\nthe email of the website owner.\n\nPocketBase is rolling out gradually: an account without access yet gets 422.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting pocketbase set-superuser mydomain.tld",
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
		username, domain, args, err := site.Resolve(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		body, err := setSuperuserBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingSetPocketBaseSuperuserV1WithBodyWithResponse(ctx, username, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(SetSuperuserCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	SetSuperuserCmd.Flags().StringP("superuser-email", "", "", "Email of the superuser to create or update. Defaults to the email of the website owner.")
	SetSuperuserCmd.Flags().StringP("superuser-password", "", "", "New password of the PocketBase superuser, 10-128 characters. Leave it out to have a strong one generated. Either way it is returned once in the response and never stored.")
	SetSuperuserCmd.RegisterFlagCompletionFunc("superuser-email", cobra.NoFileCompletions)
	SetSuperuserCmd.RegisterFlagCompletionFunc("superuser-password", cobra.NoFileCompletions)
}

func setSuperuserBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("superuser-email") {
		v, _ := cmd.Flags().GetString("superuser-email")
		body["superuser_email"] = v
	}
	if cmd.Flags().Changed("superuser-password") {
		v, _ := cmd.Flags().GetString("superuser-password")
		body["superuser_password"] = v
	}
	return body, nil
}
