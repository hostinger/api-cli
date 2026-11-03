package catalog

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PoolCmd = &cobra.Command{
	Use:               "pool",
	Short:             "Get pool list",
	Long:              "Retrieve all available GPU pools, ordered by name.\n\nA pool serves GPU accelerators of one model from a single region.\n\nUse this endpoint to choose a pool before listing its instance types.",
	Example:           "  hostinger gpu catalog pool",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GPUGetPoolListV1WithResponse(ctx)
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
	utils.RegisterArgs(PoolCmd)
}
