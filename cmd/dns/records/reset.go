package records

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ResetCmd = &cobra.Command{
	Use:     "reset [domain]",
	Short:   "Reset DNS records",
	Long:    "Reset DNS zone to the default records.\n\nUse this endpoint to restore domain DNS to original configuration.\n\nThe domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger dns records reset mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "domain",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().DomainsGetDomainListV1WithResponse(ctx)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// A destructive command acts without a question only on a target that was typed.
		typed := len(args) > 0
		domain, args, err := site.ResolveDomain(cmd, args, 0, "dns")
		if err != nil {
			return api.CommandError(err)
		}
		body, err := resetBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.ConfirmScoped(cmd, "Reset DNS records", append([]string{domain}, args...), typed); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DNSResetDNSRecordsV1WithBodyWithResponse(ctx, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(ResetCmd, 1, utils.StringArg("domain"))
	ResetCmd.Flags().BoolP("reset-email-records", "", true, "Determines if email records should be reset")
	ResetCmd.Flags().BoolP("sync", "", true, "Determines if operation should be run synchronously")
	ResetCmd.Flags().StringSliceP("whitelisted-record-types", "", nil, "Specifies which record types to not reset")
	ResetCmd.RegisterFlagCompletionFunc("reset-email-records", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	ResetCmd.RegisterFlagCompletionFunc("sync", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	ResetCmd.RegisterFlagCompletionFunc("whitelisted-record-types", cobra.NoFileCompletions)
	ResetCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func resetBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("reset-email-records") {
		v, _ := cmd.Flags().GetBool("reset-email-records")
		body["reset_email_records"] = v
	}
	if cmd.Flags().Changed("sync") {
		v, _ := cmd.Flags().GetBool("sync")
		body["sync"] = v
	}
	if cmd.Flags().Changed("whitelisted-record-types") {
		v, _ := cmd.Flags().GetStringSlice("whitelisted-record-types")
		body["whitelisted_record_types"] = v
	}
	return body, nil
}
