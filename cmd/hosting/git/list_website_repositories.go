package git

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListWebsiteRepositoriesCmd = &cobra.Command{
	Use:   "list-website-repositories <username> <domain>",
	Short: "List website Git repositories",
	Long:  "Lists the Git repositories linked to directories of the website, with\n`Deploy website Git repository` or in the Git section of hPanel: clone URL, branch and directory of\neach one. A repository whose clone failed stays listed; deploying it again retries the clone. GitHub\nand GitLab auto-deployments are not listed here; see `Get Git auto-deployment settings`.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListWebsiteGitRepositoriesV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
