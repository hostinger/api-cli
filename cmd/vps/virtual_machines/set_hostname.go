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

var SetHostnameCmd = &cobra.Command{
	Use:   "set-hostname <virtual-machine-id>",
	Short: "Set hostname",
	Long:  "Set hostname for a specified virtual machine.\n\nChanging hostname does not update PTR record automatically.\nIf you want your virtual machine to be reachable by a hostname, \nyou need to point your domain A/AAAA records to virtual machine IP as well.\n\nUse this endpoint to configure custom hostnames for VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(setHostnameBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSSetHostnameV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SetHostnameCmd.Flags().StringP("hostname", "", "", "")
	SetHostnameCmd.MarkFlagRequired("hostname")
}

func setHostnameBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	hostnameVal, _ := cmd.Flags().GetString("hostname")
	body["hostname"] = hostnameVal
	return body
}
