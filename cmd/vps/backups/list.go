package backups

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <virtual-machine-id>",
	Short: "Get backups",
	Long:  "Retrieve backups for a specified virtual machine.\n\nUse this endpoint to view available backup points for VPS data recovery.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetBackupsV1WithResponse(context.TODO(), utils.StringToInt(args[0]), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("page", "", 0, "Page number")
}

func listParams(cmd *cobra.Command) *client.VPSGetBackupsV1Params {
	params := &client.VPSGetBackupsV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	return params
}
