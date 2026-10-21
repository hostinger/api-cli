package forms

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RemoveSubmitterTagCmd = &cobra.Command{
	Use:   "remove-submitter-tag <profile-uuid> <form-uuid>",
	Short: "Remove form submitter tag",
	Long:  "Stop a form from applying a tag to its future submitters.\n\nSucceeds also when the form did not have the tag. Contacts who already received the tag keep it, and\nthe tag itself is not deleted. The automatic `form:{name}` tag cannot be removed and is rejected with\na conflict.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(removeSubmitterTagBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachRemoveFormSubmitterTagV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	RemoveSubmitterTagCmd.Flags().StringP("tag-uuid", "", "", "")
	RemoveSubmitterTagCmd.MarkFlagRequired("tag-uuid")
}

func removeSubmitterTagBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	tagUuidVal, _ := cmd.Flags().GetString("tag-uuid")
	body["tag_uuid"] = tagUuidVal
	return body
}
