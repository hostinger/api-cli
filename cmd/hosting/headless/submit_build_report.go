package headless

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SubmitBuildReportCmd = &cobra.Command{
	Use:               "submit-build-report",
	Short:             "Submit headless build report",
	Long:              "Submits the build report the headless skill writes at the end of a run (including early\nstops), so Hostinger can learn where builds get stuck. Send one report per call, with the\nreport JSON as the request body, unchanged. Only send a report when the user agreed to it.\n\nThe report must match report.schema.json v1; an invalid report returns 422 and is not\nstored. The report is linked to the authenticated account, so do not add account details\nto it.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --consent\n    {\n      \"note\": \"string\",\n      \"sent\": false\n    }\n\n  --info\n    {\n      \"domain\": \"example.com\",\n      \"outcome\": \"live\",\n      \"tool\": \"claude-code\"\n    }\n\n  --issues\n    [\n      {}\n    ]\n\n  --logging-gaps\n    [\n      {}\n    ]\n\n  --previous-runs\n    [\n      {}\n    ]\n\n  --run\n    {\n      \"follow_ups\": [\n        {}\n      ],\n      \"id\": \"2026-10-14T10-00-00Z\",\n      \"kind\": \"build\",\n      \"skill\": \"string\"\n    }\n\n  --skill-gaps\n    [\n      {}\n    ]\n\n  --timing\n    {\n      \"ended_at\": \"string\",\n      \"minutes_total\": 0,\n      \"stages\": [\n        {}\n      ],\n      \"started_at\": \"string\"\n    }\n\n  --tokens\n    {}\n\n  --user-feedback\n    [\n      {}\n    ]\n\n  --what-went-well\n    [\n      {}\n    ]",
	Example:           "  hostinger hosting headless submit-build-report --consent '{\"note\":\"string\",\"sent\":false}' --generated-at 2026-10-14T10:30:00Z --info @info.json --issues '[{}]' --previous-runs '[{}]' --run @run.json --schema-version 1 --skill-gaps '[{}]' --timing @timing.json --tokens '{}' --user-feedback '[{}]' --what-went-well '[{}]'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "schema-version", []string{"1"}); err != nil {
			return err
		}
		body, err := submitBuildReportBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingSubmitHeadlessBuildReportV1WithBodyWithResponse(ctx, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(SubmitBuildReportCmd)
	SubmitBuildReportCmd.Flags().StringP("consent", "", "", "Upload status of the report, as written by the headless skill (JSON, @file or -; shape in the help above)")
	SubmitBuildReportCmd.Flags().StringP("generated-at", "", "", "When the report was written")
	SubmitBuildReportCmd.Flags().StringP("info", "", "", "Website, tool and outcome of the run (JSON, @file or -; shape in the help above)")
	SubmitBuildReportCmd.Flags().StringP("issues", "", "", "Problems hit during the run (JSON, @file or -; shape in the help above)")
	SubmitBuildReportCmd.Flags().StringP("logging-gaps", "", "", "Where the run log is missing or unsure (JSON, @file or -; shape in the help above)")
	SubmitBuildReportCmd.Flags().StringP("previous-runs", "", "", "Earlier runs for the same website (JSON, @file or -; shape in the help above)")
	SubmitBuildReportCmd.Flags().StringP("run", "", "", "The run the report is for (JSON, @file or -; shape in the help above)")
	SubmitBuildReportCmd.Flags().StringP("schema-version", "", "", "Report schema version (one of: 1)")
	SubmitBuildReportCmd.Flags().StringP("skill-gaps", "", "", "Places where the skill instructions were missing or wrong (JSON, @file or -; shape in the help above)")
	SubmitBuildReportCmd.Flags().StringP("timing", "", "", "Run start and end times and per-stage timings (JSON, @file or -; shape in the help above)")
	SubmitBuildReportCmd.Flags().StringP("tokens", "", "", "Token usage, only when the tool reports its own usage (JSON, @file or -; shape in the help above)")
	SubmitBuildReportCmd.Flags().StringP("user-feedback", "", "", "Feedback the user gave during the run (JSON, @file or -; shape in the help above)")
	SubmitBuildReportCmd.Flags().StringP("what-went-well", "", "", "Things that worked well during the run (JSON, @file or -; shape in the help above)")
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
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("consent", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("generated-at", cobra.NoFileCompletions)
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("info", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("issues", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("logging-gaps", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("previous-runs", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("run", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("schema-version", cobra.FixedCompletions([]cobra.Completion{"1"}, cobra.ShellCompDirectiveNoFileComp))
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("skill-gaps", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("timing", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("tokens", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("user-feedback", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	SubmitBuildReportCmd.RegisterFlagCompletionFunc("what-went-well", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func submitBuildReportBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	consentVal, _ := cmd.Flags().GetString("consent")
	consentValJSON, err := utils.JSONValue(consentVal, "consent")
	if err != nil {
		return nil, err
	}
	body["consent"] = consentValJSON
	generatedAtVal, _ := cmd.Flags().GetString("generated-at")
	body["generated_at"] = generatedAtVal
	infoVal, _ := cmd.Flags().GetString("info")
	infoValJSON, err := utils.JSONValue(infoVal, "info")
	if err != nil {
		return nil, err
	}
	body["info"] = infoValJSON
	issuesVal, _ := cmd.Flags().GetString("issues")
	issuesValJSON, err := utils.JSONValue(issuesVal, "issues")
	if err != nil {
		return nil, err
	}
	body["issues"] = issuesValJSON
	if cmd.Flags().Changed("logging-gaps") {
		loggingGapsVal, _ := cmd.Flags().GetString("logging-gaps")
		loggingGapsValJSON, err := utils.JSONValue(loggingGapsVal, "logging-gaps")
		if err != nil {
			return nil, err
		}
		body["logging_gaps"] = loggingGapsValJSON
	}
	previousRunsVal, _ := cmd.Flags().GetString("previous-runs")
	previousRunsValJSON, err := utils.JSONValue(previousRunsVal, "previous-runs")
	if err != nil {
		return nil, err
	}
	body["previous_runs"] = previousRunsValJSON
	runVal, _ := cmd.Flags().GetString("run")
	runValJSON, err := utils.JSONValue(runVal, "run")
	if err != nil {
		return nil, err
	}
	body["run"] = runValJSON
	schemaVersionVal, _ := cmd.Flags().GetString("schema-version")
	body["schema_version"] = schemaVersionVal
	skillGapsVal, _ := cmd.Flags().GetString("skill-gaps")
	skillGapsValJSON, err := utils.JSONValue(skillGapsVal, "skill-gaps")
	if err != nil {
		return nil, err
	}
	body["skill_gaps"] = skillGapsValJSON
	timingVal, _ := cmd.Flags().GetString("timing")
	timingValJSON, err := utils.JSONValue(timingVal, "timing")
	if err != nil {
		return nil, err
	}
	body["timing"] = timingValJSON
	tokensVal, _ := cmd.Flags().GetString("tokens")
	tokensValJSON, err := utils.JSONValue(tokensVal, "tokens")
	if err != nil {
		return nil, err
	}
	body["tokens"] = tokensValJSON
	userFeedbackVal, _ := cmd.Flags().GetString("user-feedback")
	userFeedbackValJSON, err := utils.JSONValue(userFeedbackVal, "user-feedback")
	if err != nil {
		return nil, err
	}
	body["user_feedback"] = userFeedbackValJSON
	whatWentWellVal, _ := cmd.Flags().GetString("what-went-well")
	whatWentWellValJSON, err := utils.JSONValue(whatWentWellVal, "what-went-well")
	if err != nil {
		return nil, err
	}
	body["what_went_well"] = whatWentWellValJSON
	return body, nil
}
