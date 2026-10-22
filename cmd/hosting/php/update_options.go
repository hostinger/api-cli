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

var UpdateOptionsCmd = &cobra.Command{
	Use:     "update-options <username> <domain>",
	Short:   "Update PHP options",
	Long:    "Updates PHP options for the website (e.g. `memory_limit`, `max_execution_time`, `upload_max_filesize`).\nOnly provide the options you want to change, inside the `options` object.\n\nValues above the account plan limit are silently capped to that limit, so the request can succeed\nwith a smaller applied value. Call the Get PHP details endpoint afterwards to read the applied value.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --options\n    {}",
	Example: "  hostinger hosting php update-options u123456789 mydomain.tld --options '{}'",
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
		body, err := updateOptionsBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingUpdatePHPOptionsV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateOptionsCmd, utils.StringArg("username"), utils.StringArg("domain"))
	UpdateOptionsCmd.Flags().StringP("options", "", "", "Map of PHP options to update, keyed by option name. Only include options you want to change. (JSON, @file or -; shape in the help above)")
	UpdateOptionsCmd.MarkFlagRequired("options")
	UpdateOptionsCmd.RegisterFlagCompletionFunc("options", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func updateOptionsBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	optionsVal, _ := cmd.Flags().GetString("options")
	optionsValJSON, err := utils.JSONValue(optionsVal, "options")
	if err != nil {
		return nil, err
	}
	body["options"] = optionsValJSON
	return body, nil
}
