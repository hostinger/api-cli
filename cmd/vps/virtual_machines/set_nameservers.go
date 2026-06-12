package virtual_machines

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

var SetNameserversCmd = &cobra.Command{
	Use:   "set-nameservers <virtual-machine-id>",
	Short: "Set nameservers",
	Long:  "Set nameservers for a specified virtual machine.\n\nBe aware, that improper nameserver configuration can lead to the virtual\nmachine being unable to resolve domain names.\n\nUse this endpoint to configure custom DNS resolvers for VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(setNameserversBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSSetNameserversV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SetNameserversCmd.Flags().StringP("ns1", "", "", "")
	SetNameserversCmd.Flags().StringP("ns2", "", "", "")
	SetNameserversCmd.Flags().StringP("ns3", "", "", "")
	SetNameserversCmd.MarkFlagRequired("ns1")
}

func setNameserversBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	ns1Val, _ := cmd.Flags().GetString("ns1")
	body["ns1"] = ns1Val
	if cmd.Flags().Changed("ns2") {
		v, _ := cmd.Flags().GetString("ns2")
		body["ns2"] = v
	}
	if cmd.Flags().Changed("ns3") {
		v, _ := cmd.Flags().GetString("ns3")
		body["ns3"] = v
	}
	return body
}
