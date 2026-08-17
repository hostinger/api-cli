package redirects

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteWebsiteCmd = &cobra.Command{
	Use:   "delete-website <username> <domain>",
	Short: "Delete website redirect",
	Long:  "Permanently deletes the redirect identified by its source URL.\n\nPass the `from` value exactly as returned by the list redirects endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingDeleteWebsiteRedirectV1WithResponse(context.TODO(), args[0], args[1], deleteWebsiteParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DeleteWebsiteCmd.Flags().StringP("from", "", "", "Source URL returned by the list redirects endpoint.")
	DeleteWebsiteCmd.MarkFlagRequired("from")
}

func deleteWebsiteParams(cmd *cobra.Command) *client.HostingDeleteWebsiteRedirectV1Params {
	params := &client.HostingDeleteWebsiteRedirectV1Params{}
	fromVal, _ := cmd.Flags().GetString("from")
	params.From = fromVal
	return params
}
