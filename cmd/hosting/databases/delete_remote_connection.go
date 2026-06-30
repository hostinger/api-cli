package databases

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteRemoteConnectionCmd = &cobra.Command{
	Use:   "delete-remote-connection <username> <name>",
	Short: "Delete account database remote connection",
	Long:  "Permanently removes a remote-access rule, revoking the given host's remote access to the database.\n\nIdentify the rule with the required ip query parameter (the IPv4/IPv6 address, or \"%\",\nexactly as returned by the list remote connections endpoint). The database name must be\nthe full name returned by the list databases endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingDeleteAccountDatabaseRemoteConnectionV1WithResponse(context.TODO(), args[0], args[1], deleteRemoteConnectionParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DeleteRemoteConnectionCmd.Flags().StringP("ip", "", "", "Remote host to revoke: the IPv4/IPv6 address, or \"%\",\nexactly as returned by the list remote connections endpoint.")
	DeleteRemoteConnectionCmd.MarkFlagRequired("ip")
}

func deleteRemoteConnectionParams(cmd *cobra.Command) *client.HostingDeleteAccountDatabaseRemoteConnectionV1Params {
	params := &client.HostingDeleteAccountDatabaseRemoteConnectionV1Params{}
	ipVal, _ := cmd.Flags().GetString("ip")
	params.Ip = ipVal
	return params
}
