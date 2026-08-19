package installations

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ImportWebsiteCmd = &cobra.Command{
	Use:   "import-website <username> <domain>",
	Short: "Import WordPress website",
	Long:  "Import WordPress website to the specified domain.\n\nWARNING: this overwrites the website's existing contents and cannot be undone —\nverify this is intended before calling this endpoint.\n\nThis endpoint allows you to import a WordPress website from archive and\ndatabase files that have been uploaded to the website's directory.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(importWebsiteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingImportWordPressWebsiteV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ImportWebsiteCmd.Flags().StringP("archive-path", "", "", "Path to the WordPress archive file (relative to website root)")
	ImportWebsiteCmd.Flags().StringP("sql-path", "", "", "Path to the database SQL file (relative to website root)")
	ImportWebsiteCmd.MarkFlagRequired("archive-path")
	ImportWebsiteCmd.MarkFlagRequired("sql-path")
}

func importWebsiteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	archivePathVal, _ := cmd.Flags().GetString("archive-path")
	body["archive_path"] = archivePathVal
	sqlPathVal, _ := cmd.Flags().GetString("sql-path")
	body["sql_path"] = sqlPathVal
	return body
}
