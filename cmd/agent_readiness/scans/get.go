package scans

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:     "get <domain>",
	Short:   "Get agent readiness scan",
	Long:    "Get the latest agent readiness scan of a website.\\n\nScans started with the `Run agent readiness scan` operation are asynchronous. While the status is\n`pending`, the result fields are null; wait a few seconds and call this operation again until the status is\n`completed` or `failed`. A `failed` scan is not retried; run the scan again.\\n\nA completed scan has a 0-100 score, a letter grade, a score per category, and the result of every\ncheck with the evidence found and how to fix it. Check names are listed by the `List checks` operation.\\n\nReturns 404 when the domain has never been scanned; use the `Run agent readiness scan` operation first.",
	Example: "  hostinger agent-readiness scans get example.com",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().DomainsGetDomainListV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgentReadinessGetAgentReadinessScanV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(GetCmd, utils.StringArg("domain"))
}
