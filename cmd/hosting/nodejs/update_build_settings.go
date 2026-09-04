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

var UpdateBuildSettingsCmd = &cobra.Command{
	Use:   "update-build-settings <username> <domain>",
	Short: "Update Node.js build settings",
	Long:  "Replaces the build settings stored for the website. Send the full set: `node_version` is\nrequired and every nullable field you omit is stored as null. Creates the settings when\nnone exist yet.\n\nThis does not start a build. Stored settings drive Git auto-deployment builds; a build\nstarted through the API uses the values sent in that request, so to rebuild with corrected\nsettings call `Start Node.js build` with the same values. Typical fixes: a wrong `app_type`\nafter auto-detection, or a missing `entry_file` for express, fastify, nest, nuxt and hono\napps.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "app-type", []string{"create-react-app", "gatsby", "vite", "angular", "react", "vue", "parcel", "next", "nuxt", "nest", "express", "fastify", "astro", "svelte", "svelte-kit", "hono", "react-router", "nitro", "other"})
		utils.EnumCheck(cmd, "package-manager", []string{"npm", "yarn", "pnpm"})
		payload, err := json.Marshal(updateBuildSettingsBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingUpdateNodeJsBuildSettingsV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateBuildSettingsCmd.Flags().StringP("app-type", "", "", "Node.js application framework. Set it explicitly when auto-detection picked the wrong one. (one of: create-react-app, gatsby, vite, angular, react, vue, parcel, next, nuxt, nest, express, fastify, astro, svelte, svelte-kit, hono, react-router, nitro, other)")
	UpdateBuildSettingsCmd.Flags().StringP("build-script", "", "", "The package.json script that builds the application")
	UpdateBuildSettingsCmd.Flags().StringP("entry-file", "", "", "The main entry point file for the application\n(required for express, fastify, nest, nuxt and hono app types)")
	UpdateBuildSettingsCmd.Flags().IntP("node-version", "", 0, "Node.js major version (one of: 18, 20, 22, 24)")
	UpdateBuildSettingsCmd.Flags().StringP("output-directory", "", "", "Build output directory relative to the root directory")
	UpdateBuildSettingsCmd.Flags().StringP("package-manager", "", "", "Package manager used to install dependencies (one of: npm, yarn, pnpm)")
	UpdateBuildSettingsCmd.Flags().StringP("root-directory", "", "", "Application root directory (where package.json is located) relative to public_html.\nOmit it, or send \".\", for public_html itself.")
	UpdateBuildSettingsCmd.MarkFlagRequired("node-version")
}

func updateBuildSettingsBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("app-type") {
		v, _ := cmd.Flags().GetString("app-type")
		body["app_type"] = v
	}
	if cmd.Flags().Changed("build-script") {
		v, _ := cmd.Flags().GetString("build-script")
		body["build_script"] = v
	}
	if cmd.Flags().Changed("entry-file") {
		v, _ := cmd.Flags().GetString("entry-file")
		body["entry_file"] = v
	}
	nodeVersionVal, _ := cmd.Flags().GetInt("node-version")
	body["node_version"] = nodeVersionVal
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
