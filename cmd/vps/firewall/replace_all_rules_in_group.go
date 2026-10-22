package firewall

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ReplaceAllRulesInGroupCmd = &cobra.Command{
	Use:     "replace-all-rules-in-group <firewall-id>",
	Short:   "Replace all firewall rules in group",
	Long:    "Replaces all firewall rules within a specified firewall group with the provided set of rules\nin a single atomic operation, instead of creating or deleting rules one by one.\n\nAny virtual machine using this firewall group will need to be synchronized after replacing rules;\npass the \"sync\" parameter to trigger synchronization immediately.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --rules\n    [\n      {\n        \"port\": \"443\",\n        \"protocol\": \"TCP\",\n        \"source\": \"any\",\n        \"source_detail\": \"351.15.24.0/24\"\n      }\n    ]",
	Example: "  hostinger vps firewall replace-all-rules-in-group 9449049 --rules @rules.json",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().VPSGetFirewallListV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := replaceAllRulesInGroupBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSReplaceAllFirewallRulesInGroupV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ReplaceAllRulesInGroupCmd, utils.IntArg("firewall-id"))
	ReplaceAllRulesInGroupCmd.Flags().StringP("rules", "", "", "The complete set of firewall rules that atomically replaces all existing rules in the group (JSON, @file or -; shape in the help above)")
	ReplaceAllRulesInGroupCmd.Flags().BoolP("sync", "", false, "Synchronize the firewall group to all its virtual machines after replacing the rules")
	ReplaceAllRulesInGroupCmd.MarkFlagRequired("rules")
	ReplaceAllRulesInGroupCmd.RegisterFlagCompletionFunc("rules", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	ReplaceAllRulesInGroupCmd.RegisterFlagCompletionFunc("sync", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func replaceAllRulesInGroupBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	rulesVal, _ := cmd.Flags().GetString("rules")
	rulesValJSON, err := utils.JSONValue(rulesVal, "rules")
	if err != nil {
		return nil, err
	}
	body["rules"] = rulesValJSON
	if cmd.Flags().Changed("sync") {
		v, _ := cmd.Flags().GetBool("sync")
		body["sync"] = v
	}
	return body, nil
}
