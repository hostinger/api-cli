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

var CreateCmd = &cobra.Command{
	Use:   "create <username>",
	Short: "Create account cron job",
	Long:  "Creates a cron job for the specified account from a schedule expression and a command.\n\nReturns the created cron job, including its uid, which is required to delete the cron job or fetch its output.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingCreateAccountCronJobV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("command", "", "", "Command to execute on the schedule.")
	CreateCmd.Flags().StringP("time", "", "", "Cron schedule expression (for example \"0 2 * * *\" runs daily at 02:00).")
	CreateCmd.MarkFlagRequired("command")
	CreateCmd.MarkFlagRequired("time")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	commandVal, _ := cmd.Flags().GetString("command")
	body["command"] = commandVal
	timeVal, _ := cmd.Flags().GetString("time")
	body["time"] = timeVal
	return body
}
