package verifications

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DirectCmd = &cobra.Command{
	Use:               "direct",
	Short:             "Get domain verifications",
	Long:              "Retrieve a list of pending and completed domain verifications.",
	Example:           "  hostinger domains verifications direct --domains value",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := directBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().V2GetDomainVerificationsDIRECTWithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(DirectCmd)
	DirectCmd.Flags().StringSliceP("domains", "", nil, "The list of domains for which to get verification details for.")
	DirectCmd.MarkFlagRequired("domains")
	DirectCmd.RegisterFlagCompletionFunc("domains", cobra.NoFileCompletions)
}

func directBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainsVal, _ := cmd.Flags().GetStringSlice("domains")
	body["domains"] = domainsVal
	return body, nil
}
