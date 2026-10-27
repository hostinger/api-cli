package domains

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UnlinkFromWebsiteCmd = &cobra.Command{
	Use:     "unlink-from-website <website_uid> <domain>",
	Short:   "Unlink domain from website",
	Long:    "Unlinks a domain from the specified Agency Plan website.\n\nThe website stops serving traffic on this domain immediately.\n\nWebsite files and database are preserved, and any other linked domains remain accessible.\n\nIf this is the only domain on the website, unlinking leaves the website without an accessible domain.",
	Example: "  hostinger agency-hosting domains unlink-from-website zpwlGlp19 mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
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
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Unlink domain from website", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingUnlinkDomainFromWebsiteV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(UnlinkFromWebsiteCmd, utils.StringArg("website_uid"), utils.StringArg("domain"))
	UnlinkFromWebsiteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
