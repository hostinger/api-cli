package profiles

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ConnectedSendingDomainCmd = &cobra.Command{
	Use:               "connected-sending-domain <profile-uuid>",
	Short:             "Get connected sending domain",
	Long:              "Get the sending domain connected to the profile, its verification status and any suspended\nsender addresses.\n\nCampaigns only go out once a domain is connected and active, so this is the cheapest way to\ncheck that precondition before building one. A profile with no domain connected returns the\nsame shape with every field set to `null`. For the individual MX, SPF, DKIM and DMARC records\nbehind the status, use the DNS status endpoint.",
	Example:           "  hostinger reach profiles connected-sending-domain 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachGetConnectedSendingDomainV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(ConnectedSendingDomainCmd, utils.StringArg("profile-uuid"))
}
