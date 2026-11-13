package instances

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var MetricsCmd = &cobra.Command{
	Use:     "metrics <instance-id>",
	Short:   "Get instance metrics",
	Long:    "Retrieve utilisation metrics of a specified GPU instance for the last hour.\n\nEach data point contains CPU, memory and GPU utilisation, used GPU memory and GPU throughput.\nThe list is empty until the instance is provisioned.\n\nUse this endpoint to monitor GPU instance load.",
	Example: "  hostinger gpu instances metrics 6325",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().GPUGetInstancesV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GPUGetInstanceMetricsV1WithResponse(ctx, utils.StringToInt(args[0]))
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
	utils.RegisterArgs(MetricsCmd, utils.IntArg("instance-id"))
}
