package cache

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var EnableWebsiteCmd = &cobra.Command{
	Use:   "enable-website <username> <domain>",
	Short: "Enable website cache",
	Long:  "Turns on server-side caching for the website for better performance. Use it for faster page\nloads, reduced server load, or improved user experience. Recommended for production websites.\n\nDoes nothing if caching is already enabled.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingEnableWebsiteCacheV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
