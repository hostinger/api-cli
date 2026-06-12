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

var SetPanelPasswordCmd = &cobra.Command{
	Use:   "set-panel-password <virtual-machine-id>",
	Short: "Set panel password",
	Long:  "Set panel password for a specified virtual machine.\n\nIf virtual machine does not use panel OS, the request will still be processed without any effect.\nRequirements for password are same as in the [recreate virtual machine\nendpoint](/#tag/vps-virtual-machine/POST/api/vps/v1/virtual-machines/{virtualMachineId}/recreate).\n\nUse this endpoint to configure control panel access credentials for VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(setPanelPasswordBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSSetPanelPasswordV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SetPanelPasswordCmd.Flags().StringP("password", "", "", "Panel password for the virtual machine")
	SetPanelPasswordCmd.MarkFlagRequired("password")
}

func setPanelPasswordBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body
}
