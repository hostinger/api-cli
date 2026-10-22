package nodejs

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PatchVulnerabilitiesCmd = &cobra.Command{
	Use:     "patch-vulnerabilities <username> <domain>",
	Short:   "Patch Node.js vulnerabilities",
	Long:    "Patches the selected Node.js vulnerabilities by updating the affected package versions\nin `package.json` and opening a GitHub pull request in the connected repository. The\ncustomer reviews and merges the pull request; merging triggers the automatic deployment.\n\nAuto-fix is only available for websites deployed from a connected GitHub repository.\nWebsites deployed from an archive have no auto-fix path and return a 404. The Hostinger\nGitHub App needs write access to the repository; without it the request fails with a\n403 explaining the missing permission.\n\nOnly vulnerabilities with `is_patchable` set to `true` can be patched. Non-patchable\nIDs in the selection are skipped; the pull request covers the patchable subset, listed\nin `patched_vulnerability_ids`. Selections without any patchable vulnerability are\nrejected with a 422. Only one patch pull request can be open at a time per website;\nclose or merge it before patching again. Available on Business and Cloud Hosting plans.",
	Example: "  hostinger hosting nodejs patch-vulnerabilities u123456789 mydomain.tld --vulnerability-ids value",
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
		body, err := patchVulnerabilitiesBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingPatchNodeJsVulnerabilitiesV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(PatchVulnerabilitiesCmd, utils.StringArg("username"), utils.StringArg("domain"))
	PatchVulnerabilitiesCmd.Flags().StringSliceP("vulnerability-ids", "", nil, "List of vulnerability IDs to patch, as returned by the list vulnerabilities endpoint.")
	PatchVulnerabilitiesCmd.MarkFlagRequired("vulnerability-ids")
	PatchVulnerabilitiesCmd.RegisterFlagCompletionFunc("vulnerability-ids", cobra.NoFileCompletions)
}

func patchVulnerabilitiesBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	vulnerabilityIdsVal, _ := cmd.Flags().GetStringSlice("vulnerability-ids")
	body["vulnerability_ids"] = vulnerabilityIdsVal
	return body, nil
}
