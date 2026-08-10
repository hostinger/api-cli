package move

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CancelOutgoingCmd = &cobra.Command{
	Use:   "cancel-outgoing <domain>",
	Short: "Cancel outgoing domain move",
	Long:  "Cancel an outgoing move for a specified domain.\n\nThe move can only be cancelled while the receiving account has not accepted it yet.\nThe domain stays in your account.\n\nUse this endpoint to withdraw a move you no longer want to complete.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsCancelOutgoingDomainMoveV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
