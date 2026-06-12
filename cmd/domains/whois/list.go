package whois

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get WHOIS profile list",
	Long:  "Retrieve WHOIS contact profiles.\n\nUse this endpoint to view available contact profiles for domain registration.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsGetWHOISProfileListV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("tld", "", "", "Filter by TLD (without leading dot)")
}

func listParams(cmd *cobra.Command) *client.DomainsGetWHOISProfileListV1Params {
	params := &client.DomainsGetWHOISProfileListV1Params{}
	if cmd.Flags().Changed("tld") {
		v, _ := cmd.Flags().GetString("tld")
		params.Tld = &v
	}
	return params
}
