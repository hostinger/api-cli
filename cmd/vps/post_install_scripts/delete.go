package post_install_scripts

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <post-install script ID>",
	Short: "Delete post-install script",
	Long:  `This endpoint deletes a post-install script from your account.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSDeletePostInstallScriptV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
