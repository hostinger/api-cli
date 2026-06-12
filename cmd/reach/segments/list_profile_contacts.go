package segments

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListProfileContactsCmd = &cobra.Command{
	Use:   "list-profile-contacts <profile-uuid> <segment-uuid>",
	Short: "List profile segment contacts",
	Long:  "Retrieve contacts associated with a specific segment for a given profile.\n\nThis endpoint allows you to fetch and filter contacts that belong to a particular segment,\nidentified by its UUID, scoped to a specific profile.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachListProfileSegmentContactsV1WithResponse(context.TODO(), args[0], args[1], listProfileContactsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListProfileContactsCmd.Flags().IntP("page", "", 0, "Page number")
	ListProfileContactsCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listProfileContactsParams(cmd *cobra.Command) *client.ReachListProfileSegmentContactsV1Params {
	params := &client.ReachListProfileSegmentContactsV1Params{}
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
