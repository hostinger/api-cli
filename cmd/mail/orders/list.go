package orders

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
	Short: "Get mail order list",
	Long:  "Retrieve a paginated list of mail orders associated with your account.\n\nUse this endpoint to monitor your mail services, including their status,\nplan, attached domain, and expiration details.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"pending_setup", "active", "suspended"})
		utils.EnumCheck(cmd, "sort", []string{"created_at", "-created_at", "expires_at", "-expires_at"})
		r, err := api.Request().MailGetMailOrderListV1WithResponse(context.TODO(), listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("domain", "", "", "Filter orders by domain name (exact match)")
	ListCmd.Flags().StringP("status", "", "", "Filter orders by status (one of: pending_setup, active, suspended)")
	ListCmd.Flags().BoolP("is-trial", "", false, "Filter orders by trial state")
	ListCmd.Flags().StringP("sort", "", "-created_at", "Sort orders by field. Prefix with `-` for descending order. (one of: created_at, -created_at, expires_at, -expires_at)")
	ListCmd.Flags().IntP("page", "", 0, "Page number")
	ListCmd.Flags().IntP("per-page", "", 25, "Number of items per page")
}

func listParams(cmd *cobra.Command) *client.MailGetMailOrderListV1Params {
	params := &client.MailGetMailOrderListV1Params{}
	if cmd.Flags().Changed("domain") {
		v, _ := cmd.Flags().GetString("domain")
		params.Domain = &v
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		e := client.MailGetMailOrderListV1ParamsStatus(v)
		params.Status = &e
	}
	if cmd.Flags().Changed("is-trial") {
		v, _ := cmd.Flags().GetBool("is-trial")
		params.IsTrial = &v
	}
	if cmd.Flags().Changed("sort") {
		v, _ := cmd.Flags().GetString("sort")
		e := client.MailGetMailOrderListV1ParamsSort(v)
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
