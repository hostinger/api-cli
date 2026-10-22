package websites

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:               "get <website-id>",
	Short:             "Get website",
	Long:              "Get the link for the user to open their website in Hostinger Horizons interface.\\n\nUse this tool when the user wants the link to an existing website, or when you need its\nwebsite URL before or after editing it.\\n\n`is_in_progress` is true while changes are being generated or the website is being published;\nwait until it is false before publishing. `published_at` is when the website was last published,\n`is_template` is whether its published pages show the \"Use template\" banner, and\n`has_ecommerce_store` is whether it has an online store.\\n\nWebsites can be edited with the `Edit website` tool, or by the user in Hostinger Horizons\ninterface in the provided website URL.",
	Example:           "  hostinger horizons websites get 123e4567-e89b-12d3-a456-426614174000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HorizonsGetWebsiteV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(GetCmd, utils.StringArg("website-id"))
}
