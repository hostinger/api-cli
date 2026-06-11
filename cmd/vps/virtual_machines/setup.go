package virtual_machines

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var SetupCmd = &cobra.Command{
	Use:   "setup <virtual machine ID>",
	Short: "Setup purchased virtual machine",
	Long: `This endpoint sets up a newly purchased virtual machine that has not been set up yet, by installing the
selected operating system template in the chosen data center.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSSetupPurchasedVirtualMachineV1WithResponse(context.TODO(), utils.StringToInt(args[0]), setupRequestFromFlags(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	addSetupFlags(SetupCmd)
}

// addSetupFlags registers the virtual machine setup flags shared by the setup and purchase commands.
func addSetupFlags(cmd *cobra.Command) {
	cmd.Flags().IntP("data_center_id", "", -1, "Data center ID")
	cmd.Flags().IntP("template_id", "", -1, "Template ID")
	cmd.Flags().BoolP("enable_backups", "", false, "Enable weekly backup schedule")
	cmd.Flags().StringP("hostname", "", "", "Override default hostname of the virtual machine")
	cmd.Flags().BoolP("install_monarx", "", false, "Install Monarx malware scanner (if supported)")
	cmd.Flags().StringP("ns1", "", "", "Name server 1")
	cmd.Flags().StringP("ns2", "", "", "Name server 2")
	cmd.Flags().StringP("password", "", "", "Password for the virtual machine. If not provided, a random password will be generated")
	cmd.Flags().IntP("post_install_script_id", "", -1, "Post-install script ID")
	cmd.Flags().StringP("public_key_name", "", "", "Name of the SSH key to use")
	cmd.Flags().StringP("public_key", "", "", "Contents of the SSH key to use")

	cmd.MarkFlagRequired("data_center_id")
	cmd.MarkFlagRequired("template_id")
}

func setupRequestFromFlags(cmd *cobra.Command) client.VPSV1VirtualMachineSetupRequest {
	dataCenterId, _ := cmd.Flags().GetInt("data_center_id")
	templateId, _ := cmd.Flags().GetInt("template_id")
	hostname, _ := cmd.Flags().GetString("hostname")
	ns1, _ := cmd.Flags().GetString("ns1")
	ns2, _ := cmd.Flags().GetString("ns2")
	password, _ := cmd.Flags().GetString("password")
	postInstallScriptId, _ := cmd.Flags().GetInt("post_install_script_id")

	setup := client.VPSV1VirtualMachineSetupRequest{
		DataCenterId:        dataCenterId,
		TemplateId:          templateId,
		Hostname:            utils.StringPtrOrNil(hostname),
		Ns1:                 utils.StringPtrOrNil(ns1),
		Ns2:                 utils.StringPtrOrNil(ns2),
		Password:            utils.StringPtrOrNil(password),
		PostInstallScriptId: utils.IntPtrOrNil(postInstallScriptId),
	}

	if cmd.Flags().Changed("enable_backups") {
		enableBackups, _ := cmd.Flags().GetBool("enable_backups")
		setup.EnableBackups = &enableBackups
	}

	if cmd.Flags().Changed("install_monarx") {
		installMonarx, _ := cmd.Flags().GetBool("install_monarx")
		setup.InstallMonarx = &installMonarx
	}

	publicKeyName, _ := cmd.Flags().GetString("public_key_name")
	publicKey, _ := cmd.Flags().GetString("public_key")
	if publicKeyName != "" || publicKey != "" {
		setup.PublicKey = &struct {
			Key  *string `json:"key,omitempty"`
			Name *string `json:"name,omitempty"`
		}{
			Key:  utils.StringPtrOrNil(publicKey),
			Name: utils.StringPtrOrNil(publicKeyName),
		}
	}

	return setup
}
