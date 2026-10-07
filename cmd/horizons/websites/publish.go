package websites

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var PublishCmd = &cobra.Command{
	Use:   "publish <website-id>",
	Short: "Publish website",
	Long:  "Publish a Hostinger Horizons website so its latest changes go live.\\n\nUse this tool when the user asks to publish, deploy or make their website live.\\n\nThis tool starts the publish process and returns the URL the website will be live on.\nPublishing happens asynchronously and takes a few minutes.\\n\nSet `is_template` only when the user explicitly asks to share the website as a template:\ntrue adds a \"Use template\" banner to its published pages that copies the website into the\nvisitor's own account, and false removes it. Leave it out to keep the current setting.\\n\nAfter invoking this tool, your chat reply must be EXACTLY 1 sentence summarizing\nthat the website is being published and you should provide the published URL to the user immediately.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(publishBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HorizonsPublishWebsiteV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	PublishCmd.Flags().BoolP("is-template", "", false, "Set to true to publish the website as a template: its published pages show a \"Use template\" banner\nthat copies the website into the visitor's own account. Set to false to remove the banner.\nLeave it out to keep the current setting.")
}

func publishBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("is-template") {
		v, _ := cmd.Flags().GetBool("is-template")
		body["is_template"] = v
	}
	return body
}
