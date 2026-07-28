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

var CreateCmd = &cobra.Command{
	Use:   "create <mailbox-id>",
	Short: "Create forwarder",
	Long:  "Create a forwarder from the given mailbox to the destination address.\nThe destination receives a confirmation email and forwarding becomes\nactive only after it is confirmed.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MailCreateForwarderV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("destination", "", "", "Email address the messages will be forwarded to")
	CreateCmd.Flags().BoolP("is-keep-copy-enabled", "", false, "Whether to keep a copy of forwarded messages in the mailbox. Defaults to false.")
	CreateCmd.MarkFlagRequired("destination")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	destinationVal, _ := cmd.Flags().GetString("destination")
	body["destination"] = destinationVal
	if cmd.Flags().Changed("is-keep-copy-enabled") {
		v, _ := cmd.Flags().GetBool("is-keep-copy-enabled")
		body["is_keep_copy_enabled"] = v
	}
	return body
}
