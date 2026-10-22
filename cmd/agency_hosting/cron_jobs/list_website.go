package cron_jobs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListWebsiteCmd = &cobra.Command{
	Use:               "list-website <website_uid>",
	Short:             "List website cron jobs",
	Long:              "Returns a paginated list of cron jobs configured for an Agency Plan website.\n\nEach entry includes the schedule expression and the command executed on that schedule.",
	Example:           "  hostinger agency-hosting cron-jobs list-website zpwlGlp19",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listWebsiteParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingListWebsiteCronJobsV1WithResponse(ctx, args[0], params)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		if all, _ := cmd.Flags().GetBool("all"); all {
			call = utils.AllPages(call, func(page int) { params.Page = &page })
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(ListWebsiteCmd, utils.StringArg("website_uid"))
	ListWebsiteCmd.Flags().IntP("page", "", 0, "Page number")
	ListWebsiteCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListWebsiteCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListWebsiteCmd.RegisterFlagCompletionFunc("per-page", cobra.NoFileCompletions)
	ListWebsiteCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListWebsiteCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listWebsiteParams(cmd *cobra.Command) *client.AgencyHostingListWebsiteCronJobsV1Params {
	params := &client.AgencyHostingListWebsiteCronJobsV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	return params
}
