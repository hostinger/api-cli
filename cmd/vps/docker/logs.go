package docker

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var LogsCmd = &cobra.Command{
	Use:   "logs <virtual-machine-id> <project-name>",
	Short: "Get project logs",
	Long:  "Retrieves aggregated log entries from all services within a Docker Compose project. \n\nThis endpoint returns recent log output from each container, organized by service name with timestamps. \nThe response contains the last 300 log entries across all services. \n\nUse this for debugging, monitoring application behavior, and\ntroubleshooting issues across your entire project stack.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetProjectLogsV1WithResponse(context.TODO(), utils.StringToInt(args[0]), args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
