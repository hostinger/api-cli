package redirects

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteWebsiteCmd = &cobra.Command{
	Use:     "delete-website <username> <domain>",
	Short:   "Delete website redirect",
	Long:    "Permanently deletes the redirect identified by its source URL.\n\nPass the `from` value exactly as returned by the list redirects endpoint.",
	Example: "  hostinger hosting redirects delete-website u123456789 mydomain.tld --from https://example.com/old-page",
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
		params := deleteWebsiteParams(cmd)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete website redirect", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeleteWebsiteRedirectV1WithResponse(ctx, args[0], args[1], params)
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
	utils.RegisterArgs(DeleteWebsiteCmd, utils.StringArg("username"), utils.StringArg("domain"))
	DeleteWebsiteCmd.Flags().StringP("from", "", "", "Source URL returned by the list redirects endpoint.")
	DeleteWebsiteCmd.MarkFlagRequired("from")
	DeleteWebsiteCmd.RegisterFlagCompletionFunc("from", cobra.NoFileCompletions)
	DeleteWebsiteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func deleteWebsiteParams(cmd *cobra.Command) *client.HostingDeleteWebsiteRedirectV1Params {
	params := &client.HostingDeleteWebsiteRedirectV1Params{}
	fromVal, _ := cmd.Flags().GetString("from")
	params.From = fromVal
	return params
}
