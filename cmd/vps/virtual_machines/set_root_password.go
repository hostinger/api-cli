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

var SetRootPasswordCmd = &cobra.Command{
	Use:   "set-root-password <virtual-machine-id>",
	Short: "Set root password",
	Long:  "Set root password for a specified virtual machine.\n\nRequirements for password are same as in the [recreate virtual machine\nendpoint](/#tag/vps-virtual-machine/POST/api/vps/v1/virtual-machines/{virtualMachineId}/recreate).\n\nUse this endpoint to update administrator credentials for VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(setRootPasswordBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSSetRootPasswordV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SetRootPasswordCmd.Flags().StringP("password", "", "", "Root password for the virtual machine")
	SetRootPasswordCmd.MarkFlagRequired("password")
}

func setRootPasswordBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body
}
