package catchalls

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCatchAllCmd = &cobra.Command{
	Use:               "create-catch-all <mailbox-id>",
	Short:             "Create catch-all",
	Long:              "Create a catch-all that routes all messages sent to unknown addresses\nof the domain to the given mailbox. The mailbox address receives a\nconfirmation email and the catch-all becomes active only after it is\nconfirmed. A domain can have only one catch-all.",
	Example:           "  hostinger mail catchalls create-catch-all AC1a2b3c4d5e6f7g",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailCreateCatchAllV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(CreateCatchAllCmd, utils.StringArg("mailbox-id"))
}
