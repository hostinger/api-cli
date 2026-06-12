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

var CreateRuleCmd = &cobra.Command{
	Use:   "create-rule <firewall-id>",
	Short: "Create firewall rule",
	Long:  "Create new firewall rule for a specified firewall.\n\nBy default, the firewall drops all incoming traffic,\nwhich means you must add accept rules for all ports you want to use.\n\nAny virtual machine that has this firewall activated will lose sync with the firewall\nand will have to be synced again manually.\n\nUse this endpoint to add new security rules to firewalls.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "protocol", []string{"TCP", "UDP", "ICMP", "GRE", "any", "ESP", "AH", "ICMPv6", "SSH", "HTTP", "HTTPS", "MySQL", "PostgreSQL"})
		utils.EnumCheck(cmd, "source", []string{"any", "custom"})
		payload, err := json.Marshal(createRuleBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSCreateFirewallRuleV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateRuleCmd.Flags().StringP("port", "", "", "Port or port range, ex: 1024:2048")
	CreateRuleCmd.Flags().StringP("protocol", "", "", "(one of: TCP, UDP, ICMP, GRE, any, ESP, AH, ICMPv6, SSH, HTTP, HTTPS, MySQL, PostgreSQL)")
	CreateRuleCmd.Flags().StringP("source", "", "", "(one of: any, custom)")
	CreateRuleCmd.Flags().StringP("source-detail", "", "", "IP range, CIDR, single IP or `any`")
	CreateRuleCmd.MarkFlagRequired("port")
	CreateRuleCmd.MarkFlagRequired("protocol")
	CreateRuleCmd.MarkFlagRequired("source")
	CreateRuleCmd.MarkFlagRequired("source-detail")
}

func createRuleBody(cmd *cobra.Command) map[string]any {
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
