package docker

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <virtual-machine-id> <project-name>",
	Short: "Get project contents",
	Long:  "Retrieves the complete project information including the docker-compose.yml\nfile contents, project metadata, and current deployment status.\n\nThis endpoint provides the full configuration and state details of a specific Docker Compose project. \n\nUse this to inspect project settings, review the compose file, or check the overall project health.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetProjectContentsV1WithResponse(context.TODO(), utils.StringToInt(args[0]), args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
