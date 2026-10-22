package forms

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateTemplateCmd = &cobra.Command{
	Use:     "update-template <profile-uuid> <form-uuid>",
	Short:   "Update form template",
	Long:    "Save `content` as the latest design of a form and return the stored version.\n\nThe HTML is sanitized and, unless the plan allows removing it, the Reach signature is added. Visitors\nsubmit the form with an `email` field (required) and optional `name`, `surname` and `phone`. After\nsanitizing, the design must contain a form container (`data-reach-form`), an email input\n(`data-reach-field=\"email\"`) and a submit button (`data-reach-submit`); otherwise the request fails\nwith `422` and `errors.content` lists `missing_form_container`, `missing_email_field` and/or\n`missing_submit_button`.\n\nA `draft` or `paused` form stays unreachable until its status is set to `active`, which publishes this\ndesign. On a form that is already `active` the saved design is not public until the form is paused and\nset to `active` again. Saving a design does not consume AI credits.",
	Example: "  hostinger reach forms update-template 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000 --content '<form data-reach-form><input type=\"email\" name=\"email\" data-reach-field=\"email\" /><button type=\"submit\" data-reach-submit>Subscribe</button></form>'",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "uuid",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().ReachListFormsV1WithResponse(ctx, args[0], nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := updateTemplateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachUpdateFormTemplateV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateTemplateCmd, utils.StringArg("profile-uuid"), utils.StringArg("form-uuid"))
	UpdateTemplateCmd.Flags().StringP("content", "", "", "HTML of the form design. After sanitizing it must still contain a form container (`data-reach-form`), an email input (`data-reach-field=\"email\"`) and a submit button (`data-reach-submit`).")
	UpdateTemplateCmd.MarkFlagRequired("content")
	UpdateTemplateCmd.RegisterFlagCompletionFunc("content", cobra.NoFileCompletions)
}

func updateTemplateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	contentVal, _ := cmd.Flags().GetString("content")
	body["content"] = contentVal
	return body, nil
}
