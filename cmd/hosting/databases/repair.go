package databases

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RepairCmd = &cobra.Command{
	Use:   "repair <username> <name>",
	Short: "Repair database",
	Long:  "Repairs corrupted database tables asynchronously.\n\nUse when database errors, crashes, or corruption are reported.\nThe database name must be the full name returned by the list databases endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingRepairDatabaseV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
