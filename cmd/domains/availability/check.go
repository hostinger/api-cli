package availability

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CheckCmd = &cobra.Command{
	Use:               "check",
	Short:             "Check domain availability",
	Long:              "Check availability of domain names across multiple TLDs.\n\nMultiple TLDs can be checked at once.\nIf you want alternative domains with response, provide only one TLD and set `with_alternatives` to `true`.\nTLDs should be provided without leading dot (e.g. `com`, `net`, `org`).\n\nEndpoint has rate limit of 90 requests per minute.\n\nUse this endpoint to verify domain availability before purchase.",
	Example:           "  hostinger domains availability check --domain mydomain --tlds com,net,org",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := checkBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsCheckDomainAvailabilityV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CheckCmd)
	CheckCmd.Flags().StringP("domain", "", "", "Domain name (without TLD)")
	CheckCmd.Flags().StringSliceP("tlds", "", nil, "TLDs list")
	CheckCmd.Flags().BoolP("with-alternatives", "", false, "Should response include alternatives")
	CheckCmd.MarkFlagRequired("domain")
	CheckCmd.MarkFlagRequired("tlds")
	CheckCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	CheckCmd.RegisterFlagCompletionFunc("tlds", cobra.NoFileCompletions)
	CheckCmd.RegisterFlagCompletionFunc("with-alternatives", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func checkBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	tldsVal, _ := cmd.Flags().GetStringSlice("tlds")
	body["tlds"] = tldsVal
	if cmd.Flags().Changed("with-alternatives") {
		v, _ := cmd.Flags().GetBool("with-alternatives")
		body["with_alternatives"] = v
	}
	return body, nil
}
