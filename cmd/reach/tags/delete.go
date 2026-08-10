package tags

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <profile-uuid> <tag-uuid>",
	Short: "Delete a tag",
	Long:  "Delete a tag and remove it from every contact carrying it.\n\nThe contacts themselves are not deleted. This is idempotent: deleting a tag that does not\nexist in the profile still succeeds.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachDeleteATagV1WithResponse(context.TODO(), args[0], uuid.MustParse(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
