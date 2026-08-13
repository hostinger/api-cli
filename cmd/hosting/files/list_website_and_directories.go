package files

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListWebsiteAndDirectoriesCmd = &cobra.Command{
	Use:   "list-website-and-directories <username> <domain>",
	Short: "List website files and directories",
	Long:  "List files and directories under a website's document root.\n\nUse `directory` to browse a subdirectory relative to the document root. Symlinked entries\nare listed but never traversed into or resolved.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListWebsiteFilesAndDirectoriesV1WithResponse(context.TODO(), args[0], args[1], listWebsiteAndDirectoriesParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListWebsiteAndDirectoriesCmd.Flags().StringP("directory", "", "", "Directory path to check")
	ListWebsiteAndDirectoriesCmd.Flags().IntP("max-depth", "", 5, "How many directory levels deep to recurse.")
	ListWebsiteAndDirectoriesCmd.Flags().IntP("max-items", "", 1000, "Max number of entries to return in this page.")
	ListWebsiteAndDirectoriesCmd.Flags().IntP("offset", "", 0, "Number of entries to skip. Page with offset + item count until reaching total_items.")
	ListWebsiteAndDirectoriesCmd.Flags().StringSliceP("file-types", "", nil, "Filter by entry type, e.g. file,directory. Omit for all types. (one of: file, directory, symlink, other)")
}

func listWebsiteAndDirectoriesParams(cmd *cobra.Command) *client.HostingListWebsiteFilesAndDirectoriesV1Params {
	params := &client.HostingListWebsiteFilesAndDirectoriesV1Params{}
	if cmd.Flags().Changed("directory") {
		v, _ := cmd.Flags().GetString("directory")
		params.Directory = &v
	}
	if cmd.Flags().Changed("max-depth") {
		v, _ := cmd.Flags().GetInt("max-depth")
		params.MaxDepth = &v
	}
	if cmd.Flags().Changed("max-items") {
		v, _ := cmd.Flags().GetInt("max-items")
		params.MaxItems = &v
	}
	if cmd.Flags().Changed("offset") {
		v, _ := cmd.Flags().GetInt("offset")
		params.Offset = &v
	}
	if cmd.Flags().Changed("file-types") {
		v, _ := cmd.Flags().GetStringSlice("file-types")
		es := make([]client.HostingListWebsiteFilesAndDirectoriesV1ParamsFileTypes, len(v))
		for i, s := range v {
			es[i] = client.HostingListWebsiteFilesAndDirectoriesV1ParamsFileTypes(s)
		}
		params.FileTypes = &es
	}
	return params
}
