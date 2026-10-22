package php

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateVersionCmd = &cobra.Command{
	Use:     "update-version <username> <domain>",
	Short:   "Update PHP version",
	Long:    "Changes the PHP version of the website.\n\nUse the Get PHP details endpoint to see the versions available for the website.",
	Example: "  hostinger hosting php update-version u123456789 mydomain.tld --version 8.1",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
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
		body, err := updateVersionBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingUpdatePHPVersionV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateVersionCmd, utils.StringArg("username"), utils.StringArg("domain"))
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
