package themes

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
	Short: "List WordPress themes",
	Long:  "List WordPress themes available to install.\n\nUse the returned `slug` values with\nPOST /api/hosting/v1/accounts/{username}/wordpress/{software}/themes/install.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListWordPressThemesV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("order-id", "", 0, "Optionally scope themes to a specific order.")
	ListCmd.Flags().StringP("search", "", "", "Search term to match against theme names.")
}

func listParams(cmd *cobra.Command) *client.HostingListWordPressThemesV1Params {
	params := &client.HostingListWordPressThemesV1Params{}
	if cmd.Flags().Changed("order-id") {
		v, _ := cmd.Flags().GetInt("order-id")
		params.OrderId = &v
	}
	if cmd.Flags().Changed("search") {
		v, _ := cmd.Flags().GetString("search")
		params.Search = &v
	}
	return params
}
