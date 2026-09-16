package git

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var AutoDeploymentSettingsCmd = &cobra.Command{
	Use:   "auto-deployment-settings <username> <domain>",
	Short: "Get Git auto-deployment settings",
	Long:  "Returns the Git auto-deployment settings of the website: which repository and branch deploy\ninto which directory, and whether pushes trigger a deployment. `is_enabled` false keeps the\nrepository link but ignores pushes.\n\nWhen the website has no auto-deployment configured every field is null. Save settings with\n`Update Git auto-deployment settings`.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetGitAutoDeploymentSettingsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
