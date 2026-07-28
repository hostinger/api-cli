package forwarders

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var UpdateKeepCopySettingCmd = &cobra.Command{
	Use:   "update-keep-copy-setting <forwarder-id>",
	Short: "Update forwarder keep-copy setting",
	Long:  "Enable or disable keeping a copy of forwarded messages in the\nmailbox.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateKeepCopySettingBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MailUpdateForwarderKeepCopySettingV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateKeepCopySettingCmd.Flags().BoolP("is-keep-copy-enabled", "", false, "Whether to keep a copy of forwarded messages in the mailbox")
	UpdateKeepCopySettingCmd.MarkFlagRequired("is-keep-copy-enabled")
}

func updateKeepCopySettingBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	isKeepCopyEnabledVal, _ := cmd.Flags().GetBool("is-keep-copy-enabled")
	body["is_keep_copy_enabled"] = isKeepCopyEnabledVal
	return body
}
