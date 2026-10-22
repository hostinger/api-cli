package cron_jobs

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateWebsiteCmd = &cobra.Command{
	Use:               "create-website <website_uid>",
	Short:             "Create website cron job",
	Long:              "Creates a cron job for an Agency Plan website from a schedule expression and a command.\n\nReturns the created cron job, including its uuid, which is required to delete the cron job.",
	Example:           "  hostinger agency-hosting cron-jobs create-website zpwlGlp19 --command 'php artisan schedule:run' --time '*/30 * * * *'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createWebsiteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingCreateWebsiteCronJobV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateWebsiteCmd, utils.StringArg("website_uid"))
	CreateWebsiteCmd.Flags().StringP("command", "", "", "Command to run on the schedule. Must not contain pipe (|) or redirection (<, >) characters.")
	CreateWebsiteCmd.Flags().StringP("time", "", "", "Cron schedule expression (standard 5-field crontab syntax).")
	CreateWebsiteCmd.MarkFlagRequired("command")
	CreateWebsiteCmd.MarkFlagRequired("time")
	CreateWebsiteCmd.RegisterFlagCompletionFunc("command", cobra.NoFileCompletions)
	CreateWebsiteCmd.RegisterFlagCompletionFunc("time", cobra.NoFileCompletions)
}

func createWebsiteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	commandVal, _ := cmd.Flags().GetString("command")
	body["command"] = commandVal
	timeVal, _ := cmd.Flags().GetString("time")
	body["time"] = timeVal
	return body, nil
}
