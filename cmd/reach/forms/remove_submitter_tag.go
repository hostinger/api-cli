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

var RemoveSubmitterTagCmd = &cobra.Command{
	Use:     "remove-submitter-tag <profile-uuid> <form-uuid>",
	Short:   "Remove form submitter tag",
	Long:    "Stop a form from applying a tag to its future submitters.\n\nSucceeds also when the form did not have the tag. Contacts who already received the tag keep it, and\nthe tag itself is not deleted. The automatic `form:{name}` tag cannot be removed and is rejected with\na conflict.",
	Example: "  hostinger reach forms remove-submitter-tag 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000 --tag-uuid 550e8400-e29b-41d4-a716-446655440000",
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
		body, err := removeSubmitterTagBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Remove form submitter tag", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachRemoveFormSubmitterTagV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(RemoveSubmitterTagCmd, utils.StringArg("profile-uuid"), utils.StringArg("form-uuid"))
	RemoveSubmitterTagCmd.Flags().StringP("tag-uuid", "", "", "")
	RemoveSubmitterTagCmd.MarkFlagRequired("tag-uuid")
	RemoveSubmitterTagCmd.RegisterFlagCompletionFunc("tag-uuid", cobra.NoFileCompletions)
	RemoveSubmitterTagCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func removeSubmitterTagBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	tagUuidVal, _ := cmd.Flags().GetString("tag-uuid")
	body["tag_uuid"] = tagUuidVal
	return body, nil
}
