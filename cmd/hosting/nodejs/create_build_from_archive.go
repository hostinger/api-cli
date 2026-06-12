package nodejs

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateBuildFromArchiveCmd = &cobra.Command{
	Use:   "create-build-from-archive <username> <domain>",
	Short: "Create NodeJS build from archive",
	Long:  "Upload a project archive, auto-detect build settings, and immediately start a Node.js build.\n\nThis is the recommended single-step approach for deploying a Node.js application.\nThe archive is uploaded to the website's file storage, build settings are auto-detected\nfrom the package.json inside the archive, and the build process starts automatically.\nOptional override fields take precedence over auto-detected values.\nMaximum archive size is 50MB.\n\nBefore archiving, exclude `node_modules/` and any build output directories\n(e.g. `dist/`, `.next/`, `build/`) — they are not needed because the build\nprocess runs the install step automatically, and including them unnecessarily\nincreases the archive size. This also helps keep the archive well under the 50MB limit.\n\nExample (zip):\n```\nzip -r archive.zip . --exclude \"node_modules/*\" --exclude \"dist/*\"\n```\n\nThe returned build `uuid` can be used to poll progress and retrieve logs via\nthe `Get Node.js Build Logs` endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "app-type", []string{"create-react-app", "vite", "angular", "react", "vue", "parcel", "express", "fastify", "nest"})
		utils.EnumCheck(cmd, "package-manager", []string{"npm", "yarn", "pnpm"})
		payload, err := json.Marshal(createBuildFromArchiveBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingCreateNodeJSBuildFromArchiveV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateBuildFromArchiveCmd.Flags().StringP("app-type", "", "", "Node.js application type override (one of: create-react-app, vite, angular, react, vue, parcel, express, fastify, nest)")
	CreateBuildFromArchiveCmd.Flags().StringP("archive", "", "", "Project archive file (.zip, .tar.gz, or .tgz), maximum 50MB")
	CreateBuildFromArchiveCmd.Flags().StringP("build-script", "", "", "Build script override")
	CreateBuildFromArchiveCmd.Flags().StringP("entry-file", "", "", "Main entry point file override")
	CreateBuildFromArchiveCmd.Flags().IntP("node-version", "", 0, "Node.js version override (auto-detected from package.json if omitted) (one of: 18, 20, 22, 24)")
	CreateBuildFromArchiveCmd.Flags().StringP("output-directory", "", "", "Build output directory override relative to the root directory")
	CreateBuildFromArchiveCmd.Flags().StringP("package-manager", "", "", "Package manager override (one of: npm, yarn, pnpm)")
	CreateBuildFromArchiveCmd.Flags().StringP("root-directory", "", "", "Application root directory override (where package.json is located) relative to public_html")
	CreateBuildFromArchiveCmd.MarkFlagRequired("archive")
}

func createBuildFromArchiveBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("app-type") {
		v, _ := cmd.Flags().GetString("app-type")
		body["app_type"] = v
	}
	archiveVal, _ := cmd.Flags().GetString("archive")
	body["archive"] = archiveVal
	if cmd.Flags().Changed("build-script") {
		v, _ := cmd.Flags().GetString("build-script")
		body["build_script"] = v
	}
	if cmd.Flags().Changed("entry-file") {
		v, _ := cmd.Flags().GetString("entry-file")
		body["entry_file"] = v
	}
	if cmd.Flags().Changed("node-version") {
		v, _ := cmd.Flags().GetInt("node-version")
		body["node_version"] = v
	}
	if cmd.Flags().Changed("output-directory") {
		v, _ := cmd.Flags().GetString("output-directory")
		body["output_directory"] = v
	}
	if cmd.Flags().Changed("package-manager") {
		v, _ := cmd.Flags().GetString("package-manager")
		body["package_manager"] = v
	}
	if cmd.Flags().Changed("root-directory") {
		v, _ := cmd.Flags().GetString("root-directory")
		body["root_directory"] = v
	}
	return body
}
