package domains

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GenerateFreeSubdomainCmd = &cobra.Command{
	Use:               "generate-free-subdomain",
	Short:             "Generate a free subdomain",
	Long:              "Generate a unique free subdomain that can be used for hosting services without purchasing custom domains.\nFree subdomains allow you to start using hosting services immediately\nand you can always connect a custom domain to your site later.",
	Example:           "  hostinger hosting domains generate-free-subdomain",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingGenerateAFreeSubdomainV1WithResponse(ctx)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(GenerateFreeSubdomainCmd)
}
