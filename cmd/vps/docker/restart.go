package docker

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RestartCmd = &cobra.Command{
	Use:   "restart <virtual-machine-id> <project-name>",
	Short: "Restart project",
	Long:  "Restarts all services in a Docker Compose project by stopping and starting\ncontainers in the correct dependency order.\n\nThis operation preserves data volumes and network configurations while refreshing the running containers. \n\nUse this to apply configuration changes or recover from service failures.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSRestartProjectV1WithResponse(context.TODO(), utils.StringToInt(args[0]), args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
