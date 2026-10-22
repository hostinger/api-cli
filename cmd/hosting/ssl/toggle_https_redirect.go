package ssl

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ToggleHttpsRedirectCmd = &cobra.Command{
	Use:     "toggle-https-redirect <username> <domain>",
	Short:   "Toggle HTTPS redirect",
	Long:    "Turns the HTTP to HTTPS redirect of the website on or off, based on `is_enabled`. Does\nnothing when the redirect is already in the requested state. Turning it on requires an\ninstalled certificate (`status` `active` or `expired` on `Get SSL status`) and returns 422\nwhen there is none; turning it off is always accepted.",
	Example: "  hostinger hosting ssl toggle-https-redirect u123456789 mydomain.tld --is-enabled=true",
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
		body, err := toggleHttpsRedirectBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingToggleHTTPSRedirectV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ToggleHttpsRedirectCmd, utils.StringArg("username"), utils.StringArg("domain"))
	ToggleHttpsRedirectCmd.Flags().BoolP("is-enabled", "", false, "Turn the HTTP to HTTPS redirect on (true) or off (false) for the website.")
	ToggleHttpsRedirectCmd.MarkFlagRequired("is-enabled")
	ToggleHttpsRedirectCmd.RegisterFlagCompletionFunc("is-enabled", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func toggleHttpsRedirectBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	isEnabledVal, _ := cmd.Flags().GetBool("is-enabled")
	body["is_enabled"] = isEnabledVal
	return body, nil
}
