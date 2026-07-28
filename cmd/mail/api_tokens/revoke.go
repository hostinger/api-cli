package api_tokens

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RevokeCmd = &cobra.Command{
	Use:   "revoke <token-id>",
	Short: "Revoke API token",
	Long:  "Revoke an API token. The token immediately loses access to the\n[Hostinger Email API](https://api.mail.hostinger.com/). This action\ncannot be undone.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailRevokeAPITokenV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
