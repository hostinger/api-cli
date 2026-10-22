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

var UpdateRuleCmd = &cobra.Command{
	Use:     "update-rule <firewall-id> <rule-id>",
	Short:   "Update firewall rule",
	Long:    "Update a specific firewall rule from a specified firewall.\n\nAny virtual machine that has this firewall activated will lose sync with the firewall\nand will have to be synced again manually.\n\nUse this endpoint to modify existing firewall rules.",
	Example: "  hostinger vps firewall update-rule 9449049 8941182 --port 443 --protocol TCP --source any --source-detail 351.15.24.0/24",
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
		if err := utils.EnumCheck(cmd, "protocol", []string{"TCP", "UDP", "ICMP", "GRE", "any", "ESP", "AH", "ICMPv6", "SSH", "HTTP", "HTTPS", "MySQL", "PostgreSQL"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "source", []string{"any", "custom"}); err != nil {
			return err
		}
		body, err := updateRuleBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSUpdateFirewallRuleV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), utils.StringToInt(args[1]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateRuleCmd, utils.IntArg("firewall-id"), utils.IntArg("rule-id"))
	UpdateRuleCmd.Flags().StringP("port", "", "", "Port or port range, ex: 1024:2048")
	UpdateRuleCmd.Flags().StringP("protocol", "", "", "(one of: TCP, UDP, ICMP, GRE, any, ESP, AH, ICMPv6, SSH, HTTP, HTTPS, MySQL, PostgreSQL)")
	UpdateRuleCmd.Flags().StringP("source", "", "", "(one of: any, custom)")
	UpdateRuleCmd.Flags().StringP("source-detail", "", "", "IP range, CIDR, single IP or `any`")
	UpdateRuleCmd.MarkFlagRequired("port")
	UpdateRuleCmd.MarkFlagRequired("protocol")
	UpdateRuleCmd.MarkFlagRequired("source")
	UpdateRuleCmd.MarkFlagRequired("source-detail")
	UpdateRuleCmd.RegisterFlagCompletionFunc("port", cobra.NoFileCompletions)
	UpdateRuleCmd.RegisterFlagCompletionFunc("protocol", cobra.FixedCompletions([]cobra.Completion{"TCP", "UDP", "ICMP", "GRE", "any", "ESP", "AH", "ICMPv6", "SSH", "HTTP", "HTTPS", "MySQL", "PostgreSQL"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateRuleCmd.RegisterFlagCompletionFunc("source", cobra.FixedCompletions([]cobra.Completion{"any", "custom"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateRuleCmd.RegisterFlagCompletionFunc("source-detail", cobra.NoFileCompletions)
}

func updateRuleBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	portVal, _ := cmd.Flags().GetString("port")
	body["port"] = portVal
	protocolVal, _ := cmd.Flags().GetString("protocol")
	body["protocol"] = protocolVal
	sourceVal, _ := cmd.Flags().GetString("source")
	body["source"] = sourceVal
	sourceDetailVal, _ := cmd.Flags().GetString("source-detail")
	body["source_detail"] = sourceDetailVal
	return body, nil
}
