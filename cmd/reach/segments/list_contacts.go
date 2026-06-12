package segments

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListContactsCmd = &cobra.Command{
	Use:   "list-contacts <segment-uuid>",
	Short: "List segment contacts",
	Long:  "Retrieve contacts associated with a specific segment.\n\nThis endpoint allows you to fetch and filter contacts that belong to a particular segment,\nidentified by its UUID.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachListSegmentContactsV1WithResponse(context.TODO(), args[0], listContactsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListContactsCmd.Flags().IntP("page", "", 0, "Page number")
	ListContactsCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listContactsParams(cmd *cobra.Command) *client.ReachListSegmentContactsV1Params {
	params := &client.ReachListSegmentContactsV1Params{}
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
