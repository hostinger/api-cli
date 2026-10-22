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

var CreateCmd = &cobra.Command{
	Use:               "create <mailbox-id>",
	Short:             "Create forwarder",
	Long:              "Create a forwarder from the given mailbox to the destination address.\nThe destination receives a confirmation email and forwarding becomes\nactive only after it is confirmed.",
	Example:           "  hostinger mail forwarders create AC1a2b3c4d5e6f7g --destination jane@example.org",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailCreateForwarderV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd, utils.StringArg("mailbox-id"))
	CreateCmd.Flags().StringP("destination", "", "", "Email address the messages will be forwarded to")
	CreateCmd.Flags().BoolP("is-keep-copy-enabled", "", false, "Whether to keep a copy of forwarded messages in the mailbox. Defaults to false.")
	CreateCmd.MarkFlagRequired("destination")
	CreateCmd.RegisterFlagCompletionFunc("destination", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("is-keep-copy-enabled", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	destinationVal, _ := cmd.Flags().GetString("destination")
	body["destination"] = destinationVal
	if cmd.Flags().Changed("is-keep-copy-enabled") {
		v, _ := cmd.Flags().GetBool("is-keep-copy-enabled")
		body["is_keep_copy_enabled"] = v
	}
	return body, nil
}
