package segments

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List segments",
	Long:  "Get a list of all contact segments.\n\nThis endpoint returns a list of contact segments that can be used to organize contacts.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachListSegmentsV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
