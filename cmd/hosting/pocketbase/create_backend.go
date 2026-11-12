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

var CreateBackendCmd = &cobra.Command{
	Use:     "create-backend [domain]",
	Short:   "Create PocketBase backend",
	Long:    "Creates a PocketBase backend for a website: a database, authentication, file storage and a\nrealtime REST API, served from the website's own domain. A website has at most one\nPocketBase backend. Creation usually takes a few seconds.\n\nThe backend takes over the `/api` and `/_` paths of the website: from then on requests to\nthese paths and everything below them go to PocketBase, while every other path keeps\nserving the website. Do not create a backend for a website that serves its own `/api` or\n`/_` paths, because they stop working; confirm with the user before calling it. Deleting\nthe backend gives both paths back to the website. Headless WordPress websites and parked\ndomains are not supported; for a parked domain, use the website it is parked on. A website\nthat already has a backend is rejected with 422, unless that backend is in state\n`missing`: then it is replaced with a new, empty one. A website set up with\n`type: headless_empty` becomes `headless_pocketbase`; any other website keeps its type.\n\n`superuser_password` is optional: leave it out to get a strong generated password, or send\nyour own (10-128 characters). The response returns it in `superuser_password` once and\nHostinger never stores it, so save it in your app secrets right away; you need it to manage\nthe backend through the PocketBase REST API. If it is lost, set a new one with the\n`Set PocketBase superuser` endpoint. `superuser_email` defaults to the email of the\nwebsite owner.\n\nThe instance sleeps when idle and wakes on the first request, so the first call after a\nquiet period can take a few seconds. `next_steps` in the response lists what to do next.\n\nPocketBase is rolling out gradually: an account without access yet gets 422.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting pocketbase create-backend mydomain.tld",
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
		body, err := createBackendBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingCreatePocketBaseBackendV1WithBodyWithResponse(ctx, username, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(CreateBackendCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	CreateBackendCmd.Flags().StringP("superuser-email", "", "", "Email of the PocketBase superuser. Defaults to the email of the website owner.")
	CreateBackendCmd.Flags().StringP("superuser-password", "", "", "Password for the PocketBase superuser, 10-128 characters. Leave it out to have a strong one generated. Either way it is returned once in the response and never stored.")
	CreateBackendCmd.RegisterFlagCompletionFunc("superuser-email", cobra.NoFileCompletions)
	CreateBackendCmd.RegisterFlagCompletionFunc("superuser-password", cobra.NoFileCompletions)
}

func createBackendBody(cmd *cobra.Command) (map[string]any, error) {
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
