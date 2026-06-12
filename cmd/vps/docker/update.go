package docker

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update <virtual-machine-id> <project-name>",
	Short: "Update project",
	Long:  "Updates a Docker Compose project by pulling the latest image versions and\nrecreating containers with new configurations.\n\nThis operation preserves data volumes while applying changes from the compose file. \n\nUse this to deploy application updates, apply configuration changes, or\nrefresh container images to their latest versions.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSUpdateProjectV1WithResponse(context.TODO(), utils.StringToInt(args[0]), args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
