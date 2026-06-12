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
	Long:  "Get a link for the user to edit their website in Hostinger Horizons interface.\\n\nUse this tool when user wants to modify, edit or add new features to an existing website.\\n\nWebsites can only be edited in Hostinger Horizons interface in the provided website URL.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HorizonsGetWebsiteV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
