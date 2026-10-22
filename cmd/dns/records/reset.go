package records

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ResetCmd = &cobra.Command{
	Use:     "reset <domain>",
	Short:   "Reset DNS records",
	Long:    "Reset DNS zone to the default records.\n\nUse this endpoint to restore domain DNS to original configuration.",
	Example: "  hostinger dns records reset mydomain.tld",
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
		body, err := resetBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Reset DNS records", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DNSResetDNSRecordsV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ResetCmd, utils.StringArg("domain"))
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
