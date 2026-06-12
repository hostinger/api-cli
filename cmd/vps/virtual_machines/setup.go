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

var SetupCmd = &cobra.Command{
	Use:   "setup <virtual-machine-id>",
	Short: "Setup purchased virtual machine",
	Long:  "Setup newly purchased virtual machine with `initial` state.\n\nUse this endpoint to configure and initialize purchased VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(setupBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().VPSSetupPurchasedVirtualMachineV1WithBodyWithResponse(context.TODO(), utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SetupCmd.Flags().IntP("data-center-id", "", 0, "Data center ID")
	SetupCmd.Flags().BoolP("enable-backups", "", true, "Enable weekly backup schedule")
	SetupCmd.Flags().StringP("hostname", "", "", "Override default hostname of the virtual machine")
	SetupCmd.Flags().BoolP("install-monarx", "", false, "Install Monarx malware scanner (if supported)")
	SetupCmd.Flags().StringP("ns1", "", "", "Name server 1")
	SetupCmd.Flags().StringP("ns2", "", "", "Name server 2")
	SetupCmd.Flags().StringP("password", "", "", "Password for the virtual machine. If not provided, random password will be generated.\nPassword will not be shown in the response.")
	SetupCmd.Flags().IntP("post-install-script-id", "", 0, "Post-install script ID")
	SetupCmd.Flags().StringP("public-key", "", "", "Use SSH key (JSON)")
	SetupCmd.Flags().IntP("template-id", "", 0, "Template ID")
	SetupCmd.MarkFlagRequired("data-center-id")
	SetupCmd.MarkFlagRequired("template-id")
}

func setupBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	dataCenterIdVal, _ := cmd.Flags().GetInt("data-center-id")
	body["data_center_id"] = dataCenterIdVal
	if cmd.Flags().Changed("enable-backups") {
		v, _ := cmd.Flags().GetBool("enable-backups")
		body["enable_backups"] = v
	}
	if cmd.Flags().Changed("hostname") {
		v, _ := cmd.Flags().GetString("hostname")
		body["hostname"] = v
	}
	if cmd.Flags().Changed("install-monarx") {
		v, _ := cmd.Flags().GetBool("install-monarx")
		body["install_monarx"] = v
	}
	if cmd.Flags().Changed("ns1") {
		v, _ := cmd.Flags().GetString("ns1")
		body["ns1"] = v
	}
	if cmd.Flags().Changed("ns2") {
		v, _ := cmd.Flags().GetString("ns2")
		body["ns2"] = v
	}
	if cmd.Flags().Changed("password") {
		v, _ := cmd.Flags().GetString("password")
		body["password"] = v
	}
	if cmd.Flags().Changed("post-install-script-id") {
		v, _ := cmd.Flags().GetInt("post-install-script-id")
		body["post_install_script_id"] = v
	}
	if cmd.Flags().Changed("public-key") {
		v, _ := cmd.Flags().GetString("public-key")
		body["public_key"] = utils.JSONValue(v, "public-key")
	}
	templateIdVal, _ := cmd.Flags().GetInt("template-id")
	body["template_id"] = templateIdVal
	return body
}
