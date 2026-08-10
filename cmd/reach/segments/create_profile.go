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

var CreateProfileCmd = &cobra.Command{
	Use:   "create-profile <profile-uuid>",
	Short: "Create a profile segment",
	Long:  "Create a segment in a profile.\n\nA segment is a saved set of conditions rather than a fixed list, so its membership changes\nas contacts change. Creating one does not modify any contact.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "logic", []string{"AND", "OR"})
		payload, err := json.Marshal(createProfileBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachCreateAProfileSegmentV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateProfileCmd.Flags().StringP("conditions", "", "", "Conditions a contact must satisfy to fall into the segment (JSON)")
	CreateProfileCmd.Flags().StringP("logic", "", "", "How to combine multiple conditions (one of: AND, OR)")
	CreateProfileCmd.Flags().StringP("name", "", "", "")
	CreateProfileCmd.MarkFlagRequired("conditions")
	CreateProfileCmd.MarkFlagRequired("logic")
	CreateProfileCmd.MarkFlagRequired("name")
}

func createProfileBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	conditionsVal, _ := cmd.Flags().GetString("conditions")
	body["conditions"] = utils.JSONValue(conditionsVal, "conditions")
	logicVal, _ := cmd.Flags().GetString("logic")
	body["logic"] = logicVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body
}
