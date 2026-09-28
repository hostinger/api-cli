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
	Long:  "Returns the website setups started in the last 24 hours for the hosting accounts\naccessible to the authenticated client, newest first.\n\nMeant for polling right after creating a website: the website shows up in the\nwebsites list before its server-side setup has finished, and while the setup is\n`running` endpoints that operate on that website may respond with `404` or `409`.\nPoll this endpoint with the `domain` filter every 10 to 15 seconds and wait for\n`status: completed` before uploading files, deploying or creating databases.\n`failed` means the setup stopped before finishing or has not reported progress for\nover an hour. Setups older than 24 hours are not listed.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListWebsiteSetupsV1WithResponse(context.TODO(), listSetupsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListSetupsCmd.Flags().StringP("domain", "", "", "Filter by domain name (exact match)")
}

func listSetupsParams(cmd *cobra.Command) *client.HostingListWebsiteSetupsV1Params {
	params := &client.HostingListWebsiteSetupsV1Params{}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	return params
}
