package websites

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create website",
	Long:  "Create a new website for the authenticated client.\n\nYou must choose which hosting order to create this website on. Pass that\norder as `order_id` together with the domain name. List orders to see\navailable IDs; the website is provisioned on that order's hosting plan.\n\nOnly Web and Cloud hosting orders are accepted. To create a website on an Agency\nPlan order, use `POST /api/agency-hosting/v1/orders/{order_id}/websites/setups`.\n\nThe datacenter_code parameter is required when creating the first website\non a new hosting plan - this will set up and configure new hosting account\nin the selected datacenter.\n\nSubsequent websites will be hosted on the same datacenter automatically.\n\nWebsite creation is asynchronous and takes up to a few minutes. Poll the list website\nsetups endpoint with the `domain` filter every 10 to 15 seconds and wait for `status:\ncompleted` before uploading files, deploying or creating databases. While the setup is\n`running`, endpoints that operate on the website may respond with `404` or `409`.\n`is_enabled` on the websites list reflects suspension, not readiness.",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingCreateWebsiteV1WithBodyWithResponse(context.TODO(), "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("datacenter-code", "", "", "Datacenter code. This parameter is required when creating the first website on a new hosting plan.")
	CreateCmd.Flags().StringP("domain", "", "", "Domain name for the website. Cannot start with \"www.\"")
	CreateCmd.Flags().IntP("order-id", "", 0, "Hosting order ID to create this website on. Choose the order whose hosting plan should host the new website. List orders to find available IDs.")
	CreateCmd.MarkFlagRequired("domain")
	CreateCmd.MarkFlagRequired("order-id")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("datacenter-code") {
		v, _ := cmd.Flags().GetString("datacenter-code")
		body["datacenter_code"] = v
	}
	domainVal, _ := cmd.Flags().GetString("domain")
	body["domain"] = domainVal
	orderIdVal, _ := cmd.Flags().GetInt("order-id")
	body["order_id"] = orderIdVal
	return body
}
