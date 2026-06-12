package docker

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ContainersCmd = &cobra.Command{
	Use:   "containers <virtual-machine-id> <project-name>",
	Short: "Get project containers",
	Long:  "Retrieves a list of all containers belonging to a specific Docker Compose project on the virtual machine. \n\nThis endpoint returns detailed information about each container including\ntheir current status, port mappings, and runtime configuration.\n\nUse this to monitor the health and state of all services within your Docker Compose project.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetProjectContainersV1WithResponse(context.TODO(), utils.StringToInt(args[0]), args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
