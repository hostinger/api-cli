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

var DeleteCmd = &cobra.Command{
	Use:     "delete <domain>",
	Short:   "Delete DNS records",
	Long:    "Delete DNS records for the selected domain.\n\nTo filter which records to delete, add the `name` of the record and `type` to the filter.\nMultiple filters can be provided with single request.\n\nIf you have multiple records with the same name and type, and you want to delete only part of them,\nrefer to the `Update zone records` endpoint.\n\nUse this endpoint to remove specific DNS records from domains.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --filters\n    [\n      {\n        \"name\": \"@\",\n        \"type\": \"A\"\n      }\n    ]",
	Example: "  hostinger dns records delete mydomain.tld --filters '[{\"name\":\"@\",\"type\":\"A\"}]'",
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
		body, err := deleteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete DNS records", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DNSDeleteDNSRecordsV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(DeleteCmd, utils.StringArg("domain"))
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
