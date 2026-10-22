package datacenters

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list",
	Short:             "List available datacenters",
	Long:              "Retrieve a list of datacenters available for setting up hosting plans\nbased on available datacenter capacity and hosting plan of your order.\nThe first item in the list is the best match for your specific order\nrequirements.",
	Example:           "  hostinger hosting datacenters list --order-id 123",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListAvailableDatacentersV1WithResponse(ctx, params)
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
	utils.RegisterArgs(ListCmd)
	ListCmd.Flags().IntP("order-id", "", 0, "Order ID")
	ListCmd.MarkFlagRequired("order-id")
	ListCmd.RegisterFlagCompletionFunc("order-id", cobra.NoFileCompletions)
}

func listParams(cmd *cobra.Command) *client.HostingListAvailableDatacentersV1Params {
	params := &client.HostingListAvailableDatacentersV1Params{}
	orderIdVal, _ := cmd.Flags().GetInt("order-id")
	params.OrderId = orderIdVal
	return params
}
