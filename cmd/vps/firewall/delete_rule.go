package firewall

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteRuleCmd = &cobra.Command{
	Use:     "delete-rule <firewall-id> <rule-id>",
	Short:   "Delete firewall rule",
	Long:    "Delete a specific firewall rule from a specified firewall.\n\nAny virtual machine that has this firewall activated will lose sync with the firewall\nand will have to be synced again manually.\n\nUse this endpoint to remove specific firewall rules.",
	Example: "  hostinger vps firewall delete-rule 9449049 8941182",
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
		nil,
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Delete firewall rule", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSDeleteFirewallRuleV1WithResponse(ctx, utils.StringToInt(args[0]), utils.StringToInt(args[1]))
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
	utils.RegisterArgs(DeleteRuleCmd, utils.IntArg("firewall-id"), utils.IntArg("rule-id"))
	DeleteRuleCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
