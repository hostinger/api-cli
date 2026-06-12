package docker

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <virtual-machine-id> <project-name>",
	Short: "Delete project",
	Long:  "Completely removes a Docker Compose project from the virtual machine, stopping all containers and cleaning up \nassociated resources including networks, volumes, and images. \n\nThis operation is irreversible and will delete all project data. \n\nUse this when you want to permanently remove a project and free up system resources.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSDeleteProjectV1WithResponse(context.TODO(), utils.StringToInt(args[0]), args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
