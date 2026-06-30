package databases

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListRemoteConnectionsCmd = &cobra.Command{
	Use:   "list-remote-connections <username>",
	Short: "List account database remote connections",
	Long:  "Returns the remote-access rules for the specified account: the remote hosts\n(IPv4/IPv6 addresses, or \"%\" for any host) allowed to connect to the account databases.\n\nUse the domain filter to only return rules for databases assigned to a specific domain.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListAccountDatabaseRemoteConnectionsV1WithResponse(context.TODO(), args[0], listRemoteConnectionsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListRemoteConnectionsCmd.Flags().StringP("domain", "", "", "Filter remote connections by the domain the database is assigned to.\nRules for databases not assigned to any domain are always included.")
}

func listRemoteConnectionsParams(cmd *cobra.Command) *client.HostingListAccountDatabaseRemoteConnectionsV1Params {
	params := &client.HostingListAccountDatabaseRemoteConnectionsV1Params{}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	return params
}
