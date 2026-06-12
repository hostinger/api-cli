package verifications

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DirectCmd = &cobra.Command{
	Use:   "direct",
	Short: "Get domain verifications",
	Long:  "Retrieve a list of pending and completed domain verifications.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(directBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().V2GetDomainVerificationsDIRECTWithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DirectCmd.Flags().StringSliceP("domains", "", nil, "The list of domains for which to get verification details for.")
	DirectCmd.MarkFlagRequired("domains")
}

func directBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainsVal, _ := cmd.Flags().GetStringSlice("domains")
	body["domains"] = domainsVal
	return body
}
