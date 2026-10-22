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

var EditCmd = &cobra.Command{
	Use:               "edit <website-id>",
	Short:             "Edit website",
	Long:              "Edit an existing Hostinger Horizons website with a follow-up message.\\n\nUse this tool when the user wants to change, extend or fix a website that already exists.\\n\nThis tool queues the requested changes and returns the website URL and ID.\nThe changes are applied asynchronously.\\n\nAfter invoking this tool, your chat reply must be EXACTLY 1 sentence summarizing\nthat Hostinger Horizons is now applying the requested changes and they will be ready\nin a few minutes, and you should provide the website URL to the user immediately.\nDo not write code.\\n\nIf the tool call fails with an error, you should provide a clear explanation of the error\nand do not generate code yourself in the chat.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --message\n    [\n      {\n        \"text\": \"Add a testimonials section below the menu and make the hero background green\",\n        \"type\": \"text\"\n      }\n    ]",
	Example:           "  hostinger horizons websites edit 123e4567-e89b-12d3-a456-426614174000 --message @message.json",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := editBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HorizonsEditWebsiteV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(EditCmd, utils.StringArg("website-id"))
	EditCmd.Flags().StringP("message", "", "", "(JSON, @file or -; shape in the help above)")
	EditCmd.MarkFlagRequired("message")
	EditCmd.RegisterFlagCompletionFunc("message", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func editBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	messageVal, _ := cmd.Flags().GetString("message")
	messageValJSON, err := utils.JSONValue(messageVal, "message")
	if err != nil {
		return nil, err
	}
	body["message"] = messageValJSON
	return body, nil
}
