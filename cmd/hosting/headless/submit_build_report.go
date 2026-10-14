package headless

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SubmitBuildReportCmd = &cobra.Command{
	Use:   "submit-build-report",
	Short: "Submit headless build report",
	Long:  "Submits the build report the headless skill writes at the end of a run (including early\nstops), so Hostinger can learn where builds get stuck. Send one report per call, with the\nreport JSON as the request body, unchanged. Only send a report when the user agreed to it.\n\nThe report must match report.schema.json v1; an invalid report returns 422 and is not\nstored. The report is linked to the authenticated account, so do not add account details\nto it.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "schema-version", []string{"1"})
		payload, err := json.Marshal(submitBuildReportBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingSubmitHeadlessBuildReportV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	SubmitBuildReportCmd.Flags().StringP("consent", "", "", "Upload status of the report, as written by the headless skill (JSON)")
	SubmitBuildReportCmd.Flags().StringP("generated-at", "", "", "When the report was written")
	SubmitBuildReportCmd.Flags().StringP("info", "", "", "Website, tool and outcome of the run (JSON)")
	SubmitBuildReportCmd.Flags().StringP("issues", "", "", "Problems hit during the run (JSON)")
	SubmitBuildReportCmd.Flags().StringP("logging-gaps", "", "", "Where the run log is missing or unsure (JSON)")
	SubmitBuildReportCmd.Flags().StringP("previous-runs", "", "", "Earlier runs for the same website (JSON)")
	SubmitBuildReportCmd.Flags().StringP("run", "", "", "The run the report is for (JSON)")
	SubmitBuildReportCmd.Flags().StringP("schema-version", "", "", "Report schema version (one of: 1)")
	SubmitBuildReportCmd.Flags().StringP("skill-gaps", "", "", "Places where the skill instructions were missing or wrong (JSON)")
	SubmitBuildReportCmd.Flags().StringP("timing", "", "", "Run start and end times and per-stage timings (JSON)")
	SubmitBuildReportCmd.Flags().StringP("tokens", "", "", "Token usage, only when the tool reports its own usage (JSON)")
	SubmitBuildReportCmd.Flags().StringP("user-feedback", "", "", "Feedback the user gave during the run (JSON)")
	SubmitBuildReportCmd.Flags().StringP("what-went-well", "", "", "Things that worked well during the run (JSON)")
	SubmitBuildReportCmd.MarkFlagRequired("consent")
	SubmitBuildReportCmd.MarkFlagRequired("generated-at")
	SubmitBuildReportCmd.MarkFlagRequired("info")
	SubmitBuildReportCmd.MarkFlagRequired("issues")
	SubmitBuildReportCmd.MarkFlagRequired("previous-runs")
	SubmitBuildReportCmd.MarkFlagRequired("run")
	SubmitBuildReportCmd.MarkFlagRequired("schema-version")
	SubmitBuildReportCmd.MarkFlagRequired("skill-gaps")
	SubmitBuildReportCmd.MarkFlagRequired("timing")
	SubmitBuildReportCmd.MarkFlagRequired("tokens")
	SubmitBuildReportCmd.MarkFlagRequired("user-feedback")
	SubmitBuildReportCmd.MarkFlagRequired("what-went-well")
}

func submitBuildReportBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	consentVal, _ := cmd.Flags().GetString("consent")
	body["consent"] = utils.JSONValue(consentVal, "consent")
	generatedAtVal, _ := cmd.Flags().GetString("generated-at")
	body["generated_at"] = generatedAtVal
	infoVal, _ := cmd.Flags().GetString("info")
	body["info"] = utils.JSONValue(infoVal, "info")
	issuesVal, _ := cmd.Flags().GetString("issues")
	body["issues"] = utils.JSONValue(issuesVal, "issues")
	if cmd.Flags().Changed("logging-gaps") {
		v, _ := cmd.Flags().GetString("logging-gaps")
		body["logging_gaps"] = utils.JSONValue(v, "logging-gaps")
	}
	previousRunsVal, _ := cmd.Flags().GetString("previous-runs")
	body["previous_runs"] = utils.JSONValue(previousRunsVal, "previous-runs")
	runVal, _ := cmd.Flags().GetString("run")
	body["run"] = utils.JSONValue(runVal, "run")
	schemaVersionVal, _ := cmd.Flags().GetString("schema-version")
	body["schema_version"] = schemaVersionVal
	skillGapsVal, _ := cmd.Flags().GetString("skill-gaps")
	body["skill_gaps"] = utils.JSONValue(skillGapsVal, "skill-gaps")
	timingVal, _ := cmd.Flags().GetString("timing")
	body["timing"] = utils.JSONValue(timingVal, "timing")
	tokensVal, _ := cmd.Flags().GetString("tokens")
	body["tokens"] = utils.JSONValue(tokensVal, "tokens")
	userFeedbackVal, _ := cmd.Flags().GetString("user-feedback")
	body["user_feedback"] = utils.JSONValue(userFeedbackVal, "user-feedback")
	whatWentWellVal, _ := cmd.Flags().GetString("what-went-well")
	body["what_went_well"] = utils.JSONValue(whatWentWellVal, "what-went-well")
	return body
}
