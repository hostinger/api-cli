package whois

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Get WHOIS profile list",
	Long:              "Retrieve WHOIS contact profiles.\n\nUse this endpoint to view available contact profiles for domain registration.",
	Example:           "  hostinger domains whois list",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsGetWHOISProfileListV1WithResponse(ctx, params)
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
	utils.RegisterArgs(ListCmd)
	ListCmd.Flags().StringP("tld", "", "", "Filter by TLD (without leading dot)")
	ListCmd.RegisterFlagCompletionFunc("tld", cobra.NoFileCompletions)
}

func listParams(cmd *cobra.Command) *client.DomainsGetWHOISProfileListV1Params {
	params := &client.DomainsGetWHOISProfileListV1Params{}
	if cmd.Flags().Changed("tld") {
		v, _ := cmd.Flags().GetString("tld")
		params.Tld = &v
	}
	return params
}
