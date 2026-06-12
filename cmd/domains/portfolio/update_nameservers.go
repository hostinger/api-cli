package portfolio

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var UpdateNameserversCmd = &cobra.Command{
	Use:   "update-nameservers <domain>",
	Short: "Update domain nameservers",
	Long:  "Set nameservers for a specified domain.\n\nBe aware, that improper nameserver configuration can lead to the domain being unresolvable or unavailable.\n\nUse this endpoint to configure custom DNS hosting for domains.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateNameserversBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().DomainsUpdateDomainNameserversV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateNameserversCmd.Flags().StringP("ns1", "", "", "First name server")
	UpdateNameserversCmd.Flags().StringP("ns2", "", "", "Second name server")
	UpdateNameserversCmd.Flags().StringP("ns3", "", "", "Third name server")
	UpdateNameserversCmd.Flags().StringP("ns4", "", "", "Fourth name server")
	UpdateNameserversCmd.MarkFlagRequired("ns1")
	UpdateNameserversCmd.MarkFlagRequired("ns2")
}

func updateNameserversBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	ns1Val, _ := cmd.Flags().GetString("ns1")
	body["ns1"] = ns1Val
	ns2Val, _ := cmd.Flags().GetString("ns2")
	body["ns2"] = ns2Val
	if cmd.Flags().Changed("ns3") {
		v, _ := cmd.Flags().GetString("ns3")
		body["ns3"] = v
	}
	if cmd.Flags().Changed("ns4") {
		v, _ := cmd.Flags().GetString("ns4")
		body["ns4"] = v
	}
	return body
}
