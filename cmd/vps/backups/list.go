package backups

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
	Short: "Get backup list",
	Long:  `This endpoint retrieves a list of backups for a specified virtual machine.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetBackupsV1WithResponse(context.TODO(), utils.StringToInt(args[0]), backupListRequestParameters(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("page", "", 1, "Page number")
}

func backupListRequestParameters(cmd *cobra.Command) *client.VPSGetBackupsV1Params {
	pageId, _ := cmd.Flags().GetInt("page")

	return &client.VPSGetBackupsV1Params{
		Page: utils.IntPtrOrNil(pageId),
	}
}
