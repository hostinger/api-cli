package websites

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
	Short: "List websites",
	Long:  "Retrieve a paginated list of websites (main and addon types) accessible to the authenticated client.\n\nThis endpoint returns websites from your hosting accounts as well as\nwebsites from other client hosting accounts that have shared access\nwith you.\n\nUse the available query parameters to filter results by username,\norder ID, enabled status, or domain name for more targeted results.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListWebsitesV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCmd.Flags().StringP("username", "", "", "Filter by specific username")
	ListCmd.Flags().IntP("order-id", "", 0, "Order ID")
	ListCmd.Flags().BoolP("is-enabled", "", false, "Filter by enabled status")
	ListCmd.Flags().StringP("domain", "", "", "Filter by domain name (exact match)")
}

func listParams(cmd *cobra.Command) *client.HostingListWebsitesV1Params {
	params := &client.HostingListWebsitesV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	if cmd.Flags().Changed("username") {
		v, _ := cmd.Flags().GetString("username")
		params.Username = &v
	}
	if cmd.Flags().Changed("order-id") {
		v, _ := cmd.Flags().GetInt("order-id")
		params.OrderId = &v
	}
	if cmd.Flags().Changed("is-enabled") {
		v, _ := cmd.Flags().GetBool("is-enabled")
		params.IsEnabled = &v
	}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	return params
}
