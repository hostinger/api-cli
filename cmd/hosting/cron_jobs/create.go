package cron_jobs

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:     "create [username]",
	Short:   "Create account cron job",
	Long:    "Creates a cron job for the specified account from a schedule expression and a command.\n\nReturns the created cron job, including its uid, which is required to delete the cron job or fetch its output.\n\nThe hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting cron-jobs create u123456789 --command 'php /home/u123456789/cleanup.php' --time '0 2 * * *'",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "username",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		username, args, err := site.ResolveUsername(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
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
			r, err := api.Request().HostingCreateAccountCronJobV1WithBodyWithResponse(ctx, username, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(CreateCmd, 1, utils.StringArg("username"))
	CreateCmd.Flags().StringP("command", "", "", "Command to execute on the schedule.")
	CreateCmd.Flags().StringP("time", "", "", "Cron schedule expression (for example \"0 2 * * *\" runs daily at 02:00).")
	CreateCmd.MarkFlagRequired("command")
	CreateCmd.MarkFlagRequired("time")
	CreateCmd.RegisterFlagCompletionFunc("command", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("time", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	commandVal, _ := cmd.Flags().GetString("command")
	body["command"] = commandVal
	timeVal, _ := cmd.Flags().GetString("time")
	body["time"] = timeVal
	return body, nil
}
