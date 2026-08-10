package segments

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListProfileCmd = &cobra.Command{
	Use:   "list-profile <profile-uuid>",
	Short: "List profile segments",
	Long:  "Get a paginated list of the segments defined in a profile.\n\nEach entry carries the number of contacts currently matching it, which is recalculated on\nread rather than stored. Use `count_type` to count either every matching contact or only\nthe subscribed ones.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "count-type", []string{"all", "subscribed"})
		r, err := api.Request().ReachListProfileSegmentsV1WithResponse(context.TODO(), args[0], listProfileParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListProfileCmd.Flags().StringP("count-type", "", "all", "Which matching contacts to count for each segment (one of: all, subscribed)")
	ListProfileCmd.Flags().IntP("page", "", 0, "Page number")
	ListProfileCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listProfileParams(cmd *cobra.Command) *client.ReachListProfileSegmentsV1Params {
	params := &client.ReachListProfileSegmentsV1Params{}
	if cmd.Flags().Changed("count-type") {
		v, _ := cmd.Flags().GetString("count-type")
		e := client.ReachListProfileSegmentsV1ParamsCountType(v)
		params.CountType = &e
	}
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
