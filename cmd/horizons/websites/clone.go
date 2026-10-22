package websites

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CloneCmd = &cobra.Command{
	Use:               "clone <website-id>",
	Short:             "Clone website",
	Long:              "Clone a Hostinger Horizons website into a new website.\\n\nUse this tool when the user wants a copy of an existing website, for example to try out\nchanges without touching the original.\\n\nThis tool returns the ID and URL of the newly created copy.\nThe original website is left untouched.\\n\nTo edit the copy, use the `Edit website` tool with the returned website ID, or the user can\nopen the provided website URL in Hostinger Horizons interface.",
	Example:           "  hostinger horizons websites clone 123e4567-e89b-12d3-a456-426614174000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HorizonsCloneWebsiteV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(CloneCmd, utils.StringArg("website-id"))
}
