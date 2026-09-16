package git

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListInstallationsCmd = &cobra.Command{
	Use:   "list-installations",
	Short: "List Git installations",
	Long:  "Lists the Git provider accounts the customer has connected. Only installations with status\n`active` are returned unless the `status` filter says otherwise.\n\nAn empty list means the customer has no active installation. Check `status=suspended` and\n`status=pending` as well. If there is none at all, a Git provider (GitHub or GitLab) has to be\nconnected once in hPanel (Websites, Manage, Advanced, Git; or Add Website, Node.js Web App,\nImport Git Repository); this endpoint then lists the new installation.\n\nUse `uuid` as the path parameter of `List Git installation repositories`, and as\n`installation_uuid` in `Start Node.js build` with `source_type` `git` and in\n`Update Git auto-deployment settings`.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "provider", []string{"github", "gitlab", "bitbucket"})
		utils.EnumCheck(cmd, "status", []string{"pending", "active", "suspended"})
		r, err := api.Request().HostingListGitInstallationsV1WithResponse(context.TODO(), listInstallationsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListInstallationsCmd.Flags().StringP("provider", "", "", "Filter by Git provider (one of: github, gitlab, bitbucket)")
	ListInstallationsCmd.Flags().StringP("status", "", "active", "Filter by installation status (one of: pending, active, suspended)")
}

func listInstallationsParams(cmd *cobra.Command) *client.HostingListGitInstallationsV1Params {
	params := &client.HostingListGitInstallationsV1Params{}
	if cmd.Flags().Changed("provider") {
		v, _ := cmd.Flags().GetString("provider")
		e := client.HostingListGitInstallationsV1ParamsProvider(v)
		params.Provider = &e
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		e := client.HostingListGitInstallationsV1ParamsStatus(v)
		params.Status = &e
	}
	return params
}
