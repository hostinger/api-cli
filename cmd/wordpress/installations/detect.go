package installations

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DetectCmd = &cobra.Command{
	Use:   "detect <username>",
	Short: "Detect WordPress installations",
	Long:  "Trigger a background scan to detect WordPress installations for the account.\n\nThis operation is asynchronous: a successful response only means the scan has\nbeen queued. Poll GET /api/hosting/v1/wordpress/installations to fetch the\ndetected installations once the scan completes.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingDetectWordPressInstallationsV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
