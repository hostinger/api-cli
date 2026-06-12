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

var CreateCmd = &cobra.Command{
	Use:   "create <username>",
	Short: "Create account database",
	Long:  "Creates a database with a database user and password for the specified account.\n\nThe database name and user are automatically prefixed with the account username when needed.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingCreateAccountDatabaseV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("name", "", "", "Database name. If the account username prefix is omitted, it is added automatically.")
	CreateCmd.Flags().StringP("password", "", "", "Database user password.")
	CreateCmd.Flags().StringP("user", "", "", "Database user. If the account username prefix is omitted, it is added automatically.")
	CreateCmd.Flags().StringP("website-domain", "", "", "Website domain assigned to the database.")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("password")
	CreateCmd.MarkFlagRequired("user")
	CreateCmd.MarkFlagRequired("website-domain")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	userVal, _ := cmd.Flags().GetString("user")
	body["user"] = userVal
	websiteDomainVal, _ := cmd.Flags().GetString("website-domain")
	body["website_domain"] = websiteDomainVal
	return body
}
