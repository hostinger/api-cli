package databases

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateRemoteConnectionCmd = &cobra.Command{
	Use:     "create-remote-connection <username> <name>",
	Short:   "Create database remote connection",
	Long:    "Allows a remote host to connect to the specified database.\n\nProvide an IPv4/IPv6 address, or \"%\" to allow any host. The database name must be\nthe full name returned by the list databases endpoint. Database creation is synchronous,\nso a 404 here means no database with that name exists under the username, not that it\nis still being created.",
	Example: "  hostinger hosting databases create-remote-connection u123456789 u123456789_test_db --ip 192.0.2.10",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
		nil,
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createRemoteConnectionBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingCreateDatabaseRemoteConnectionV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateRemoteConnectionCmd, utils.StringArg("username"), utils.StringArg("name"))
	CreateRemoteConnectionCmd.Flags().StringP("ip", "", "", "Remote host to allow: an IPv4/IPv6 address, or \"%\" for any host.")
	CreateRemoteConnectionCmd.MarkFlagRequired("ip")
	CreateRemoteConnectionCmd.RegisterFlagCompletionFunc("ip", cobra.NoFileCompletions)
}

func createRemoteConnectionBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	ipVal, _ := cmd.Flags().GetString("ip")
	body["ip"] = ipVal
	return body, nil
}
