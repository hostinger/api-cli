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

var AddSubmitterTagsCmd = &cobra.Command{
	Use:   "add-submitter-tags <profile-uuid> <form-uuid>",
	Short: "Add form submitter tags",
	Long:  "Add tags to a form so that every contact who submits it from now on receives them.\n\nA tag the form already has is skipped; an unknown tag is rejected. Contacts who submitted earlier are\nnot tagged.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(addSubmitterTagsBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachAddFormSubmitterTagsV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	AddSubmitterTagsCmd.Flags().StringSliceP("tag-uuids", "", nil, "")
	AddSubmitterTagsCmd.MarkFlagRequired("tag-uuids")
}

func addSubmitterTagsBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	tagUuidsVal, _ := cmd.Flags().GetStringSlice("tag-uuids")
	body["tag_uuids"] = tagUuidsVal
	return body
}
