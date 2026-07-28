package catchalls

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCatchAllsCmd = &cobra.Command{
	Use:   "list-catch-alls <order-id>",
	Short: "List catch-alls",
	Long:  "Retrieve a paginated list of catch-alls across all mailboxes of a\nmail order.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().MailListCatchAllsV1WithResponse(context.TODO(), args[0], listCatchAllsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCatchAllsCmd.Flags().IntP("page", "", 0, "Page number")
	ListCatchAllsCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listCatchAllsParams(cmd *cobra.Command) *client.MailListCatchAllsV1Params {
	params := &client.MailListCatchAllsV1Params{}
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
