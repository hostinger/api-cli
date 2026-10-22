package website_setups

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:               "status <order_id> <setup_uuid>",
	Short:             "Get website setup status",
	Long:              "Returns the current status of an Agency Plan website setup started via the setups\nendpoint.\n\nPoll this endpoint using the `setup_uuid` returned from the provisioning request until\n`status` becomes `completed`, at which point `website_uid` identifies the new website.",
	Example:           "  hostinger agency-hosting website-setups status 123456 0193b6d4-fabb-70e0-8ea4-cfe060a45898",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingGetWebsiteSetupStatusV1WithResponse(ctx, utils.StringToInt(args[0]), uuid.MustParse(args[1]))
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
	utils.RegisterArgs(StatusCmd, utils.IntArg("order_id"), utils.UUIDArg("setup_uuid"))
}
