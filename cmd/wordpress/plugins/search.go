package plugins

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search WordPress plugins",
	Long:  "Search the WordPress.org plugin directory for plugins available to install.\n\nUse the returned `slug` values with\nPOST /api/hosting/v1/accounts/{username}/wordpress/{software}/plugins/install.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingSearchWordPressPluginsV1WithResponse(context.TODO(), searchParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SearchCmd.Flags().StringP("search", "", "", "Search term to match against plugin names. Minimum 3 characters.")
	SearchCmd.MarkFlagRequired("search")
}

func searchParams(cmd *cobra.Command) *client.HostingSearchWordPressPluginsV1Params {
	params := &client.HostingSearchWordPressPluginsV1Params{}
	searchVal, _ := cmd.Flags().GetString("search")
	params.Search = searchVal
	return params
}
