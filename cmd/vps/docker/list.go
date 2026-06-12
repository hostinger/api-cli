package docker

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <virtual-machine-id>",
	Short: "Get project list",
	Long:  "Retrieves a list of all Docker Compose projects currently deployed on the virtual machine. \n\nThis endpoint returns basic information about each project including name,\nstatus, file path and list of containers with details about their names,\nimage, status, health and ports. Container stats are omitted in this\nendpoint. If you need to get detailed information about container with\nstats included, use the `Get project containers` endpoint.\n\nUse this to get an overview of all Docker projects on your VPS instance.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetProjectListV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
