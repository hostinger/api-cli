package files

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListWebsiteAndDirectoriesCmd = &cobra.Command{
	Use:     "list-website-and-directories <username> <domain>",
	Short:   "List website files and directories",
	Long:    "List files and directories under a website's document root.\n\nUse `directory` to browse a subdirectory relative to the document root. Symlinked entries\nare listed but never traversed into or resolved.",
	Example: "  hostinger hosting files list-website-and-directories u123456789 mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listWebsiteAndDirectoriesParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingListWebsiteFilesAndDirectoriesV1WithResponse(ctx, args[0], args[1], params)
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
	utils.RegisterArgs(ListWebsiteAndDirectoriesCmd, utils.StringArg("username"), utils.StringArg("domain"))
	ListWebsiteAndDirectoriesCmd.Flags().StringP("directory", "", "", "Directory path to check")
	ListWebsiteAndDirectoriesCmd.Flags().IntP("max-depth", "", 5, "How many directory levels deep to recurse.")
	ListWebsiteAndDirectoriesCmd.Flags().IntP("max-items", "", 1000, "Max number of entries to return in this page.")
	ListWebsiteAndDirectoriesCmd.Flags().IntP("offset", "", 0, "Number of entries to skip. Page with offset + item count until reaching total_items.")
	ListWebsiteAndDirectoriesCmd.Flags().StringSliceP("file-types", "", nil, "Filter by entry type, e.g. file,directory. Omit for all types. (one of: file, directory, symlink, other)")
	ListWebsiteAndDirectoriesCmd.RegisterFlagCompletionFunc("directory", cobra.NoFileCompletions)
	ListWebsiteAndDirectoriesCmd.RegisterFlagCompletionFunc("max-depth", cobra.NoFileCompletions)
	ListWebsiteAndDirectoriesCmd.RegisterFlagCompletionFunc("max-items", cobra.NoFileCompletions)
	ListWebsiteAndDirectoriesCmd.RegisterFlagCompletionFunc("offset", cobra.NoFileCompletions)
	ListWebsiteAndDirectoriesCmd.RegisterFlagCompletionFunc("file-types", cobra.FixedCompletions([]cobra.Completion{"file", "directory", "symlink", "other"}, cobra.ShellCompDirectiveNoFileComp))
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
