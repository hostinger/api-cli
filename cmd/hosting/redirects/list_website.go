package redirects

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListWebsiteCmd = &cobra.Command{
	Use:   "list-website <username> <domain>",
	Short: "List website redirects",
	Long:  "Returns a paginated list of redirects configured for the selected website.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListWebsiteRedirectsV1WithResponse(context.TODO(), args[0], args[1], listWebsiteParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListWebsiteCmd.Flags().IntP("page", "", 0, "Page number")
	ListWebsiteCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listWebsiteParams(cmd *cobra.Command) *client.HostingListWebsiteRedirectsV1Params {
	params := &client.HostingListWebsiteRedirectsV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	return params
}
