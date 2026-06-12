package databases

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <username>",
	Short: "List account databases",
	Long:  "Returns a paginated list of databases for the specified account.\n\nUse the domain and is_assigned filters to find databases assigned to a specific domain.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListAccountDatabasesV1WithResponse(context.TODO(), args[0], listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
	ListCmd.Flags().StringP("domain", "", "", "Filter by domain name (exact match)")
	ListCmd.Flags().BoolP("is-assigned", "", false, "When used with domain, return only databases assigned to that domain.")
	ListCmd.Flags().StringP("search", "", "", "Search databases by name, user, or creation date.")
}

func listParams(cmd *cobra.Command) *client.HostingListAccountDatabasesV1Params {
	params := &client.HostingListAccountDatabasesV1Params{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("per-page") {
		v, _ := cmd.Flags().GetInt("per-page")
		params.PerPage = &v
	}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	if cmd.Flags().Changed("is-assigned") {
		v, _ := cmd.Flags().GetBool("is-assigned")
		params.IsAssigned = &v
	}
	if cmd.Flags().Changed("search") {
		v, _ := cmd.Flags().GetString("search")
		params.Search = &v
	}
	return params
}
