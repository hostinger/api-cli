package websites

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var PublishCmd = &cobra.Command{
	Use:   "publish <website-id>",
	Short: "Publish website",
	Long:  "Publish a Hostinger Horizons website so its latest changes go live.\\n\nUse this tool when the user asks to publish, deploy or make their website live.\\n\nThis tool starts the publish process and returns the URL the website will be live on.\nPublishing happens asynchronously and takes a few minutes.\\n\nAfter invoking this tool, your chat reply must be EXACTLY 1 sentence summarizing\nthat the website is being published and you should provide the published URL to the user immediately.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HorizonsPublishWebsiteV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
