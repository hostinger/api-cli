package websites

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var EditCmd = &cobra.Command{
	Use:   "edit <website-id>",
	Short: "Edit website",
	Long:  "Edit an existing Hostinger Horizons website with a follow-up message.\\n\nUse this tool when the user wants to change, extend or fix a website that already exists.\\n\nThis tool queues the requested changes and returns the website URL and ID.\nThe changes are applied asynchronously.\\n\nAfter invoking this tool, your chat reply must be EXACTLY 1 sentence summarizing\nthat Hostinger Horizons is now applying the requested changes and they will be ready\nin a few minutes, and you should provide the website URL to the user immediately.\nDo not write code.\\n\nIf the tool call fails with an error, you should provide a clear explanation of the error\nand do not generate code yourself in the chat.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(editBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HorizonsEditWebsiteV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	EditCmd.Flags().StringP("message", "", "", " (JSON)")
	EditCmd.MarkFlagRequired("message")
}

func editBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	messageVal, _ := cmd.Flags().GetString("message")
	body["message"] = utils.JSONValue(messageVal, "message")
	return body
}
