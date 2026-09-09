package websites

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get website list",
	Long:  "List the Hostinger Horizons websites the user owns.\\n\nUse this tool when the user asks which websites they have, or when you need a website ID\nbefore editing, publishing or cloning a website.\\n\nEach website is returned with its ID, status, domain and the URL to open it\nin Hostinger Horizons interface.\\n\nThe complete list of websites is returned in a single response - it is not paginated.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HorizonsGetWebsiteListV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
