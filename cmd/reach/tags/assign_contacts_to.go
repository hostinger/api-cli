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

var AssignContactsToCmd = &cobra.Command{
	Use:   "assign-contacts-to <profile-uuid> <tag-uuid>",
	Short: "Assign contacts to a tag",
	Long:  "Assign a tag to many contacts at once.\n\nPass `contact_uuids` to target specific contacts, or `all_contacts` to target every contact\nin the profile. The work is queued, so a success response means it was accepted rather than\nfinished. Contacts that already carry the tag are left alone.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(assignContactsToBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachAssignContactsToATagV1WithBodyWithResponse(context.TODO(), args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	AssignContactsToCmd.Flags().BoolP("all-contacts", "", false, "Apply to every contact in the profile")
	AssignContactsToCmd.Flags().StringSliceP("contact-uuids", "", nil, "Contacts to apply the change to. Required unless all_contacts is true.")
}

func assignContactsToBody(cmd *cobra.Command) map[string]any {
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
