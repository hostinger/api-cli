package websites

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PublishCmd = &cobra.Command{
	Use:               "publish <website-id>",
	Short:             "Publish website",
	Long:              "Publish a Hostinger Horizons website so its latest changes go live.\\n\nUse this tool when the user asks to publish, deploy or make their website live.\\n\nThis tool starts the publish process and returns the URL the website will be live on.\nPublishing happens asynchronously and takes a few minutes.\\n\nSet `is_template` only when the user explicitly asks to share the website as a template:\ntrue adds a \"Use template\" banner to its published pages that copies the website into the\nvisitor's own account, and false removes it. Leave it out to keep the current setting.\\n\nAfter invoking this tool, your chat reply must be EXACTLY 1 sentence summarizing\nthat the website is being published and you should provide the published URL to the user immediately.",
	Example:           "  hostinger horizons websites publish 123e4567-e89b-12d3-a456-426614174000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := publishBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HorizonsPublishWebsiteV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(PublishCmd, utils.StringArg("website-id"))
	PublishCmd.Flags().BoolP("is-template", "", false, "Set to true to publish the website as a template: its published pages show a \"Use template\" banner that copies the website into the visitor's own account. Set to false to remove the banner. Leave it out to keep the current setting.")
	PublishCmd.RegisterFlagCompletionFunc("is-template", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func publishBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("is-template") {
		v, _ := cmd.Flags().GetBool("is-template")
		body["is_template"] = v
	}
	return body, nil
}
