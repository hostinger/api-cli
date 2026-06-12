package profiles

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Profiles",
	Long:  "This endpoint returns all profiles available to the client, including their basic information.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachListProfilesV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
