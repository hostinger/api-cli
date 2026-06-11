package docker

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var RestartCmd = &cobra.Command{
	Use:   "restart <virtual machine ID> <project name>",
	Short: "Restart Docker project",
	Long:  `This endpoint restarts a specified Docker Compose project on a virtual machine.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSRestartProjectV1WithResponse(context.TODO(), utils.StringToInt(args[0]), args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
