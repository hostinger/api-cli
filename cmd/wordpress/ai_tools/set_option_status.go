package ai_tools

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

var SetOptionStatusCmd = &cobra.Command{
	Use:     "set-option-status [username] <software>",
	Short:   "Set AI option status",
	Long:    "Enable or disable an AI option for the Hostinger Tools plugin on the specified\nWordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThe hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger wordpress ai-tools set-option-status u123456789 1232456789 --enable=true --option llmstxt",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "username",
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
		username, args, err := site.ResolveUsername(cmd, args, 1)
		if err != nil {
			return api.CommandError(err)
		}
		if err := utils.EnumCheck(cmd, "option", []string{"llmstxt", "web2agent"}); err != nil {
			return err
		}
		body, err := setOptionStatusBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingSetAIOptionStatusV1WithBodyWithResponse(ctx, username, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(SetOptionStatusCmd, 1, utils.StringArg("username"), utils.StringArg("software"))
	SetOptionStatusCmd.Flags().BoolP("enable", "", false, "Enable (true) or disable (false) the AI option.")
	SetOptionStatusCmd.Flags().StringP("option", "", "", "AI option name (one of: llmstxt, web2agent)")
	SetOptionStatusCmd.MarkFlagRequired("enable")
	SetOptionStatusCmd.MarkFlagRequired("option")
	SetOptionStatusCmd.RegisterFlagCompletionFunc("enable", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	SetOptionStatusCmd.RegisterFlagCompletionFunc("option", cobra.FixedCompletions([]cobra.Completion{"llmstxt", "web2agent"}, cobra.ShellCompDirectiveNoFileComp))
}

func setOptionStatusBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	enableVal, _ := cmd.Flags().GetBool("enable")
	body["enable"] = enableVal
	optionVal, _ := cmd.Flags().GetString("option")
	body["option"] = optionVal
	return body, nil
}
