package aliases

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteAliasCmd = &cobra.Command{
	Use:   "delete-alias <alias-id>",
	Short: "Delete alias",
	Long:  "Delete an alias. Messages sent to the alias address are no longer\ndelivered to the mailbox.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailDeleteAliasV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
