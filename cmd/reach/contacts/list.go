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

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contacts",
	Long:  "Get a list of contacts, optionally filtered by group and subscription status.\n\nThis endpoint returns a paginated list of contacts with their basic information.\nYou can filter contacts by group UUID and subscription status.\n\n**Deprecated.** This endpoint cannot target a profile, so it always falls back to the\nclient's default profile and cannot list contacts of any other profile. Use\n`GET /api/reach/v1/profiles/{profileUuid}/contacts` instead, which also replaces the\ngroup filter with a tag filter.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "subscription-status", []string{"subscribed", "unsubscribed", "confirmed", "pending"})
		r, err := api.Request().ReachListContactsV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("group-uuid", "", "", "Filter contacts by group UUID")
	ListCmd.Flags().StringP("subscription-status", "", "", "Filter contacts by subscription status (one of: subscribed, unsubscribed, confirmed, pending)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
}

func listParams(cmd *cobra.Command) *client.ReachListContactsV1Params {
	params := &client.ReachListContactsV1Params{}
	if cmd.Flags().Changed("group-uuid") {
		v, _ := cmd.Flags().GetString("group-uuid")
		params.GroupUuid = &v
	}
	if cmd.Flags().Changed("subscription-status") {
		v, _ := cmd.Flags().GetString("subscription-status")
		e := client.ReachListContactsV1ParamsSubscriptionStatus(v)
		params.SubscriptionStatus = &e
	}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	return params
}
