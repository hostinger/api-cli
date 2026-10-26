package contacts

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateInBulkCmd = &cobra.Command{
	Use:               "create-in-bulk <profile-uuid>",
	Short:             "Create contacts in bulk",
	Long:              "Create many contacts in a profile in a single call.\n\nThe contacts are imported in the background, so a success response means the import was\naccepted rather than finished. Contacts whose email already exists in the profile keep\ntheir current values; only the ones that are still empty, custom fields included, are\nfilled in. If double opt-in is enabled, new contacts start off pending and are sent a\nconfirmation email.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --contacts\n    [\n      {\n        \"email\": \"john.doe@example.com\",\n        \"fields\": [\n          {\n            \"selected_option_uuids\": [\n              \"string\"\n            ],\n            \"uuid\": \"string\",\n            \"value\": \"string\"\n          }\n        ],\n        \"name\": \"John\",\n        \"phone\": \"+14155552671\",\n        \"surname\": \"Doe\"\n      }\n    ]",
	Example:           "  hostinger reach contacts create-in-bulk 550e8400-e09b-41d4-a716-400055000000 --contacts @contacts.json",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createInBulkBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachCreateContactsInBulkV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateInBulkCmd, utils.StringArg("profile-uuid"))
	CreateInBulkCmd.Flags().StringP("contacts", "", "", "(JSON, @file or -; shape in the help above)")
	CreateInBulkCmd.Flags().StringP("note", "", "", "Note applied to every created contact")
	CreateInBulkCmd.Flags().StringSliceP("tag-uuids", "", nil, "Existing tags to attach to every created contact")
	CreateInBulkCmd.MarkFlagRequired("contacts")
	CreateInBulkCmd.RegisterFlagCompletionFunc("contacts", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateInBulkCmd.RegisterFlagCompletionFunc("note", cobra.NoFileCompletions)
	CreateInBulkCmd.RegisterFlagCompletionFunc("tag-uuids", cobra.NoFileCompletions)
}

func createInBulkBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	contactsVal, _ := cmd.Flags().GetString("contacts")
	contactsValJSON, err := utils.JSONValue(contactsVal, "contacts")
	if err != nil {
		return nil, err
	}
	body["contacts"] = contactsValJSON
	if cmd.Flags().Changed("note") {
		v, _ := cmd.Flags().GetString("note")
		body["note"] = v
	}
	if cmd.Flags().Changed("tag-uuids") {
		v, _ := cmd.Flags().GetStringSlice("tag-uuids")
		body["tag_uuids"] = v
	}
	return body, nil
}
