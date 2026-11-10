package forwarding

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
	Short:             "Create domain forwarding",
	Long:              "Create domain forwarding configuration.\n\nCreating a forwarding changes the domain's DNS. If the domain does not use Hostinger nameservers, they are\nswitched to Hostinger's nameservers automatically, so website, email and other services whose DNS is hosted\nelsewhere stop working. The root (@) A and AAAA records in the domain's Hostinger DNS zone are also replaced\nwith records pointing to Hostinger, even when the domain already uses Hostinger nameservers. The domain's\ncurrent nameservers are returned by the domain details endpoint (#tag/domains-portfolio); confirm with the\nuser before creating a forwarding on a domain that uses other nameservers. DNS changes can take some time to\npropagate.\n\nUse this endpoint to set up domain redirects to other URLs.",
	Example:           "  hostinger domains forwarding create --domain mydomain.tld --redirect-type 301 --redirect-url https://forward.to.my.url",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "redirect-type", []string{"301", "302"}); err != nil {
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
			r, err := api.Request().DomainsCreateDomainForwardingV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	CreateCmd.Flags().StringP("domain", "", "", "Domain name")
	CreateCmd.Flags().StringP("redirect-type", "", "", "Redirect type (one of: 301, 302)")
	CreateCmd.Flags().StringP("redirect-url", "", "", "URL to forward domain to")
	CreateCmd.MarkFlagRequired("domain")
	CreateCmd.MarkFlagRequired("redirect-type")
	CreateCmd.MarkFlagRequired("redirect-url")
	CreateCmd.RegisterFlagCompletionFunc("domain", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("redirect-type", cobra.FixedCompletions([]cobra.Completion{"301", "302"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("redirect-url", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	redirectTypeVal, _ := cmd.Flags().GetString("redirect-type")
	body["redirect_type"] = redirectTypeVal
	redirectUrlVal, _ := cmd.Flags().GetString("redirect-url")
	body["redirect_url"] = redirectUrlVal
	return body, nil
}
