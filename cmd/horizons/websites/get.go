package websites

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <website-id>",
	Short: "Get website",
	Long:  "Get the link for the user to open their website in Hostinger Horizons interface.\\n\nUse this tool when the user wants the link to an existing website, or when you need its\nwebsite URL before or after editing it.\\n\nWebsites can be edited with the `Edit website` tool, or by the user in Hostinger Horizons\ninterface in the provided website URL.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HorizonsGetWebsiteV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
