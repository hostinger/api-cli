package domains

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateWebsiteSubdomainCmd = &cobra.Command{
	Use:   "create-website-subdomain <username> <domain>",
	Short: "Create website subdomain",
	Long:  "Create a new subdomain for the selected website.\n\nProvide a subdomain prefix and, optionally, a custom directory or the\nwebsite public directory to use as the subdomain root.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createWebsiteSubdomainBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingCreateWebsiteSubdomainV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateWebsiteSubdomainCmd.Flags().StringP("directory", "", "", "Directory name for the subdomain relative to the website root")
	CreateWebsiteSubdomainCmd.Flags().BoolP("is-using-public-directory", "", false, "Use the website public directory as the subdomain root directory")
	CreateWebsiteSubdomainCmd.Flags().StringP("subdomain", "", "", "Subdomain prefix to create under the selected website")
	CreateWebsiteSubdomainCmd.MarkFlagRequired("subdomain")
}

func createWebsiteSubdomainBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("directory") {
		v, _ := cmd.Flags().GetString("directory")
		body["directory"] = v
	}
	if cmd.Flags().Changed("is-using-public-directory") {
		v, _ := cmd.Flags().GetBool("is-using-public-directory")
		body["is_using_public_directory"] = v
	}
	subdomainVal, _ := cmd.Flags().GetString("subdomain")
	body["subdomain"] = subdomainVal
	return body
}
