package scans

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RunWebsiteCmd = &cobra.Command{
	Use:   "run-website",
	Short: "Run website scan",
	Long:  "Start a scan of how ready a public website is for AI agents to discover, read and act on.\\n\nScanning is asynchronous: this operation only starts the scan, and the returned scan is usually `pending`\nwith no results yet. To get the results, poll the `Get website scan` operation with the same domain every\nfew seconds until its status is `completed` or `failed`. The scan itself is limited to 60 seconds,\nplus any time it waits in the queue. A `failed` scan is not retried; run this operation again.\\n\nStarting a scan for a domain that already has a pending scan from the last 5 minutes\nreturns that scan instead of starting a new one.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(runWebsiteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgentReadinessRunWebsiteScanV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	RunWebsiteCmd.Flags().StringP("domain", "", "", "Domain name or full URL of the website to scan. Only the host is used.")
	RunWebsiteCmd.MarkFlagRequired("domain")
}

func runWebsiteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body
}
