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

var AddSubmitterTagsCmd = &cobra.Command{
	Use:     "add-submitter-tags <profile-uuid> <form-uuid>",
	Short:   "Add form submitter tags",
	Long:    "Add tags to a form so that every contact who submits it from now on receives them.\n\nA tag the form already has is skipped; an unknown tag is rejected. Contacts who submitted earlier are\nnot tagged.",
	Example: "  hostinger reach forms add-submitter-tags 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000 --tag-uuids 550e8400-e29b-41d4-a716-446655440000",
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
		body, err := addSubmitterTagsBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachAddFormSubmitterTagsV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(AddSubmitterTagsCmd, utils.StringArg("profile-uuid"), utils.StringArg("form-uuid"))
	AddSubmitterTagsCmd.Flags().StringSliceP("tag-uuids", "", nil, "")
	AddSubmitterTagsCmd.MarkFlagRequired("tag-uuids")
	AddSubmitterTagsCmd.RegisterFlagCompletionFunc("tag-uuids", cobra.NoFileCompletions)
}

func addSubmitterTagsBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	tagUuidsVal, _ := cmd.Flags().GetStringSlice("tag-uuids")
	body["tag_uuids"] = tagUuidsVal
	return body, nil
}
