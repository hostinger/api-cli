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

var StartBuildCmd = &cobra.Command{
	Use:   "start-build <username> <domain>",
	Short: "Start Node.js build",
	Long:  "Start a Node.js build process using files already present on the website's file storage.\n\nWARNING: on success this overwrites the website's existing contents and cannot be\nundone — verify this is intended before calling this endpoint.\n\nThe `source_type` must be `archive` and `source_options.archive_path` must point to an\nexisting archive file on the server (relative to the website document root).\nUse the `Generate Upload URL` endpoint to obtain credentials and upload the archive first.\n\nTo auto-detect build settings from an archive before starting, first call the\n`Get Node.js Build Settings from Archive` endpoint.\n\nThe returned build `uuid` can be used to poll progress and retrieve logs via\nthe `Get Node.js Build Logs` endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "app-type", []string{"create-react-app", "gatsby", "vite", "angular", "react", "vue", "parcel", "next", "nuxt", "nest", "express", "fastify", "astro", "svelte", "svelte-kit", "hono", "react-router", "nitro", "other"})
		utils.EnumCheck(cmd, "package-manager", []string{"npm", "yarn", "pnpm"})
		utils.EnumCheck(cmd, "source-type", []string{"archive"})
		payload, err := json.Marshal(startBuildBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingStartNodeJsBuildV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	StartBuildCmd.Flags().StringP("app-type", "", "", "Node.js application type (one of: create-react-app, gatsby, vite, angular, react, vue, parcel, next, nuxt, nest, express, fastify, astro, svelte, svelte-kit, hono, react-router, nitro, other)")
	StartBuildCmd.Flags().StringP("build-script", "", "", "Build script that will be ran to build the application")
	StartBuildCmd.Flags().StringP("entry-file", "", "", "The main entry point file for the application")
	StartBuildCmd.Flags().IntP("node-version", "", 0, "Node.js version (one of: 18, 20, 22, 24)")
	StartBuildCmd.Flags().StringP("output-directory", "", "", "Build output directory relative to the root directory")
	StartBuildCmd.Flags().StringP("package-manager", "", "", "Package manager (one of: npm, yarn, pnpm)")
	StartBuildCmd.Flags().StringP("root-directory", "", "", "Application root directory (where package.json is located) relative to public_html")
	StartBuildCmd.Flags().StringP("source-options", "", "", "Source-specific options (JSON)")
	StartBuildCmd.Flags().StringP("source-type", "", "", "The source type of the files (one of: archive)")
	StartBuildCmd.MarkFlagRequired("app-type")
	StartBuildCmd.MarkFlagRequired("build-script")
	StartBuildCmd.MarkFlagRequired("node-version")
	StartBuildCmd.MarkFlagRequired("output-directory")
	StartBuildCmd.MarkFlagRequired("root-directory")
	StartBuildCmd.MarkFlagRequired("source-options")
	StartBuildCmd.MarkFlagRequired("source-type")
}

func startBuildBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	appTypeVal, _ := cmd.Flags().GetString("app-type")
	body["app_type"] = appTypeVal
	buildScriptVal, _ := cmd.Flags().GetString("build-script")
	body["build_script"] = buildScriptVal
	if cmd.Flags().Changed("entry-file") {
		v, _ := cmd.Flags().GetString("entry-file")
		body["entry_file"] = v
	}
	nodeVersionVal, _ := cmd.Flags().GetInt("node-version")
	body["node_version"] = nodeVersionVal
	outputDirectoryVal, _ := cmd.Flags().GetString("output-directory")
	body["output_directory"] = outputDirectoryVal
	if cmd.Flags().Changed("package-manager") {
		v, _ := cmd.Flags().GetString("package-manager")
		body["package_manager"] = v
	}
	rootDirectoryVal, _ := cmd.Flags().GetString("root-directory")
	body["root_directory"] = rootDirectoryVal
	sourceOptionsVal, _ := cmd.Flags().GetString("source-options")
	body["source_options"] = utils.JSONValue(sourceOptionsVal, "source-options")
	sourceTypeVal, _ := cmd.Flags().GetString("source-type")
	body["source_type"] = sourceTypeVal
	return body
}
