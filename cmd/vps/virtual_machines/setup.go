package virtual_machines

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SetupCmd = &cobra.Command{
	Use:     "setup <virtual-machine-id>",
	Short:   "Setup purchased virtual machine",
	Long:    "Setup newly purchased virtual machine with `initial` state.\n\nUse this endpoint to configure and initialize purchased VPS instances.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --public-key\n    {\n      \"key\": \"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC2X...\",\n      \"name\": \"my-key\"\n    }",
	Example: "  hostinger vps virtual-machines setup 1268054 --data-center-id 19 --template-id 1130",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().VPSGetVirtualMachinesV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := setupBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSSetupPurchasedVirtualMachineV1WithBodyWithResponse(ctx, utils.StringToInt(args[0]), "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(SetupCmd, utils.IntArg("virtual-machine-id"))
	SetupCmd.Flags().IntP("data-center-id", "", 0, "Data center ID")
	SetupCmd.Flags().BoolP("enable-backups", "", true, "Enable weekly backup schedule")
	SetupCmd.Flags().StringP("hostname", "", "", "Override default hostname of the virtual machine")
	SetupCmd.Flags().BoolP("install-monarx", "", false, "Install Monarx malware scanner (if supported)")
	SetupCmd.Flags().StringP("ns1", "", "", "Name server 1")
	SetupCmd.Flags().StringP("ns2", "", "", "Name server 2")
	SetupCmd.Flags().StringP("password", "", "", "Password for the virtual machine. If not provided, random password will be generated. Password will not be shown in the response.")
	SetupCmd.Flags().IntP("post-install-script-id", "", 0, "Post-install script ID")
	SetupCmd.Flags().StringP("public-key", "", "", "Use SSH key (JSON, @file or -; shape in the help above)")
	SetupCmd.Flags().IntP("template-id", "", 0, "Template ID")
	SetupCmd.MarkFlagRequired("data-center-id")
	SetupCmd.MarkFlagRequired("template-id")
	SetupCmd.RegisterFlagCompletionFunc("data-center-id", cobra.NoFileCompletions)
	SetupCmd.RegisterFlagCompletionFunc("enable-backups", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	SetupCmd.RegisterFlagCompletionFunc("hostname", cobra.NoFileCompletions)
	SetupCmd.RegisterFlagCompletionFunc("install-monarx", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	SetupCmd.RegisterFlagCompletionFunc("ns1", cobra.NoFileCompletions)
	SetupCmd.RegisterFlagCompletionFunc("ns2", cobra.NoFileCompletions)
	SetupCmd.RegisterFlagCompletionFunc("password", cobra.NoFileCompletions)
	SetupCmd.RegisterFlagCompletionFunc("post-install-script-id", cobra.NoFileCompletions)
	SetupCmd.RegisterFlagCompletionFunc("public-key", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	SetupCmd.RegisterFlagCompletionFunc("template-id", cobra.NoFileCompletions)
}

func setupBody(cmd *cobra.Command) (map[string]any, error) {
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
		publicKeyVal, _ := cmd.Flags().GetString("public-key")
		publicKeyValJSON, err := utils.JSONValue(publicKeyVal, "public-key")
		if err != nil {
			return nil, err
		}
		body["public_key"] = publicKeyValJSON
	}
	templateIdVal, _ := cmd.Flags().GetInt("template-id")
	body["template_id"] = templateIdVal
	return body, nil
}
