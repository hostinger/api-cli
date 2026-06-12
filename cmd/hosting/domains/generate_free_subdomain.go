package domains

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var GenerateFreeSubdomainCmd = &cobra.Command{
	Use:   "generate-free-subdomain",
	Short: "Generate a free subdomain",
	Long:  "Generate a unique free subdomain that can be used for hosting services without purchasing custom domains.\nFree subdomains allow you to start using hosting services immediately\nand you can always connect a custom domain to your site later.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGenerateAFreeSubdomainV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
