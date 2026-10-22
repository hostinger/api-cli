package move

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var OutgoingListCmd = &cobra.Command{
	Use:               "outgoing-list",
	Short:             "Get outgoing domain move list",
	Long:              "Retrieve all domains you are moving to other Hostinger accounts.\n\nOnly moves which have not completed yet are returned.\n\nUse this endpoint to track moves you have initiated and the accounts they are addressed to.",
	Example:           "  hostinger domains move outgoing-list",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsGetOutgoingDomainMoveListV1WithResponse(ctx)
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
	utils.RegisterArgs(OutgoingListCmd)
}
