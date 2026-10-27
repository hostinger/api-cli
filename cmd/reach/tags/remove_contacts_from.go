package tags

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

var RemoveContactsFromCmd = &cobra.Command{
	Use:               "remove-contacts-from <profile-uuid> <tag-uuid>",
	Short:             "Remove contacts from a tag",
	Long:              "Remove a tag from many contacts at once.\n\nPass `contact_uuids` to target specific contacts, or `all_contacts` to target every contact\nin the profile. The work is queued, so a success response means it was accepted rather than\nfinished. The tag itself and the contacts are not deleted.",
	Example:           "  hostinger reach tags remove-contacts-from 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := removeContactsFromBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Remove contacts from a tag", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachRemoveContactsFromATagV1WithBodyWithResponse(ctx, args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(RemoveContactsFromCmd, utils.StringArg("profile-uuid"), utils.UUIDArg("tag-uuid"))
	RemoveContactsFromCmd.Flags().BoolP("all-contacts", "", false, "Apply to every contact in the profile")
	RemoveContactsFromCmd.Flags().StringSliceP("contact-uuids", "", nil, "Contacts to apply the change to. Required unless all_contacts is true.")
	RemoveContactsFromCmd.RegisterFlagCompletionFunc("all-contacts", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	RemoveContactsFromCmd.RegisterFlagCompletionFunc("contact-uuids", cobra.NoFileCompletions)
	RemoveContactsFromCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func removeContactsFromBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("all-contacts") {
		v, _ := cmd.Flags().GetBool("all-contacts")
		body["all_contacts"] = v
	}
	if cmd.Flags().Changed("contact-uuids") {
		v, _ := cmd.Flags().GetStringSlice("contact-uuids")
		body["contact_uuids"] = v
	}
	return body, nil
}
