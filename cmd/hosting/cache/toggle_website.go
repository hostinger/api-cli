package cache

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ToggleWebsiteCmd = &cobra.Command{
	Use:     "toggle-website <username> <domain>",
	Short:   "Toggle website cache",
	Long:    "Turns server-side caching for the website on or off, based on the enabled flag. Enable it for\nfaster page loads, reduced server load, and improved user experience; recommended for production\nwebsites. Disabling may impact performance; to temporarily bypass caching while developing or\ndebugging, prefer toggling cacheless mode instead.\n\nDoes nothing if caching is already in the requested state.",
	Example: "  hostinger hosting cache toggle-website u123456789 mydomain.tld --enabled=true",
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
		body, err := toggleWebsiteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingToggleWebsiteCacheV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ToggleWebsiteCmd, utils.StringArg("username"), utils.StringArg("domain"))
	ToggleWebsiteCmd.Flags().BoolP("enabled", "", false, "Turn server-side caching on (true) or off (false) for the website.")
	ToggleWebsiteCmd.MarkFlagRequired("enabled")
	ToggleWebsiteCmd.RegisterFlagCompletionFunc("enabled", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func toggleWebsiteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	enabledVal, _ := cmd.Flags().GetBool("enabled")
	body["enabled"] = enabledVal
	return body, nil
}
