package move

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var IncomingListCmd = &cobra.Command{
	Use:               "incoming-list",
	Short:             "Get incoming domain move list",
	Long:              "Retrieve all domains other Hostinger accounts are moving to your account.\n\nMoves of every status are returned, including the ones which already completed.\n\nUse this endpoint to find domains waiting for you to accept them.",
	Example:           "  hostinger domains move incoming-list",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsGetIncomingDomainMoveListV1WithResponse(ctx)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(IncomingListCmd)
}
