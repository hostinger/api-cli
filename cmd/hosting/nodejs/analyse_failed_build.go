package nodejs

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var AnalyseFailedBuildCmd = &cobra.Command{
	Use:   "analyse-failed-build <username> <domain> <uuid>",
	Short: "Analyse failed Node.js build",
	Long:  "Returns an AI analysis of why a build failed and how to fix it, based on the build logs,\nthe project file list and package.json. Only builds in the `failed` state can be analysed;\nany other state returns 422. When no analysis could be produced both `analysis` and\n`solution` are null, in which case read `Get NodeJS build logs` instead.\n\nEach call runs the analysis again, so call it once per failed build and keep the result.\nLimited to 5 calls per minute per API client (429 above that).",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingAnalyseFailedNodeJsBuildV1WithResponse(context.TODO(), args[0], args[1], uuid.MustParse(args[2]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
