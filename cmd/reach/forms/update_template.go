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

var UpdateTemplateCmd = &cobra.Command{
	Use:   "update-template <profile-uuid> <form-uuid>",
	Short: "Update form template",
	Long:  "Save `content` as the latest design of a form and return the stored version.\n\nThe HTML is sanitized and, unless the plan allows removing it, the Reach signature is added. Visitors\nsubmit the form with an `email` field (required) and optional `name`, `surname` and `phone`. After\nsanitizing, the design must contain a form container (`data-reach-form`), an email input\n(`data-reach-field=\"email\"`) and a submit button (`data-reach-submit`); otherwise the request fails\nwith `422` and `errors.content` lists `missing_form_container`, `missing_email_field` and/or\n`missing_submit_button`.\n\nA `draft` or `paused` form stays unreachable until its status is set to `active`, which publishes this\ndesign. On a form that is already `active` the saved design is not public until the form is paused and\nset to `active` again. Saving a design does not consume AI credits.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateTemplateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachUpdateFormTemplateV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateTemplateCmd.Flags().StringP("content", "", "", "HTML of the form design. After sanitizing it must still contain a form container\n(`data-reach-form`), an email input (`data-reach-field=\"email\"`) and a submit button\n(`data-reach-submit`).")
	UpdateTemplateCmd.MarkFlagRequired("content")
}

func updateTemplateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	contentVal, _ := cmd.Flags().GetString("content")
	body["content"] = contentVal
	return body
}
