package redirects

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateWebsiteCmd = &cobra.Command{
	Use:   "create-website <username> <domain>",
	Short: "Create website redirect",
	Long:  "Creates a redirect from a URL on the selected website to another URL or IP address.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createWebsiteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingCreateWebsiteRedirectV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateWebsiteCmd.Flags().StringP("from", "", "", "Source URL on the selected website")
	CreateWebsiteCmd.Flags().StringP("to", "", "", "Destination URL or IP address")
	CreateWebsiteCmd.MarkFlagRequired("from")
	CreateWebsiteCmd.MarkFlagRequired("to")
}

func createWebsiteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	fromVal, _ := cmd.Flags().GetString("from")
	body["from"] = fromVal
	toVal, _ := cmd.Flags().GetString("to")
	body["to"] = toVal
	return body
}
