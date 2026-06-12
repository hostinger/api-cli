package firewall

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateRuleCmd = &cobra.Command{
	Use:   "update-rule <firewall-id> <rule-id>",
	Short: "Update firewall rule",
	Long:  "Update a specific firewall rule from a specified firewall.\n\nAny virtual machine that has this firewall activated will lose sync with the firewall\nand will have to be synced again manually.\n\nUse this endpoint to modify existing firewall rules.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "protocol", []string{"TCP", "UDP", "ICMP", "GRE", "any", "ESP", "AH", "ICMPv6", "SSH", "HTTP", "HTTPS", "MySQL", "PostgreSQL"})
		utils.EnumCheck(cmd, "source", []string{"any", "custom"})
		payload, err := json.Marshal(updateRuleBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSUpdateFirewallRuleV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), utils.StringToInt(args[1]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateRuleCmd.Flags().StringP("port", "", "", "Port or port range, ex: 1024:2048")
	UpdateRuleCmd.Flags().StringP("protocol", "", "", "(one of: TCP, UDP, ICMP, GRE, any, ESP, AH, ICMPv6, SSH, HTTP, HTTPS, MySQL, PostgreSQL)")
	UpdateRuleCmd.Flags().StringP("source", "", "", "(one of: any, custom)")
	UpdateRuleCmd.Flags().StringP("source-detail", "", "", "IP range, CIDR, single IP or `any`")
	UpdateRuleCmd.MarkFlagRequired("port")
	UpdateRuleCmd.MarkFlagRequired("protocol")
	UpdateRuleCmd.MarkFlagRequired("source")
	UpdateRuleCmd.MarkFlagRequired("source-detail")
}

func updateRuleBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	portVal, _ := cmd.Flags().GetString("port")
	body["port"] = portVal
	protocolVal, _ := cmd.Flags().GetString("protocol")
	body["protocol"] = protocolVal
	sourceVal, _ := cmd.Flags().GetString("source")
	body["source"] = sourceVal
	sourceDetailVal, _ := cmd.Flags().GetString("source-detail")
	body["source_detail"] = sourceDetailVal
	return body
}
