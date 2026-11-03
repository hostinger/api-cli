package catalog

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PoolInstanceTypesCmd = &cobra.Command{
	Use:               "pool-instance-types <pool-id>",
	Short:             "Get pool instance types",
	Long:              "Retrieve the instance types offered by a specified GPU pool, ordered by CPU core count.\n\nEach instance type describes its CPU, memory and storage, and its hourly price.\nOnly instance types that can currently be ordered are returned.\n\nUse this endpoint to choose the `pool_instance_id` for a new GPU instance.",
	Example:           "  hostinger gpu catalog pool-instance-types 12",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GPUGetPoolInstanceTypesV1WithResponse(ctx, utils.StringToInt(args[0]))
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
	utils.RegisterArgs(PoolInstanceTypesCmd, utils.IntArg("pool-id"))
}
