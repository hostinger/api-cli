package cron_jobs

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateWebsiteCmd = &cobra.Command{
	Use:   "create-website <website_uid>",
	Short: "Create website cron job",
	Long:  "Creates a cron job for an Agency Plan website from a schedule expression and a command.\n\nReturns the created cron job, including its uuid, which is required to delete the cron job.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createWebsiteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingCreateWebsiteCronJobV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateWebsiteCmd.Flags().StringP("command", "", "", "Command to run on the schedule. Must not contain pipe (|) or redirection (<, >) characters.")
	CreateWebsiteCmd.Flags().StringP("time", "", "", "Cron schedule expression (standard 5-field crontab syntax).")
	CreateWebsiteCmd.MarkFlagRequired("command")
	CreateWebsiteCmd.MarkFlagRequired("time")
}

func createWebsiteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	commandVal, _ := cmd.Flags().GetString("command")
	body["command"] = commandVal
	timeVal, _ := cmd.Flags().GetString("time")
	body["time"] = timeVal
	return body
}
