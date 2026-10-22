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

var UpdateCmd = &cobra.Command{
	Use:     "update <profile-uuid> <form-uuid>",
	Short:   "Update form",
	Long:    "Rename a form and/or change its status.\n\nSetting a `draft` or `paused` form to `active` takes it live: its latest saved design is published\nand visitors can submit it. `paused` and `draft` stop the form from accepting submissions. A form is\nnot reachable by visitors until it is taken live. Design changes saved while the form is already\n`active` stay unpublished until the form is paused and set to `active` again. Renaming does not rename\nthe automatic `form:{name}` tag.",
	Example: "  hostinger reach forms update 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000",
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
		if err := utils.EnumCheck(cmd, "status", []string{"active", "paused", "draft"}); err != nil {
			return err
		}
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachUpdateFormV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("profile-uuid"), utils.StringArg("form-uuid"))
	UpdateCmd.Flags().StringP("name", "", "", "Form name. Letters, digits, spaces and - . ' _ only.")
	UpdateCmd.Flags().StringP("status", "", "", "`active` takes the form live, `paused` and `draft` stop it accepting submissions. (one of: active, paused, draft)")
	UpdateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"active", "paused", "draft"}, cobra.ShellCompDirectiveNoFileComp))
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		body["status"] = v
	}
	return body, nil
}
