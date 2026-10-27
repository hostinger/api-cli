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

var ValidateCmd = &cobra.Command{
	Use:     "validate [domain]",
	Short:   "Validate DNS records",
	Long:    "Validate DNS records prior to update for the selected domain.\n\nIf the validation is successful, the response will contain `200 Success` code.\nIf there is validation error, the response will fail with `422 Validation error` code.\n\nUse this endpoint to verify DNS record validity before applying changes.\n\nThe domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --zone\n    [\n      {\n        \"name\": \"www\",\n        \"records\": [\n          {\n            \"content\": \"mydomain.tld.\"\n          }\n        ],\n        \"ttl\": 14400,\n        \"type\": \"A\"\n      }\n    ]",
	Example: "  hostinger dns records validate mydomain.tld --zone @zone.json",
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
		domain, args, err := site.ResolveDomain(cmd, args, 0, "dns")
		if err != nil {
			return api.CommandError(err)
		}
		body, err := validateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DNSValidateDNSRecordsV1WithBodyWithResponse(ctx, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(ValidateCmd, 1, utils.StringArg("domain"))
	ValidateCmd.Flags().BoolP("overwrite", "", true, "If `true`, resource records (RRs) matching name and type will be deleted and new RRs will be created, otherwise resource records' ttl's are updated and new records are appended. If no matching RRs are found, they are created.")
	ValidateCmd.Flags().StringP("zone", "", "", "(JSON, @file or -; shape in the help above)")
	ValidateCmd.MarkFlagRequired("zone")
	ValidateCmd.RegisterFlagCompletionFunc("overwrite", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	ValidateCmd.RegisterFlagCompletionFunc("zone", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func validateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("overwrite") {
		v, _ := cmd.Flags().GetBool("overwrite")
		body["overwrite"] = v
	}
	zoneVal, _ := cmd.Flags().GetString("zone")
	zoneValJSON, err := utils.JSONValue(zoneVal, "zone")
	if err != nil {
		return nil, err
	}
	body["zone"] = zoneValJSON
	return body, nil
}
