package websites

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Get website list",
	Long:              "List the Hostinger Horizons websites the user owns.\\n\nUse this tool when the user asks which websites they have, or when you need a website ID\nbefore editing, publishing or cloning a website.\\n\nEach website is returned with its ID, status, domain and the URL to open it\nin Hostinger Horizons interface.\\n\nThe complete list of websites is returned in a single response - it is not paginated.",
	Example:           "  hostinger horizons websites list",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HorizonsGetWebsiteListV1WithResponse(ctx)
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
	utils.RegisterArgs(ListCmd)
}
