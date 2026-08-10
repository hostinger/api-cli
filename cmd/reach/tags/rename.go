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

var RenameCmd = &cobra.Command{
	Use:   "rename <profile-uuid> <tag-uuid>",
	Short: "Rename a tag",
	Long:  "Rename a tag.\n\nThe contacts assigned to the tag are unaffected. Names are unique within a profile, so\nrenaming a tag to a name that is already taken is rejected.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(renameBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachRenameATagV1WithBodyWithResponse(context.TODO(), args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	RenameCmd.Flags().StringP("value", "", "", "New tag name")
	RenameCmd.MarkFlagRequired("value")
}

func renameBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	valueVal, _ := cmd.Flags().GetString("value")
	body["value"] = valueVal
	return body
}
