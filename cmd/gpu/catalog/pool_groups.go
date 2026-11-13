package catalog

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PoolGroupsCmd = &cobra.Command{
	Use:               "pool-groups",
	Short:             "Get pool groups",
	Long:              "Retrieve GPU pool groups, ordered by name.\n\nEach group bundles the pools of one GPU model and lists its currently available pools,\nwhether any of them is available and the lowest hourly price, in credits, among its instance types.\nThe price is returned even when the group is out of stock.\n\nUse this endpoint to show which GPU models can be ordered and from what price.",
	Example:           "  hostinger gpu catalog pool-groups",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GPUGetPoolGroupsV1WithResponse(ctx)
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
	utils.RegisterArgs(PoolGroupsCmd)
}
