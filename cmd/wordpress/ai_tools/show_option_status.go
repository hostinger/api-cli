package ai_tools

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ShowOptionStatusCmd = &cobra.Command{
	Use:     "show-option-status <username> <software>",
	Short:   "Show AI option status",
	Long:    "Show the current AI option status for the Hostinger Tools plugin on the\nspecified WordPress installation. Filter by `option` to return a single\noption, or omit it to return all options.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
	Example: "  hostinger wordpress ai-tools show-option-status u123456789 1232456789",
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
		nil,
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "option", []string{"llmstxt", "web2agent"}); err != nil {
			return err
		}
		params := showOptionStatusParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingShowAIOptionStatusV1WithResponse(ctx, args[0], args[1], params)
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
	utils.RegisterArgs(ShowOptionStatusCmd, utils.StringArg("username"), utils.StringArg("software"))
	ShowOptionStatusCmd.Flags().StringP("option", "", "", "Filter the status by a single AI option. (one of: llmstxt, web2agent)")
	ShowOptionStatusCmd.RegisterFlagCompletionFunc("option", cobra.FixedCompletions([]cobra.Completion{"llmstxt", "web2agent"}, cobra.ShellCompDirectiveNoFileComp))
}

func showOptionStatusParams(cmd *cobra.Command) *client.HostingShowAIOptionStatusV1Params {
	params := &client.HostingShowAIOptionStatusV1Params{}
	if cmd.Flags().Changed("option") {
		v, _ := cmd.Flags().GetString("option")
		e := client.HostingShowAIOptionStatusV1ParamsOption(v)
		params.Option = &e
	}
	return params
}
