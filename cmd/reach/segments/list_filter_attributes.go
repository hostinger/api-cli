package segments

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListFilterAttributesCmd = &cobra.Command{
	Use:               "list-filter-attributes <profile-uuid>",
	Short:             "List segment filter attributes",
	Long:              "List every attribute a segment condition can filter on, with the operators each attribute\naccepts, the value format they expect and, where the value is constrained, the allowed\nvalues.\n\nThe list is profile specific: it includes the profile's custom contact fields, its tags and\nits 20 most recently published campaigns, so the valid attributes cannot be hardcoded. Read\nit before creating or updating a segment to discover the valid `attribute`, `operator` and\n`value` combinations.",
	Example:           "  hostinger reach segments list-filter-attributes 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachListSegmentFilterAttributesV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(ListFilterAttributesCmd, utils.StringArg("profile-uuid"))
}
