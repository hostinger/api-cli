package transfer

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get transfer list",
	Long:  "Retrieve all domain transfers in your portfolio.\n\nUse this endpoint to monitor incoming and outgoing registrar transfers across your domains.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsGetTransferListV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
