package websites

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListSetupsCmd = &cobra.Command{
	Use:   "list-setups",
	Short: "List website setups",
	Long:  "Returns the website setups started in the last 24 hours for the hosting accounts\naccessible to the authenticated client, newest first. Narrow the list with the\n`order_id`, `subscription_id` or `domain` filters.\n\nMeant for polling right after creating a website or starting a website setup: the\nwebsite shows up in the websites list before its server-side setup has finished, and\nwhile the setup is `running` endpoints that operate on that website may respond with\n`404` or `409`. Poll this endpoint with the `domain` filter every 10 to 15 seconds and\nwait for `status: completed` before uploading files, deploying or creating databases.\n`failed` means the setup stopped before finishing or has not reported progress for\nover an hour. Setups older than 24 hours are not listed.\n\n`type` is the website type the setup was started with (`wordpress`, `headless_wordpress`,\n`headless_ecommerce`, `headless_pocketbase`), or `null` for an empty website.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListWebsiteSetupsV1WithResponse(context.TODO(), listSetupsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListSetupsCmd.Flags().IntP("order-id", "", 0, "Order ID")
	ListSetupsCmd.Flags().StringP("subscription-id", "", "", "Filter by hosting order subscription ID")
	ListSetupsCmd.Flags().StringP("domain", "", "", "Filter by domain name (exact match)")
}

func listSetupsParams(cmd *cobra.Command) *client.HostingListWebsiteSetupsV1Params {
	params := &client.HostingListWebsiteSetupsV1Params{}
	if cmd.Flags().Changed("order-id") {
		v, _ := cmd.Flags().GetInt("order-id")
		params.OrderId = &v
	}
	if cmd.Flags().Changed("subscription-id") {
		v, _ := cmd.Flags().GetString("subscription-id")
		params.SubscriptionId = &v
	}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	return params
}
