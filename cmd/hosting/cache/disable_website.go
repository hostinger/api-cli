package cache

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DisableWebsiteCmd = &cobra.Command{
	Use:   "disable-website <username> <domain>",
	Short: "Disable website cache",
	Long:  "Turns off server-side caching for the website until it is enabled again. May impact performance.\nUse it when experiencing cache-related issues; to temporarily bypass caching while developing\nor debugging, prefer enabling cacheless mode instead.\n\nDoes nothing if caching is already disabled.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingDisableWebsiteCacheV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
