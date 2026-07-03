package plugins

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListSuggestedCmd = &cobra.Command{
	Use:   "list-suggested",
	Short: "List suggested WordPress plugins",
	Long:  "List curated plugin suggestions grouped by website type.\n\nUse the returned `slug` values with\nPOST /api/hosting/v1/accounts/{username}/wordpress/{software}/plugins/install.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListSuggestedWordPressPluginsV1WithResponse(context.TODO(), listSuggestedParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListSuggestedCmd.Flags().IntP("order-id", "", 0, "Optionally scope suggestions to a specific order.")
}

func listSuggestedParams(cmd *cobra.Command) *client.HostingListSuggestedWordPressPluginsV1Params {
	params := &client.HostingListSuggestedWordPressPluginsV1Params{}
	if cmd.Flags().Changed("order-id") {
		v, _ := cmd.Flags().GetInt("order-id")
		params.OrderId = &v
	}
	return params
}
