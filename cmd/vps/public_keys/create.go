package public_keys

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:               "create",
	Short:             "Create public key",
	Long:              "Deprecated: account-level public keys have no direct replacement. Root SSH keys are now managed per\nvirtual machine via `POST /api/vps/v1/virtual-machines/{virtualMachineId}/ssh-keys`.\n\nAdd a new public key to your account.\n\nUse this endpoint to register SSH keys for VPS authentication.",
	Example:           "  hostinger vps public-keys create --key 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD...' --name 'My Public Key'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSCreatePublicKeyV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd)
	CreateCmd.Flags().StringP("key", "", "", "")
	CreateCmd.Flags().StringP("name", "", "", "")
	CreateCmd.MarkFlagRequired("key")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.RegisterFlagCompletionFunc("key", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	keyVal, _ := cmd.Flags().GetString("key")
	body["key"] = keyVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body, nil
}
