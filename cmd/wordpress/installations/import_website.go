package installations

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ImportWebsiteCmd = &cobra.Command{
	Use:     "import-website <username> <domain>",
	Short:   "Import WordPress website",
	Long:    "Import WordPress website to the specified domain.\n\nWARNING: this overwrites the website's existing contents and cannot be undone —\nverify this is intended before calling this endpoint.\n\nThis endpoint allows you to import a WordPress website from archive and\ndatabase files that have been uploaded to the website's directory.",
	Example: "  hostinger wordpress installations import-website u123456789 mydomain.tld --archive-path backup.zip --sql-path database.sql",
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
				r, err := api.CompletionRequest().DomainsGetDomainListV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := importWebsiteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingImportWordPressWebsiteV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ImportWebsiteCmd, utils.StringArg("username"), utils.StringArg("domain"))
	ImportWebsiteCmd.Flags().StringP("archive-path", "", "", "Path to the WordPress archive file (relative to website root)")
	ImportWebsiteCmd.Flags().StringP("sql-path", "", "", "Path to the database SQL file (relative to website root)")
	ImportWebsiteCmd.MarkFlagRequired("archive-path")
	ImportWebsiteCmd.MarkFlagRequired("sql-path")
	ImportWebsiteCmd.RegisterFlagCompletionFunc("archive-path", cobra.NoFileCompletions)
	ImportWebsiteCmd.RegisterFlagCompletionFunc("sql-path", cobra.NoFileCompletions)
}

func importWebsiteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	archivePathVal, _ := cmd.Flags().GetString("archive-path")
	body["archive_path"] = archivePathVal
	sqlPathVal, _ := cmd.Flags().GetString("sql-path")
	body["sql_path"] = sqlPathVal
	return body, nil
}
