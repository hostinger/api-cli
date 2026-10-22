package forwarders

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateKeepCopySettingCmd = &cobra.Command{
	Use:               "update-keep-copy-setting <forwarder-id>",
	Short:             "Update forwarder keep-copy setting",
	Long:              "Enable or disable keeping a copy of forwarded messages in the\nmailbox.",
	Example:           "  hostinger mail forwarders update-keep-copy-setting FW1a2b3c4d5e6f7g --is-keep-copy-enabled=true",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := updateKeepCopySettingBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailUpdateForwarderKeepCopySettingV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateKeepCopySettingCmd, utils.StringArg("forwarder-id"))
	UpdateKeepCopySettingCmd.Flags().BoolP("is-keep-copy-enabled", "", false, "Whether to keep a copy of forwarded messages in the mailbox")
	UpdateKeepCopySettingCmd.MarkFlagRequired("is-keep-copy-enabled")
	UpdateKeepCopySettingCmd.RegisterFlagCompletionFunc("is-keep-copy-enabled", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func updateKeepCopySettingBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	isKeepCopyEnabledVal, _ := cmd.Flags().GetBool("is-keep-copy-enabled")
	body["is_keep_copy_enabled"] = isKeepCopyEnabledVal
	return body, nil
}
