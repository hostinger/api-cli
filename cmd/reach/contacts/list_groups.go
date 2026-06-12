package contacts

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "List contact groups",
	Long:  "Get a list of all contact groups.\n\nThis endpoint returns a list of contact groups that can be used to organize contacts.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachListContactGroupsV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
