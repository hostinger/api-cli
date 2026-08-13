package websites

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <domain>",
	Short: "Delete website",
	Long:  "This endpoint permanently removes a website and all of its data. This action\ncannot be undone. Before calling it, make sure the user understands the\nconsequences and explicitly confirms that they want to proceed.\n\nAll website files, databases and related configuration will be removed.\nThe hosting plan itself is kept, so a new website can be created on it afterwards.\n\nSupported websites: main and addon domain websites on web hosting plans, and\nWebsite Builder websites. Parked domains and subdomains cannot be deleted with\nthis endpoint. The domain must be the exact website domain, not a preview\ndomain or an alias.\n\nReturns 404 when the domain does not exist or does not belong to the\nauthenticated client.\n\nWebsite removal is processed asynchronously and can take a few minutes to\ncomplete. The response returns before the removal finishes.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingDeleteWebsiteV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
