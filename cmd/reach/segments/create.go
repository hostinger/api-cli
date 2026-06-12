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

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new contact segment",
	Long:  "Create a new contact segment.\n\nThis endpoint allows creating a new contact segment that can be used to organize contacts.\nThe segment can be configured with specific criteria like email, name, subscription status, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "logic", []string{"AND", "OR"})
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachCreateANewContactSegmentV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("conditions", "", "", " (JSON)")
	CreateCmd.Flags().StringP("logic", "", "", "(one of: AND, OR)")
	CreateCmd.Flags().StringP("name", "", "", "")
	CreateCmd.MarkFlagRequired("conditions")
	CreateCmd.MarkFlagRequired("logic")
	CreateCmd.MarkFlagRequired("name")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	conditionsVal, _ := cmd.Flags().GetString("conditions")
	body["conditions"] = utils.JSONValue(conditionsVal, "conditions")
	logicVal, _ := cmd.Flags().GetString("logic")
	body["logic"] = logicVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	return body
}
