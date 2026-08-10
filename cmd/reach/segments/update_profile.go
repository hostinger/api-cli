package segments

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateProfileCmd = &cobra.Command{
	Use:   "update-profile <profile-uuid> <segment-uuid>",
	Short: "Update a profile segment",
	Long:  "Rename a segment and/or replace the conditions that define it.\n\n`name` is always required. Omit `conditions` to rename without touching the conditions;\nsupply them and they replace the existing set entirely rather than being merged into it.\nContacts are never modified, but which of them match the segment can change immediately.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "logic", []string{"AND", "OR"})
		payload, err := json.Marshal(updateProfileBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachUpdateAProfileSegmentV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateProfileCmd.Flags().StringP("conditions", "", "", "Replaces the existing conditions entirely. Omit to keep the current ones. (JSON)")
	UpdateProfileCmd.Flags().StringP("logic", "", "", "How to combine multiple conditions. Required when conditions are given. (one of: AND, OR)")
	UpdateProfileCmd.Flags().StringP("name", "", "", "")
	UpdateProfileCmd.MarkFlagRequired("name")
}

func updateProfileBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("conditions") {
		v, _ := cmd.Flags().GetString("conditions")
		body["conditions"] = utils.JSONValue(v, "conditions")
	}
	if cmd.Flags().Changed("logic") {
		v, _ := cmd.Flags().GetString("logic")
		body["logic"] = v
	}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body
}
