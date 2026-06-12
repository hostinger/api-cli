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

var ChangePasswordCmd = &cobra.Command{
	Use:   "change-password <username> <name>",
	Short: "Change database password",
	Long:  "Changes the password for the specified database user.\n\nThe database name must be the full name returned by the list databases endpoint.\nThe password must also be updated in any website configuration that uses this database.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(changePasswordBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingChangeDatabasePasswordV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ChangePasswordCmd.Flags().StringP("password", "", "", "New database user password.")
	ChangePasswordCmd.MarkFlagRequired("password")
}

func changePasswordBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body
}
