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
	Long:  "Deprecated: use `POST /api/hosting/v1/orders/{order_id}/onboardings` (start website setup)\ninstead. It takes the same `domain` and `datacenter_code`, can also install WordPress, and\nreturns the created setup.\n\nCreate a new website for the authenticated client.\n\nYou must choose which hosting order to create this website on. Pass that\norder as `order_id` together with the domain name. List orders to see\navailable IDs; the website is provisioned on that order's hosting plan.\n\nOnly Web and Cloud hosting orders are accepted. To create a website on an Agency\nPlan order, use `POST /api/agency-hosting/v1/orders/{order_id}/websites/setups`.\n\nWhen the order has no hosting account yet, a new hosting account is set up first, in the\ndatacenter given by `datacenter_code` or the order's default datacenter when it is omitted.\nSubsequent websites will be hosted on the same datacenter automatically.\n\nWebsite creation is asynchronous and takes up to a few minutes. Poll the list website\nsetups endpoint with the `domain` filter every 10 to 15 seconds and wait for `status:\ncompleted` before uploading files, deploying or creating databases. While the setup is\n`running`, endpoints that operate on the website may respond with `404` or `409`.\n`is_enabled` on the websites list reflects suspension, not readiness.",
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
	CreateCmd.Flags().StringP("datacenter-code", "", "", "Datacenter for the hosting account when the order does not have one yet. Omit to use the order's default datacenter. Ignored when the order already has a hosting account. Must be an available datacenter; list available datacenters to find valid codes.")
	CreateCmd.Flags().StringP("domain", "", "", "Domain name for the website. Cannot start with \"www.\" and can have at most one subdomain level, e.g. `blog.example.com`.")
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
