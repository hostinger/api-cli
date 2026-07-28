package mailboxes

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <order-id>",
	Short: "List mailboxes",
	Long:  "Retrieve a paginated list of mailboxes belonging to a mail order.\n\nUse this endpoint to monitor mailboxes of your mail service, including\ntheir status, enabled protocols, attached resource counts, and\nperiodically synced usage numbers (usage may lag behind live values).",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "sort", []string{"address", "-address"})
		r, err := api.Request().MailListMailboxesV1WithResponse(context.TODO(), args[0], listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("search", "", "", "Filter mailboxes whose email address contains the given string")
	ListCmd.Flags().StringP("sort", "", "address", "Sort mailboxes by field. Prefix with `-` for descending order. (one of: address, -address)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listParams(cmd *cobra.Command) *client.MailListMailboxesV1Params {
	params := &client.MailListMailboxesV1Params{}
	if cmd.Flags().Changed("search") {
		v, _ := cmd.Flags().GetString("search")
		params.Search = &v
	}
	if cmd.Flags().Changed("sort") {
		v, _ := cmd.Flags().GetString("sort")
		e := client.MailListMailboxesV1ParamsSort(v)
		params.Sort = &e
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
