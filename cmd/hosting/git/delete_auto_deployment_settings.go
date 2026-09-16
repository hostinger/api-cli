package git

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteAutoDeploymentSettingsCmd = &cobra.Command{
	Use:   "delete-auto-deployment-settings <username> <domain>",
	Short: "Delete Git auto-deployment settings",
	Long:  "Removes the Git auto-deployment settings of the website. Files already deployed stay on the\nwebsite; pushes stop deploying until settings are saved again. Succeeds also when nothing is\nconfigured.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingDeleteGitAutoDeploymentSettingsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
