package cache

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DisableCachelessCmd = &cobra.Command{
	Use:   "disable-cacheless <username> <domain>",
	Short: "Disable cacheless mode",
	Long:  "Turns off development (cacheless) mode and returns the website to normal caching. Use it after\nfinishing development work to restore the performance benefits of caching.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingDisableCachelessModeV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
