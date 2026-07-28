package aliases

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateAliasCmd = &cobra.Command{
	Use:   "create-alias <mailbox-id>",
	Short: "Create alias",
	Long:  "Create an alias for the given mailbox. The alias address is formed\nfrom the given local part and the domain of the mailbox. Messages\nsent to the alias are delivered to the mailbox.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createAliasBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MailCreateAliasV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateAliasCmd.Flags().StringP("local-part", "", "", "Local part of the alias address (the part before the @). The domain is taken from the mailbox. Case-insensitive and stored lowercase; must start and end with a letter or digit; single dots, underscores and hyphens are allowed in between.")
	CreateAliasCmd.MarkFlagRequired("local-part")
}

func createAliasBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	localPartVal, _ := cmd.Flags().GetString("local-part")
	body["local_part"] = localPartVal
	return body
}
