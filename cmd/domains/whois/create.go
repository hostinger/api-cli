package whois

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:               "create",
	Short:             "Create WHOIS profile",
	Long:              "Create WHOIS contact profile.\n\nUse this endpoint to add new contact information for domain registration.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --tld-details\n    {}\n\n  --whois-details\n    {}",
	Example:           "  hostinger domains whois create --country NL --entity-type individual --tld com --whois-details '{}'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "entity-type", []string{"individual", "organization"}); err != nil {
			return err
		}
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsCreateWHOISProfileV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd)
	CreateCmd.Flags().StringP("country", "", "", "ISO 3166 2-letter country code")
	CreateCmd.Flags().StringP("entity-type", "", "", "Legal entity type (one of: individual, organization)")
	CreateCmd.Flags().StringP("tld", "", "", "TLD of the domain (without leading dot)")
	CreateCmd.Flags().StringP("tld-details", "", "", "TLD details (JSON, @file or -; shape in the help above)")
	CreateCmd.Flags().StringP("whois-details", "", "", "WHOIS details (JSON, @file or -; shape in the help above)")
	CreateCmd.MarkFlagRequired("country")
	CreateCmd.MarkFlagRequired("entity-type")
	CreateCmd.MarkFlagRequired("tld")
	CreateCmd.MarkFlagRequired("whois-details")
	CreateCmd.RegisterFlagCompletionFunc("country", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("entity-type", cobra.FixedCompletions([]cobra.Completion{"individual", "organization"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("tld", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("tld-details", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateCmd.RegisterFlagCompletionFunc("whois-details", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	countryVal, _ := cmd.Flags().GetString("country")
	body["country"] = countryVal
	entityTypeVal, _ := cmd.Flags().GetString("entity-type")
	body["entity_type"] = entityTypeVal
	tldVal, _ := cmd.Flags().GetString("tld")
	body["tld"] = tldVal
	if cmd.Flags().Changed("tld-details") {
		tldDetailsVal, _ := cmd.Flags().GetString("tld-details")
		tldDetailsValJSON, err := utils.JSONValue(tldDetailsVal, "tld-details")
		if err != nil {
			return nil, err
		}
		body["tld_details"] = tldDetailsValJSON
	}
	whoisDetailsVal, _ := cmd.Flags().GetString("whois-details")
	whoisDetailsValJSON, err := utils.JSONValue(whoisDetailsVal, "whois-details")
	if err != nil {
		return nil, err
	}
	body["whois_details"] = whoisDetailsValJSON
	return body, nil
}
