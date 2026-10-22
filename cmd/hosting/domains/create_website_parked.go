package domains

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateWebsiteParkedCmd = &cobra.Command{
	Use:     "create-website-parked <username> <domain>",
	Short:   "Create website parked domain",
	Long:    "Create a parked or alias domain for the selected website.\n\nProvide a domain name or IP address to park on the website so it serves the same content\nas the parent domain.",
	Example: "  hostinger hosting domains create-website-parked u123456789 mydomain.tld --parked-domain parked-domain.com",
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
		body, err := createWebsiteParkedBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingCreateWebsiteParkedDomainV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateWebsiteParkedCmd, utils.StringArg("username"), utils.StringArg("domain"))
	CreateWebsiteParkedCmd.Flags().StringP("parked-domain", "", "", "Domain name or IP address to park on the selected website")
	CreateWebsiteParkedCmd.MarkFlagRequired("parked-domain")
	CreateWebsiteParkedCmd.RegisterFlagCompletionFunc("parked-domain", cobra.NoFileCompletions)
}

func createWebsiteParkedBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	parkedDomainVal, _ := cmd.Flags().GetString("parked-domain")
	body["parked_domain"] = parkedDomainVal
	return body, nil
}
