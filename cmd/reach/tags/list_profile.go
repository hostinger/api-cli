package tags

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListProfileCmd = &cobra.Command{
	Use:   "list-profile <profile-uuid>",
	Short: "List profile tags",
	Long:  "Get all tags defined in a profile.\n\nTags are the way contacts are grouped in Reach, and can be used to filter the contact\nlist or to build segments.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachListProfileTagsV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
