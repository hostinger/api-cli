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

var UpdateVersionCmd = &cobra.Command{
	Use:     "update-version [domain]",
	Short:   "Update PHP version",
	Long:    "Changes the PHP version of the website.\n\nUse the Get PHP details endpoint to see the versions available for the website.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting php update-version mydomain.tld --version 8.1",
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
		body, err := updateVersionBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingUpdatePHPVersionV1WithBodyWithResponse(ctx, username, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(UpdateVersionCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	UpdateVersionCmd.Flags().StringP("version", "", "", "PHP version to switch the website to.")
	UpdateVersionCmd.MarkFlagRequired("version")
	UpdateVersionCmd.RegisterFlagCompletionFunc("version", cobra.NoFileCompletions)
}

func updateVersionBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	versionVal, _ := cmd.Flags().GetString("version")
	body["version"] = versionVal
	return body, nil
}
