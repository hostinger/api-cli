package git

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListInstallationRepositoriesCmd = &cobra.Command{
	Use:   "list-installation-repositories <uuid>",
	Short: "List Git installation repositories",
	Long:  "Lists the repositories the Git installation can access, read live from the provider. Works\nfor github and gitlab installations. Use an active installation: a suspended or pending one\nis still queried and the call fails with whatever the provider answers. The list is cut at\nthe first 500 repositories in the order the provider returns them; when the account has\nmore, name the repository directly instead of searching this list.\n\n`owner`, `name` and `default_branch` identify a repository and a branch to deploy. Returns\n404 when the installation does not belong to the customer. Limited to 10 calls per minute\nper API client (429 above that).",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListGitInstallationRepositoriesV1WithResponse(context.TODO(), uuid.MustParse(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
