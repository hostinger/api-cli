package contacts

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
	Use:     "update <profile-uuid> <contact-uuid>",
	Short:   "Update a contact",
	Long:    "Update a contact's attributes and custom field values.\n\nOnly the properties present in the request body are changed, so a partial body is enough\nto change a single attribute. Sending a property as `null` clears it.\n\nThe response carries the contact's core attributes. Read back its tags, custom field\nvalues, source and note with `GET /api/reach/v1/profiles/{profileUuid}/contacts/{contactUuid}`.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --fields\n    [\n      {\n        \"selected_option_uuids\": [\n          \"string\"\n        ],\n        \"uuid\": \"string\",\n        \"value\": \"string\"\n      }\n    ]",
	Example: "  hostinger reach contacts update 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "uuid",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().ReachListContactsV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "subscription-status", []string{"subscribed", "unsubscribed", "confirmed", "pending"}); err != nil {
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
			r, err := api.Request().ReachUpdateAContactV1WithBodyWithResponse(ctx, args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("profile-uuid"), utils.UUIDArg("contact-uuid"))
	UpdateCmd.Flags().StringP("email", "", "", "")
	UpdateCmd.Flags().StringP("fields", "", "", "Set custom field values. Omit to leave untouched, send an empty array to clear them all. (JSON, @file or -; shape in the help above)")
	UpdateCmd.Flags().StringP("name", "", "", "")
	UpdateCmd.Flags().StringP("note", "", "", "")
	UpdateCmd.Flags().StringP("phone", "", "", "Phone number in E.164 format (leading \"+\" then 7-15 digits)")
	UpdateCmd.Flags().StringP("subscription-status", "", "", "(one of: subscribed, unsubscribed, confirmed, pending)")
	UpdateCmd.Flags().StringP("surname", "", "", "")
	UpdateCmd.RegisterFlagCompletionFunc("email", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("fields", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	UpdateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("note", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("phone", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("subscription-status", cobra.FixedCompletions([]cobra.Completion{"subscribed", "unsubscribed", "confirmed", "pending"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("surname", cobra.NoFileCompletions)
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("email") {
		v, _ := cmd.Flags().GetString("email")
		body["email"] = v
	}
	if cmd.Flags().Changed("fields") {
		fieldsVal, _ := cmd.Flags().GetString("fields")
		fieldsValJSON, err := utils.JSONValue(fieldsVal, "fields")
		if err != nil {
			return nil, err
		}
		body["fields"] = fieldsValJSON
	}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("note") {
		v, _ := cmd.Flags().GetString("note")
		body["note"] = v
	}
	if cmd.Flags().Changed("phone") {
		v, _ := cmd.Flags().GetString("phone")
		body["phone"] = v
	}
	if cmd.Flags().Changed("subscription-status") {
		v, _ := cmd.Flags().GetString("subscription-status")
		body["subscription_status"] = v
	}
	if cmd.Flags().Changed("surname") {
		v, _ := cmd.Flags().GetString("surname")
		body["surname"] = v
	}
	return body, nil
}
