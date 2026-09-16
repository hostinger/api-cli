package git

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var UpdateAutoDeploymentSettingsCmd = &cobra.Command{
	Use:   "update-auto-deployment-settings <username> <domain>",
	Short: "Update Git auto-deployment settings",
	Long:  "Creates or replaces the Git auto-deployment settings of the website: repository, branch, the\ndirectory under the document root to deploy into, and `is_enabled`. Send the full set;\n`is_enabled` defaults to true and `directory` to the document root. `installation_uuid` must\nbe an installation from `List Git installations` that belongs to the same customer as the\nwebsite.\n\nFor PHP and static websites, saving with `is_enabled` true deploys the branch right away and\nevery later push to that branch deploys again. For Node.js and Website Builder websites saving\ndoes not clone anything. On a Node.js website start the first deploy with\n`Start Node.js build` using `source_type` `git`; pushes then trigger new builds with the build\nsettings stored for the website.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateAutoDeploymentSettingsBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingUpdateGitAutoDeploymentSettingsV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateAutoDeploymentSettingsCmd.Flags().StringP("branch", "", "", "Branch to deploy")
	UpdateAutoDeploymentSettingsCmd.Flags().StringP("directory", "", "", "Subdirectory under the website document root to deploy into. Empty, null or omitted means\nthe document root.")
	UpdateAutoDeploymentSettingsCmd.Flags().StringP("installation-uuid", "", "", "Active Git installation from `List Git installations`")
	UpdateAutoDeploymentSettingsCmd.Flags().BoolP("is-enabled", "", true, "Whether pushes to the branch deploy automatically")
	UpdateAutoDeploymentSettingsCmd.Flags().StringP("owner", "", "", "Repository owner login, as returned by `List Git installation repositories`. GitLab group\npaths use slashes.")
	UpdateAutoDeploymentSettingsCmd.Flags().StringP("repository", "", "", "Repository name without the .git suffix")
	UpdateAutoDeploymentSettingsCmd.MarkFlagRequired("branch")
	UpdateAutoDeploymentSettingsCmd.MarkFlagRequired("installation-uuid")
	UpdateAutoDeploymentSettingsCmd.MarkFlagRequired("owner")
	UpdateAutoDeploymentSettingsCmd.MarkFlagRequired("repository")
}

func updateAutoDeploymentSettingsBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	branchVal, _ := cmd.Flags().GetString("branch")
	body["branch"] = branchVal
	if cmd.Flags().Changed("directory") {
		v, _ := cmd.Flags().GetString("directory")
		body["directory"] = v
	}
	installationUuidVal, _ := cmd.Flags().GetString("installation-uuid")
	body["installation_uuid"] = installationUuidVal
	if cmd.Flags().Changed("is-enabled") {
		v, _ := cmd.Flags().GetBool("is-enabled")
		body["is_enabled"] = v
	}
	ownerVal, _ := cmd.Flags().GetString("owner")
	body["owner"] = ownerVal
	repositoryVal, _ := cmd.Flags().GetString("repository")
	body["repository"] = repositoryVal
	return body
}
