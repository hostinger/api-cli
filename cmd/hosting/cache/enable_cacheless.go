package cache

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var EnableCachelessCmd = &cobra.Command{
	Use:   "enable-cacheless <username> <domain>",
	Short: "Enable cacheless mode",
	Long:  "Enables development (cacheless) mode where nothing is cached, effectively turning off all\ncaching for the website. Use it while actively developing, testing changes, debugging issues,\nor when real-time updates must be visible. Disable cacheless mode afterwards to restore\nnormal caching.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingEnableCachelessModeV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
