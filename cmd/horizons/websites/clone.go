package websites

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CloneCmd = &cobra.Command{
	Use:   "clone <website-id>",
	Short: "Clone website",
	Long:  "Clone a Hostinger Horizons website into a new website.\\n\nUse this tool when the user wants a copy of an existing website, for example to try out\nchanges without touching the original.\\n\nThis tool returns the ID and URL of the newly created copy.\nThe original website is left untouched.\\n\nTo edit the copy, use the `Edit website` tool with the returned website ID, or the user can\nopen the provided website URL in Hostinger Horizons interface.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HorizonsCloneWebsiteV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
