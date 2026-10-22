package nodejs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListVulnerabilitiesCmd = &cobra.Command{
	Use:     "list-vulnerabilities <username> <domain>",
	Short:   "List Node.js vulnerabilities",
	Long:    "Lists known npm package vulnerabilities detected on a Node.js website, enriched with\nadvisory metadata (severity, CVSS score, CVE, advisory URL). Results are sorted from\nthe most severe to the least severe, then by publish date (newest first). Use the\n`severities` query parameter to filter.\n\nVulnerabilities with `is_patchable` set to `true` can be auto-fixed via the\n`Patch Node.js Vulnerabilities` endpoint, which opens a GitHub pull request with\nupdated package versions. Auto-fix is only available for websites deployed from a\nconnected GitHub repository. Vulnerabilities with `is_patching_in_progress` set to\n`true` are already included in an open patch pull request; while any patch pull\nrequest is open, new patch requests for this website are rejected until it is merged\nor closed.\n\nData comes from periodic dependency scans, so it may lag behind the latest deployment.\nAn empty list means the most recent scan found no vulnerabilities; it does not\nguarantee the current deployment is vulnerability-free. Available on Business and\nCloud Hosting plans.",
	Example: "  hostinger hosting nodejs list-vulnerabilities u123456789 mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listVulnerabilitiesParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListNodeJsVulnerabilitiesV1WithResponse(ctx, args[0], args[1], params)
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
	utils.RegisterArgs(ListVulnerabilitiesCmd, utils.StringArg("username"), utils.StringArg("domain"))
	ListVulnerabilitiesCmd.Flags().StringSliceP("severities", "", nil, "Severities to filter by (one of: low, moderate, high, critical, unknown)")
	ListVulnerabilitiesCmd.RegisterFlagCompletionFunc("severities", cobra.FixedCompletions([]cobra.Completion{"low", "moderate", "high", "critical", "unknown"}, cobra.ShellCompDirectiveNoFileComp))
}

func listVulnerabilitiesParams(cmd *cobra.Command) *client.HostingListNodeJsVulnerabilitiesV1Params {
	params := &client.HostingListNodeJsVulnerabilitiesV1Params{}
	if cmd.Flags().Changed("severities") {
		v, _ := cmd.Flags().GetStringSlice("severities")
		es := make([]client.HostingListNodeJsVulnerabilitiesV1ParamsSeverities, len(v))
		for i, s := range v {
			es[i] = client.HostingListNodeJsVulnerabilitiesV1ParamsSeverities(s)
		}
		params.Severities = &es
	}
	return params
}
