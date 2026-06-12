package datacenters

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available datacenters",
	Long:  "Retrieve a list of datacenters available for setting up hosting plans\nbased on available datacenter capacity and hosting plan of your order.\nThe first item in the list is the best match for your specific order\nrequirements.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListAvailableDatacentersV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("order-id", "", 0, "Order ID")
	ListCmd.MarkFlagRequired("order-id")
}

func listParams(cmd *cobra.Command) *client.HostingListAvailableDatacentersV1Params {
	params := &client.HostingListAvailableDatacentersV1Params{}
	orderIdVal, _ := cmd.Flags().GetInt("order-id")
	params.OrderId = orderIdVal
	return params
}
