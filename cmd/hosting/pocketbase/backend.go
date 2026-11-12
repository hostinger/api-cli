package pocketbase

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var BackendCmd = &cobra.Command{
	Use:     "backend [domain]",
	Short:   "Get PocketBase backend",
	Long:    "Returns the PocketBase backend of a website: its REST API and dashboard URLs, live state\nand version. A PocketBase backend gives the website a database, authentication, file\nstorage and a realtime REST API, served from the website's own domain under `/api` and\n`/_`. Returns 404 when the website has no PocketBase backend.\n\nThe superuser password is only shown once, when it is set, and is never returned here. If it\nis lost, set a new one with the `Set PocketBase superuser` endpoint.\n\nThe instance sleeps when idle and wakes on the first request, so `asleep` is a normal state\nand the first call to a sleeping instance can take a few seconds. `starting`, `running` and\n`stopping` are healthy; `failed` or `unregistered` mean the instance is broken; `missing`\nmeans its files are gone from the server, and creating the PocketBase backend again sets up\na new, empty one. `state` is null when the server could not be reached; retry later.\n\nPocketBase is rolling out gradually: an account without access yet gets 422.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting pocketbase backend mydomain.tld",
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
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGetPocketBaseBackendV1WithResponse(ctx, username, domain)
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
	utils.RegisterScopedArgs(BackendCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
}
