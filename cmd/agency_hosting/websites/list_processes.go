package websites

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListProcessesCmd = &cobra.Command{
	Use:               "list-processes <website_uid>",
	Short:             "List website processes",
	Long:              "Lists active and recently completed asynchronous processes for an Agency Plan website.\n\nEach process has a unique ID (for tracking), a type, and a status (running, completed, failed).\nPoll this endpoint after initiating async operations (SSL setup, backups, cloning) to track progress.",
	Example:           "  hostinger agency-hosting websites list-processes zpwlGlp19",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingListWebsiteProcessesV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(ListProcessesCmd, utils.StringArg("website_uid"))
}
