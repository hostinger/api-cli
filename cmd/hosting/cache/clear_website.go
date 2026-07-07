package cache

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ClearWebsiteCmd = &cobra.Command{
	Use:   "clear-website <username> <domain>",
	Short: "Clear website cache",
	Long:  "Permanently clears all server-side cache for the website at once. Use it when content was\nupdated and needs to be visible immediately, or after making major changes.\n\nAlso purges the Hostinger CDN cache when CDN is enabled on the website. For a WordPress\ninstallation living in a subdirectory, pass the directory query parameter to clear its cache.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingClearWebsiteCacheV1WithResponse(context.TODO(), args[0], args[1], clearWebsiteParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ClearWebsiteCmd.Flags().StringP("directory", "", "", "Directory of the website installation to clear, relative to the website root.\nDefaults to the website root.")
}

func clearWebsiteParams(cmd *cobra.Command) *client.HostingClearWebsiteCacheV1Params {
	params := &client.HostingClearWebsiteCacheV1Params{}
	if cmd.Flags().Changed("directory") {
		v, _ := cmd.Flags().GetString("directory")
		params.Directory = &v
	}
	return params
}
