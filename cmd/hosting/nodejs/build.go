package nodejs

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var BuildCmd = &cobra.Command{
	Use:   "build <username> <domain> <uuid>",
	Short: "Get Node.js build details",
	Long:  "Returns one build by UUID: its state (`pending`, `running`, `completed`, `failed`), the\noptions it ran with and timestamps. Poll this while a build is pending or running. When it\nis failed, read `Get NodeJS build logs` and `Analyse failed Node.js build` for the cause.\nReturns 404 when the UUID does not belong to a build of this website.",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetNodeJsBuildDetailsV1WithResponse(context.TODO(), args[0], args[1], uuid.MustParse(args[2]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
