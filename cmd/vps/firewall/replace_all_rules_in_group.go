package firewall

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ReplaceAllRulesInGroupCmd = &cobra.Command{
	Use:   "replace-all-rules-in-group <firewall-id>",
	Short: "Replace all firewall rules in group",
	Long:  "Replaces all firewall rules within a specified firewall group with the provided set of rules\nin a single atomic operation, instead of creating or deleting rules one by one.\n\nAny virtual machine using this firewall group will need to be synchronized after replacing rules;\npass the \"sync\" query parameter to trigger synchronization immediately.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(replaceAllRulesInGroupBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSReplaceAllFirewallRulesInGroupV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), replaceAllRulesInGroupParams(cmd), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ReplaceAllRulesInGroupCmd.Flags().BoolP("sync", "", false, "Synchronize the firewall group to all its virtual machines after replacing the rules")
	ReplaceAllRulesInGroupCmd.Flags().StringP("rules", "", "", "The complete set of firewall rules that atomically replaces all existing rules in the group (JSON)")
	ReplaceAllRulesInGroupCmd.MarkFlagRequired("rules")
}

func replaceAllRulesInGroupParams(cmd *cobra.Command) *client.VPSReplaceAllFirewallRulesInGroupV1Params {
	params := &client.VPSReplaceAllFirewallRulesInGroupV1Params{}
	if cmd.Flags().Changed("sync") {
		v, _ := cmd.Flags().GetBool("sync")
		params.Sync = &v
	}
	return params
}

func replaceAllRulesInGroupBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	rulesVal, _ := cmd.Flags().GetString("rules")
	body["rules"] = utils.JSONValue(rulesVal, "rules")
	return body
}
