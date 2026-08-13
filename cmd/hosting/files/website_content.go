package files

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var WebsiteContentCmd = &cobra.Command{
	Use:   "website-content <username> <domain>",
	Short: "Get website file content",
	Long:  "Get a single file's content, relative to a website's document root.\n\nRead-only; refuses symlinks, oversized files, non-text file types, and files identified as\ncontaining secrets (e.g. credential files) — none of these are returned by this endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetWebsiteFileContentV1WithResponse(context.TODO(), args[0], args[1], websiteContentParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	WebsiteContentCmd.Flags().StringP("path", "", "", "File path, relative to the document root.")
	WebsiteContentCmd.Flags().IntP("from-line", "", 0, "Line offset to start reading from.")
	WebsiteContentCmd.Flags().IntP("max-lines", "", 5000, "Max number of lines to return.")
	WebsiteContentCmd.MarkFlagRequired("path")
}

func websiteContentParams(cmd *cobra.Command) *client.HostingGetWebsiteFileContentV1Params {
	params := &client.HostingGetWebsiteFileContentV1Params{}
	pathVal, _ := cmd.Flags().GetString("path")
	params.Path = pathVal
	if cmd.Flags().Changed("from-line") {
		v, _ := cmd.Flags().GetInt("from-line")
		params.FromLine = &v
	}
	if cmd.Flags().Changed("max-lines") {
		v, _ := cmd.Flags().GetInt("max-lines")
		params.MaxLines = &v
	}
	return params
}
