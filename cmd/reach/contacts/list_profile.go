package contacts

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
	Short: "List profile contacts",
	Long:  "Get a paginated list of contacts belonging to a profile.\n\nContacts can be filtered by subscription status, by tag, and by an email search term.\nThe `meta.total` field of the response is the number of contacts matching the filters,\nso calling this endpoint without filters gives the profile's total contact count.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "subscription-status", []string{"subscribed", "unsubscribed", "confirmed", "pending"})
		r, err := api.Request().ReachListProfileContactsV1WithResponse(context.TODO(), args[0], listProfileParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListProfileCmd.Flags().StringP("subscription-status", "", "", "Filter contacts by subscription status (one of: subscribed, unsubscribed, confirmed, pending)")
	ListProfileCmd.Flags().StringP("tag-uuid", "", "", "Filter contacts by tag UUID")
	ListProfileCmd.Flags().StringP("search", "", "", "Search contacts by email")
	ListProfileCmd.Flags().IntP("page", "", 0, "Page number")
	ListProfileCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listProfileParams(cmd *cobra.Command) *client.ReachListProfileContactsV1Params {
	params := &client.ReachListProfileContactsV1Params{}
	if cmd.Flags().Changed("subscription-status") {
		v, _ := cmd.Flags().GetString("subscription-status")
		e := client.ReachListProfileContactsV1ParamsSubscriptionStatus(v)
		params.SubscriptionStatus = &e
	}
	if cmd.Flags().Changed("tag-uuid") {
		v, _ := cmd.Flags().GetString("tag-uuid")
		t := utils.StringToUUID(v)
		params.TagUuid = &t
	}
	if cmd.Flags().Changed("search") {
		v, _ := cmd.Flags().GetString("search")
		params.Search = &v
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
