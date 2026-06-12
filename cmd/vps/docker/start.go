package docker

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start <virtual-machine-id> <project-name>",
	Short: "Start project",
	Long:  "Starts all services in a Docker Compose project that are currently stopped. \n\nThis operation brings up containers in the correct dependency order as defined in the compose file. \n\nUse this to resume a project that was previously stopped or to start services after a system reboot.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSStartProjectV1WithResponse(context.TODO(), utils.StringToInt(args[0]), args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
