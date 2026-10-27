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

var DeleteCmd = &cobra.Command{
	Use:     "delete [domain]",
	Short:   "Delete DNS records",
	Long:    "Delete DNS records for the selected domain.\n\nTo filter which records to delete, add the `name` of the record and `type` to the filter.\nMultiple filters can be provided with single request.\n\nIf you have multiple records with the same name and type, and you want to delete only part of them,\nrefer to the `Update zone records` endpoint.\n\nUse this endpoint to remove specific DNS records from domains.\n\nThe domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --filters\n    [\n      {\n        \"name\": \"@\",\n        \"type\": \"A\"\n      }\n    ]",
	Example: "  hostinger dns records delete mydomain.tld --filters '[{\"name\":\"@\",\"type\":\"A\"}]'",
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
		body, err := deleteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.ConfirmScoped(cmd, "Delete DNS records", append([]string{domain}, args...), typed); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DNSDeleteDNSRecordsV1WithBodyWithResponse(ctx, domain, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(DeleteCmd, 1, utils.StringArg("domain"))
	DeleteCmd.Flags().StringP("filters", "", "", "Filter records for deletion (JSON, @file or -; shape in the help above)")
	DeleteCmd.MarkFlagRequired("filters")
	DeleteCmd.RegisterFlagCompletionFunc("filters", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	DeleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func deleteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	filtersVal, _ := cmd.Flags().GetString("filters")
	filtersValJSON, err := utils.JSONValue(filtersVal, "filters")
	if err != nil {
		return nil, err
	}
	body["filters"] = filtersValJSON
	return body, nil
}
