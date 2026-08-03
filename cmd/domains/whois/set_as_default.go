package whois

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SetAsDefaultCmd = &cobra.Command{
	Use:   "set-as-default <whois-id>",
	Short: "Set WHOIS profile as default",
	Long:  "Set WHOIS contact profile as default.\n\nThe default profile is pre-selected for the TLD it belongs to when registering new domains.\n\nUse this endpoint to avoid picking contact information for every registration.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsSetWHOISProfileAsDefaultV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
