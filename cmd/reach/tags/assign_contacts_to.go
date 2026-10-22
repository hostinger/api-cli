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

var AssignContactsToCmd = &cobra.Command{
	Use:               "assign-contacts-to <profile-uuid> <tag-uuid>",
	Short:             "Assign contacts to a tag",
	Long:              "Assign a tag to many contacts at once.\n\nPass `contact_uuids` to target specific contacts, or `all_contacts` to target every contact\nin the profile. The work is queued, so a success response means it was accepted rather than\nfinished. Contacts that already carry the tag are left alone.",
	Example:           "  hostinger reach tags assign-contacts-to 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := assignContactsToBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachAssignContactsToATagV1WithBodyWithResponse(ctx, args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(AssignContactsToCmd, utils.StringArg("profile-uuid"), utils.UUIDArg("tag-uuid"))
	AssignContactsToCmd.Flags().BoolP("all-contacts", "", false, "Apply to every contact in the profile")
	AssignContactsToCmd.Flags().StringSliceP("contact-uuids", "", nil, "Contacts to apply the change to. Required unless all_contacts is true.")
	AssignContactsToCmd.RegisterFlagCompletionFunc("all-contacts", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	AssignContactsToCmd.RegisterFlagCompletionFunc("contact-uuids", cobra.NoFileCompletions)
}

func assignContactsToBody(cmd *cobra.Command) (map[string]any, error) {
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
