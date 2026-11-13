package instances

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
	Short:             "Create instance",
	Long:              "Create a new GPU instance.\n\nPick `pool_instance_id` from the pool instance types list and `os_template_id` from the OS templates\navailable for that instance type. Your account must have enough credits to run the instance for its\nfirst hour.\n\nThe instance is created asynchronously: a successful request returns immediately, and the instance\nappears in the instance list while its `status` moves from `created` through `creating` to `running`.\n\nUse this endpoint to provision GPU compute.",
	Example:           "  hostinger gpu instances create --name my-instance-1 --os-template-id 23 --pool-instance-id 232",
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
			r, err := api.Request().GPUCreateInstanceV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	CreateCmd.Flags().StringP("name", "", "", "Instance name, unique among your instances. Letters, numbers, dashes and underscores only")
	CreateCmd.Flags().IntP("os-template-id", "", 0, "OS template ID available for the selected instance type")
	CreateCmd.Flags().IntP("pool-instance-id", "", 0, "Instance type ID from the pool instance types list")
	CreateCmd.Flags().StringSliceP("public-keys", "", nil, "SSH public keys to add to the instance")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("os-template-id")
	CreateCmd.MarkFlagRequired("pool-instance-id")
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("os-template-id", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("pool-instance-id", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("public-keys", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	osTemplateIdVal, _ := cmd.Flags().GetInt("os-template-id")
	body["os_template_id"] = osTemplateIdVal
	poolInstanceIdVal, _ := cmd.Flags().GetInt("pool-instance-id")
	body["pool_instance_id"] = poolInstanceIdVal
	if cmd.Flags().Changed("public-keys") {
		v, _ := cmd.Flags().GetStringSlice("public-keys")
		body["public_keys"] = v
	}
	return body, nil
}
