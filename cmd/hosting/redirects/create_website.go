package redirects

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateWebsiteCmd = &cobra.Command{
	Use:     "create-website <username> <domain>",
	Short:   "Create website redirect",
	Long:    "Creates a redirect from a URL on the selected website to another URL or IP address.",
	Example: "  hostinger hosting redirects create-website u123456789 mydomain.tld --from https://example.com/old-page --to https://example.com/new-page",
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
		body, err := createWebsiteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingCreateWebsiteRedirectV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateWebsiteCmd, utils.StringArg("username"), utils.StringArg("domain"))
	CreateWebsiteCmd.Flags().StringP("from", "", "", "Source URL on the selected website")
	CreateWebsiteCmd.Flags().StringP("to", "", "", "Destination URL or IP address")
	CreateWebsiteCmd.MarkFlagRequired("from")
	CreateWebsiteCmd.MarkFlagRequired("to")
	CreateWebsiteCmd.RegisterFlagCompletionFunc("from", cobra.NoFileCompletions)
	CreateWebsiteCmd.RegisterFlagCompletionFunc("to", cobra.NoFileCompletions)
}

func createWebsiteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	fromVal, _ := cmd.Flags().GetString("from")
	body["from"] = fromVal
	toVal, _ := cmd.Flags().GetString("to")
	body["to"] = toVal
	return body, nil
}
