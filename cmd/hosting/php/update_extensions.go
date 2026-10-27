package php

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

var UpdateExtensionsCmd = &cobra.Command{
	Use:     "update-extensions [domain]",
	Short:   "Update PHP extensions",
	Long:    "Enables or disables PHP extensions (modules) for the website.\n\nUse the Get PHP details endpoint to check the current extension states before changing them.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting php update-extensions mydomain.tld",
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
		body, err := updateExtensionsBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingUpdatePHPExtensionsV1WithBodyWithResponse(ctx, username, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(UpdateExtensionsCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	UpdateExtensionsCmd.Flags().StringSliceP("disable", "", nil, "PHP extensions to disable.")
	UpdateExtensionsCmd.Flags().StringSliceP("enable", "", nil, "PHP extensions to enable.")
	UpdateExtensionsCmd.RegisterFlagCompletionFunc("disable", cobra.NoFileCompletions)
	UpdateExtensionsCmd.RegisterFlagCompletionFunc("enable", cobra.NoFileCompletions)
}

func updateExtensionsBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("disable") {
		v, _ := cmd.Flags().GetStringSlice("disable")
		body["disable"] = v
	}
	if cmd.Flags().Changed("enable") {
		v, _ := cmd.Flags().GetStringSlice("enable")
		body["enable"] = v
	}
	return body, nil
}
