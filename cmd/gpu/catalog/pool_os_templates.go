package catalog

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PoolOsTemplatesCmd = &cobra.Command{
	Use:               "pool-os-templates <pool-id> <pool-instance-id>",
	Short:             "Get pool OS templates",
	Long:              "Retrieve the OS templates that can be installed on a specified GPU pool instance type.\n\nUse this endpoint to choose the `os_template_id` for a new GPU instance.",
	Example:           "  hostinger gpu catalog pool-os-templates 12 232",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GPUGetPoolOSTemplatesV1WithResponse(ctx, utils.StringToInt(args[0]), utils.StringToInt(args[1]))
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
	utils.RegisterArgs(PoolOsTemplatesCmd, utils.IntArg("pool-id"), utils.IntArg("pool-instance-id"))
}
