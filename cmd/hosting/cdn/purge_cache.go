package cdn

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var PurgeCacheCmd = &cobra.Command{
	Use:   "purge-cache <username> <domain>",
	Short: "Purge CDN cache",
	Long:  "Purges the website's content from the Hostinger CDN edge cache only. The origin server cache\nis not touched; to clear both, use `Clear website cache`.\n\nReturns 422 when CDN is not enabled on the website. Returns 404 when the domain is not a\nwebsite on the account.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingPurgeCDNCacheV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
