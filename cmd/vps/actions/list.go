package actions

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var ListCmd = &cobra.Command{
	Use:   "list <virtual machine ID>",
	Short: "Get action list",
	Long: `Actions are operations or events that have been executed on the virtual machine, such as starting, stopping, 
or modifying the machine. This endpoint allows you to view the history of these actions, providing details about each 
action, such as the action name, timestamp, and status.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetActionsV1WithResponse(context.TODO(), utils.StringToInt(args[0]), listRequestParameters(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("page", "", 1, "Page number")
}

func listRequestParameters(cmd *cobra.Command) *client.VPSGetActionsV1Params {
	pageId, _ := cmd.Flags().GetInt("page")

	return &client.VPSGetActionsV1Params{
		Page: utils.IntPtrOrNil(pageId),
	}
}
