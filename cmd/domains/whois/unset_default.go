package whois

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UnsetDefaultCmd = &cobra.Command{
	Use:   "unset-default <whois-id>",
	Short: "Unset default WHOIS profile",
	Long:  "Unset WHOIS contact profile as default.\n\nThe profile itself is kept, it is only no longer pre-selected for its TLD.\n\nUse this endpoint to stop reusing contact information for new registrations.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsUnsetDefaultWHOISProfileV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
