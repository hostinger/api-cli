package databases

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateRemoteConnectionCmd = &cobra.Command{
	Use:   "create-remote-connection <username> <name>",
	Short: "Create account database remote connection",
	Long:  "Allows a remote host to connect to the specified database.\n\nProvide an IPv4/IPv6 address, or \"%\" to allow any host. The database name must be\nthe full name returned by the list databases endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createRemoteConnectionBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingCreateAccountDatabaseRemoteConnectionV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateRemoteConnectionCmd.Flags().StringP("ip", "", "", "Remote host to allow: an IPv4/IPv6 address, or \"%\" for any host.")
	CreateRemoteConnectionCmd.MarkFlagRequired("ip")
}

func createRemoteConnectionBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	ipVal, _ := cmd.Flags().GetString("ip")
	body["ip"] = ipVal
	return body
}
