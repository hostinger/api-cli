package ai_tools

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

var SetOptionStatusCmd = &cobra.Command{
	Use:   "set-option-status <username> <software>",
	Short: "Set AI option status",
	Long:  "Enable or disable an AI option for the Hostinger Tools plugin on the specified\nWordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "option", []string{"llmstxt", "web2agent"})
		payload, err := json.Marshal(setOptionStatusBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingSetAIOptionStatusV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SetOptionStatusCmd.Flags().BoolP("enable", "", false, "Enable (true) or disable (false) the AI option.")
	SetOptionStatusCmd.Flags().StringP("option", "", "", "AI option name (one of: llmstxt, web2agent)")
	SetOptionStatusCmd.MarkFlagRequired("enable")
	SetOptionStatusCmd.MarkFlagRequired("option")
}

func setOptionStatusBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	enableVal, _ := cmd.Flags().GetBool("enable")
	body["enable"] = enableVal
	optionVal, _ := cmd.Flags().GetString("option")
	body["option"] = optionVal
	return body
}
