package forms

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:     "get <profile-uuid> <form-uuid>",
	Short:   "Get form details",
	Long:    "Get a single form with the URL of its hosted template and the tags it applies to the contacts\nit captures.\n\nThere is no ready-made embed snippet in the response - either serve the template HTML yourself\nor build your own embed around the form uuid.",
	Example: "  hostinger reach forms get 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "uuid",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().ReachListFormsV1WithResponse(ctx, args[0], nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachGetFormDetailsV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(GetCmd, utils.StringArg("profile-uuid"), utils.StringArg("form-uuid"))
}
