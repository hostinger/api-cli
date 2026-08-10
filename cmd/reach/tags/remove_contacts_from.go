package tags

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RemoveContactsFromCmd = &cobra.Command{
	Use:   "remove-contacts-from <profile-uuid> <tag-uuid>",
	Short: "Remove contacts from a tag",
	Long:  "Remove a tag from many contacts at once.\n\nPass `contact_uuids` to target specific contacts, or `all_contacts` to target every contact\nin the profile. The work is queued, so a success response means it was accepted rather than\nfinished. The tag itself and the contacts are not deleted.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(removeContactsFromBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachRemoveContactsFromATagV1WithBodyWithResponse(context.TODO(), args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	RemoveContactsFromCmd.Flags().BoolP("all-contacts", "", false, "Apply to every contact in the profile")
	RemoveContactsFromCmd.Flags().StringSliceP("contact-uuids", "", nil, "Contacts to apply the change to. Required unless all_contacts is true.")
}

func removeContactsFromBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("all-contacts") {
		v, _ := cmd.Flags().GetBool("all-contacts")
		body["all_contacts"] = v
	}
	if cmd.Flags().Changed("contact-uuids") {
		v, _ := cmd.Flags().GetStringSlice("contact-uuids")
		body["contact_uuids"] = v
	}
	return body
}
