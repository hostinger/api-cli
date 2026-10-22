package contact_fields

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:               "update <profile-uuid> <field-uuid>",
	Short:             "Update a contact field",
	Long:              "Rename a custom contact field and, for the choice types, replace its option set.\n\nOptions carrying a uuid are kept and relabelled, options without one are created, and any\nexisting option left out of the list is deleted along with the values contacts hold for\nit. The field type and slug cannot be changed.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --options\n    [\n      {\n        \"label\": \"string\",\n        \"uuid\": \"string\"\n      }\n    ]",
	Example:           "  hostinger reach contact-fields update 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000 --label 'Job title'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachUpdateAContactFieldV1WithBodyWithResponse(ctx, args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("profile-uuid"), utils.UUIDArg("field-uuid"))
	UpdateCmd.Flags().StringP("label", "", "", "")
	UpdateCmd.Flags().StringP("options", "", "", "Replaces the option set when provided. Entries carrying a uuid are kept and relabelled, entries without one are created, and any existing option missing from the list is deleted along with the values contacts hold for it. (JSON, @file or -; shape in the help above)")
	UpdateCmd.MarkFlagRequired("label")
	UpdateCmd.RegisterFlagCompletionFunc("label", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("options", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	labelVal, _ := cmd.Flags().GetString("label")
	body["label"] = labelVal
	if cmd.Flags().Changed("options") {
		optionsVal, _ := cmd.Flags().GetString("options")
		optionsValJSON, err := utils.JSONValue(optionsVal, "options")
		if err != nil {
			return nil, err
		}
		body["options"] = optionsValJSON
	}
	return body, nil
}
