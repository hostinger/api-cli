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

var RecreateCmd = &cobra.Command{
	Use:   "recreate <virtual-machine-id>",
	Short: "Recreate virtual machine",
	Long:  "Recreate a virtual machine from scratch.\n\nThe recreation process involves reinstalling the operating system and\nresetting the virtual machine to its initial state.\nSnapshots, if there are any, will be deleted.\n\n## Password Requirements\nPassword will be checked against leaked password databases. \nRequirements for the password are:\n- At least 12 characters long\n- At least one uppercase letter\n- At least one lowercase letter\n- At least one number\n- Is not leaked publicly\n\n**This operation is irreversible and will result in the loss of all data stored on the virtual machine!**\n\nUse this endpoint to completely rebuild VPS instances with fresh OS installation.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(recreateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSRecreateVirtualMachineV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	RecreateCmd.Flags().StringP("panel-password", "", "", "Panel password for the panel-based OS template. If not provided, random password will be generated.\nIf OS does not support panel_password this field will be ignored.\nPassword will not be shown in the response.")
	RecreateCmd.Flags().StringP("password", "", "", "Root password for the virtual machine. If not provided, random password will be generated.\nPassword will not be shown in the response.")
	RecreateCmd.Flags().IntP("post-install-script-id", "", 0, "Post-install script to execute after virtual machine was recreated")
	RecreateCmd.Flags().IntP("template-id", "", 0, "Template ID")
	RecreateCmd.MarkFlagRequired("template-id")
}

func recreateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("panel-password") {
		v, _ := cmd.Flags().GetString("panel-password")
		body["panel_password"] = v
	}
	if cmd.Flags().Changed("password") {
		v, _ := cmd.Flags().GetString("password")
		body["password"] = v
	}
	if cmd.Flags().Changed("post-install-script-id") {
		v, _ := cmd.Flags().GetInt("post-install-script-id")
		body["post_install_script_id"] = v
	}
	templateIdVal, _ := cmd.Flags().GetInt("template-id")
	body["template_id"] = templateIdVal
	return body
}
