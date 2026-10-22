package scans

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:               "run",
	Short:             "Run agent readiness scan",
	Long:              "Start a scan of how ready a public website is for AI agents to discover, read and act on.\\n\nScanning is asynchronous: this operation only starts the scan, and the returned scan is usually `pending`\nwith no results yet. To get the results, poll the `Get agent readiness scan` operation with the same\ndomain every few seconds until its status is `completed` or `failed`. The scan itself is limited to 60 seconds,\nplus any time it waits in the queue. A `failed` scan is not retried; run this operation again.\\n\nStarting a scan for a domain that already has a pending scan from the last 5 minutes\nreturns that scan instead of starting a new one.",
	Example:           "  hostinger agent-readiness scans run --domain example.com",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := runBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgentReadinessRunAgentReadinessScanV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(RunCmd)
	RunCmd.Flags().StringP("domain", "", "", "Domain name or full URL of the website to scan. Only the host is used.")
	RunCmd.MarkFlagRequired("domain")
	RunCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
}

func runBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	return body, nil
}
