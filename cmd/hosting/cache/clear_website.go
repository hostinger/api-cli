package cache

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ClearWebsiteCmd = &cobra.Command{
	Use:     "clear-website <username> <domain>",
	Short:   "Clear website cache",
	Long:    "Permanently clears all server-side cache for the website at once. Use it when content was\nupdated and needs to be visible immediately, or after making major changes.\n\nAlso purges the Hostinger CDN cache when CDN is enabled on the website. To purge only the CDN\ncache and leave the server cache untouched, use `Purge CDN cache`. For a WordPress\ninstallation living in a subdirectory, pass the directory query parameter to clear its cache.",
	Example: "  hostinger hosting cache clear-website u123456789 mydomain.tld",
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
		params := clearWebsiteParams(cmd)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Clear website cache", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingClearWebsiteCacheV1WithResponse(ctx, args[0], args[1], params)
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
	utils.RegisterArgs(ClearWebsiteCmd, utils.StringArg("username"), utils.StringArg("domain"))
	ClearWebsiteCmd.Flags().StringP("directory", "", "", "Directory of the website installation to clear, relative to the website root. Defaults to the website root.")
	ClearWebsiteCmd.RegisterFlagCompletionFunc("directory", cobra.NoFileCompletions)
	ClearWebsiteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func clearWebsiteParams(cmd *cobra.Command) *client.HostingClearWebsiteCacheV1Params {
	params := &client.HostingClearWebsiteCacheV1Params{}
	if cmd.Flags().Changed("directory") {
		v, _ := cmd.Flags().GetString("directory")
		params.Directory = &v
	}
	return params
}
