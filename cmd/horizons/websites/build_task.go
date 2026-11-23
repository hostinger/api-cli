package websites

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var BuildTaskCmd = &cobra.Command{
	Use:               "build-task <website-id> <task-id>",
	Short:             "Get website build task",
	Long:              "Get the state and outcome of a website build started by the `Create website` or `Edit website` tool.\\n\nUse this tool with the `task.id` those tools return to find out whether the build completed or failed.\nPoll it every 30 seconds or so while `state` is `TASK_STATE_SUBMITTED` or `TASK_STATE_WORKING`.\\n\nWhen `state` is `TASK_STATE_FAILED` or `TASK_STATE_REJECTED`, the website does not contain the requested\nchanges: tell the user, using `message` and `reason_code`, and do not publish it.\nWhen `state` is `TASK_STATE_CANCELED`, the user stopped the build and changes made before that may be kept:\npublish only if the user asks to.\n`reason_code` `insufficient_credits` means the user must get more credits before building again.\\n\nWhen `state` is `TASK_STATE_COMPLETED`, the website can be published. A `reason_code` of\n`iteration_limit` on a completed build means it stopped early and kept the work done so far.",
	Example:           "  hostinger horizons websites build-task 123e4567-e89b-12d3-a456-426614174000 223e4567-e89b-12d3-a456-426614174000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HorizonsGetWebsiteBuildTaskV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(BuildTaskCmd, utils.StringArg("website-id"), utils.StringArg("task-id"))
}
