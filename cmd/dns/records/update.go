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

var UpdateCmd = &cobra.Command{
	Use:     "update <domain>",
	Short:   "Update DNS records",
	Long:    "Update DNS records for the selected domain.\n\nUsing `overwrite = true` will replace existing records with the provided ones.\nOtherwise existing records will be updated and new records will be added.\n\nUse this endpoint to modify domain DNS configuration.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --zone\n    [\n      {\n        \"name\": \"www\",\n        \"records\": [\n          {\n            \"content\": \"mydomain.tld.\"\n          }\n        ],\n        \"ttl\": 14400,\n        \"type\": \"A\"\n      }\n    ]",
	Example: "  hostinger dns records update mydomain.tld --zone @zone.json",
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
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DNSUpdateDNSRecordsV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("domain"))
	UpdateCmd.Flags().BoolP("overwrite", "", true, "If `true`, resource records (RRs) matching name and type will be deleted and new RRs will be created, otherwise resource records' ttl's are updated and new records are appended. If no matching RRs are found, they are created.")
	UpdateCmd.Flags().StringP("zone", "", "", "(JSON, @file or -; shape in the help above)")
	UpdateCmd.MarkFlagRequired("zone")
	UpdateCmd.RegisterFlagCompletionFunc("overwrite", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("zone", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
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
